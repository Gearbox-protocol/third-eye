drop MATERIALIZED VIEW ranking_7d;
drop MATERIALIZED VIEW ranking_30d;
drop procedure rankings;
drop FUNCTION ranking_by_period;
create OR REPLACE FUNCTION ranking_by_period(BIGINT)
returns table (
		old_collateral DOUBLE PRECISION,
		old_profit DOUBLE PRECISION,
		old_total DOUBLE PRECISION,
		old_collateral_underlying DOUBLE PRECISION,
		old_profit_underlying DOUBLE PRECISION,
        sid varchar(100),
		new_collateral DOUBLE PRECISION,
		new_profit DOUBLE PRECISION,
		new_total DOUBLE PRECISION,
		new_collateral_underlying DOUBLE PRECISION,
		new_profit_underlying DOUBLE PRECISION,
        session_id varchar(100),
        current_block integer,
		profit_usd DOUBLE PRECISION,
        collateral_usd DOUBLE PRECISION,
		profit_underlying DOUBLE PRECISION,
		collateral_underlying DOUBLE PRECISION,
        roi DOUBLE PRECISION
	) 
language plpgsql
as $$
DECLARE
BEGIN
    RETURN QUERY SELECT *, (t2.new_profit-t1.old_profit) profit_usd, t1.old_collateral collateral_usd,
		(t2.new_profit_underlying-t1.old_profit_underlying) profit_underlying, t1.old_collateral_underlying collateral_underlying,
			(t2.new_profit-t1.old_profit)/(t1.old_collateral) roi  FROM
        (SELECT distinct on (d.session_id) d.collateral_usd old_collateral, d.profit_usd as old_profit, d.total_value_usd old_total,
			d.collateral_underlying old_collateral_underlying, d.profit_underlying as old_profit_underlying,
            d.session_id sid
			FROM debts d WHERE block_num >= (SELECT min(id) FROM blocks WHERE timestamp > (extract(epoch from now())::bigint - $1)) 
            order by d.session_id, block_num) t1
        JOIN (SELECT distinct on (d.session_id) d.collateral_usd new_collateral, d.profit_usd as new_profit, d.total_value_usd new_total,
			d.collateral_underlying new_collateral_underlying, d.profit_underlying as new_profit_underlying,
			d.session_id, block_num current_block
			FROM debts d WHERE block_num >= (SELECT min(id) FROM blocks WHERE timestamp > (extract(epoch from now())::bigint - $1)) 
            order by d.session_id, block_num DESC) t2
            ON t1.sid = t2.session_id;
END $$;

CREATE  MATERIALIZED VIEW ranking_7d AS
    SELECT * FROM ranking_by_period(7*86400);
CREATE  MATERIALIZED VIEW ranking_30d AS
    SELECT * FROM ranking_by_period(30*86400);

create OR REPLACE PROCEDURE rankings()
language plpgsql
as $$
DECLARE
BEGIN
    REFRESH MATERIALIZED VIEW ranking_7d;
    REFRESH MATERIALIZED VIEW ranking_30d;
END $$;



-- SELECT array_to_json(array_agg(row_to_json(t))) FROM  (select distinct on(type) * from dao_operations order by type) t ;

-- update price_feeds set feed=address from 
-- (select address, details->>'token' as token, type from sync_adapters where type='QueryPriceFeed' and disabled='f') t
-- WHERE t.token=price_feeds.token;