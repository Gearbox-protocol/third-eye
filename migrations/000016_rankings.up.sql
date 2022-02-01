create OR REPLACE FUNCTION ranking_by_period(BIGINT)
returns table (
		old_collateral DOUBLE PRECISION,
		old_profit DOUBLE PRECISION,
		old_total DOUBLE PRECISION,
        sid varchar(100),
		new_collateral DOUBLE PRECISION,
		new_profit DOUBLE PRECISION,
		new_total DOUBLE PRECISION,
        session_id varchar(100),
        current_block integer,
		profit_usd DOUBLE PRECISION,
        collateral_usd DOUBLE PRECISION,
        apy DOUBLE PRECISION
	) 
language plpgsql
as $$
DECLARE
BEGIN
    RETURN QUERY WITH common AS 
		(SELECT min(d1.block_num) , max(d1.block_num) , d1.session_id from debts d1
			JOIN (SELECT * FROM blocks WHERE timestamp > (extract(epoch from now())::bigint - $1)) b 
			ON b.id = d1.block_num group by d1.session_id)
		SELECT *, (t2.new_profit-t1.old_profit) profit_usd, t1.old_collateral collateral_usd, 
			(t2.new_profit-t1.old_profit)/(t1.old_collateral) apy  FROM
		(SELECT 
			d.collateral_usd old_collateral, d.profit_usd as old_profit, d.total_value_usd old_total,
			d.session_id sid
			FROM debts d JOIN common ON common.min = d.block_num AND d.session_id=common.session_id) t1
		JOIN (SELECT 
			d.collateral_usd new_collateral, d.profit_usd as new_profit, d.total_value_usd new_total,
			d.session_id, common.max current_block
			FROM debts d JOIN common ON common.max = d.block_num AND d.session_id=common.session_id) t2 
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

-- drop MATERIALIZED VIEW ranking_7d;
-- drop MATERIALIZED VIEW ranking_30d;
-- drop procedure rankings;
-- drop FUNCTION ranking_by_period;

-- SELECT array_to_json(array_agg(row_to_json(t))) FROM  (select distinct on(type) * from dao_operations order by type) t ;