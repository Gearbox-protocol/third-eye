drop MATERIALIZED VIEW IF EXISTS ranking_7d;
drop MATERIALIZED VIEW IF EXISTS ranking_30d;
drop procedure IF EXISTS rankings;
drop FUNCTION  IF EXISTS ranking_by_period;
create OR REPLACE FUNCTION ranking_by_period(BIGINT)
returns table (
		old_total DOUBLE PRECISION,
		old_collateral_underlying DOUBLE PRECISION,
		old_profit_underlying DOUBLE PRECISION,
		old_borrowed_amount_usd DOUBLE PRECISION,
		old_twv_usd DOUBLE PRECISION,
		old_hf varchar(80),
        sid varchar(100),
		new_total DOUBLE PRECISION,
		new_collateral_underlying DOUBLE PRECISION,
		new_profit_underlying DOUBLE PRECISION,
        session_id varchar(100),
        current_block integer,

		old_collateral DOUBLE PRECISION,
		old_profit DOUBLE PRECISION,
		new_collateral DOUBLE PRECISION,
		new_profit DOUBLE PRECISION,

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
    RETURN QUERY 
	WITH cm_prices AS (SELECT cm.address credit_manager, price FROM 
		public.token_current_price tcp JOIN public.credit_managers cm ON cm.underlying_token = tcp.token WHERE price_source='chainlink')
	SELECT t1.*, t2.*, 

		price * t1.old_collateral_underlying old_collateral, price * t1.old_profit_underlying old_profit,
		price * t2.new_collateral_underlying new_collateral, price * t2.new_profit_underlying new_profit,

		(t2.new_profit_underlying-t1.old_profit_underlying) * price profit_usd, t1.old_collateral_underlying*price collateral_usd,
		(t2.new_profit_underlying-t1.old_profit_underlying) profit_underlying, t1.old_collateral_underlying collateral_underlying, 

		(case when t1.old_collateral_underlying=0 then 0 else 
		(t2.new_profit_underlying-t1.old_profit_underlying)/(t1.old_collateral_underlying) 
		end) roi FROM

        (SELECT distinct on (d.session_id) d.total_value_usd old_total,
			d.collateral_underlying old_collateral_underlying, d.profit_underlying as old_profit_underlying,
			(case when cal_total_value_bi::float8=0  then 0 else 
			(cal_borrowed_amt_with_interest_bi::float8 * total_value_usd)/ cal_total_value_bi::float8
			end) old_borrowed_amount_usd,
			(case when cal_total_value_bi::float8=0  then 0 else 
			(cal_threshold_value_bi::float8 * total_value_usd)/ cal_total_value_bi::float8
			end) old_twv_usd,
			cal_health_factor old_hf, d.session_id sid
			FROM public.debts d WHERE block_num >= (SELECT min(id) FROM public.blocks WHERE timestamp > (extract(epoch from now())::bigint - $1)) 
            order by d.session_id, block_num) t1
        JOIN (SELECT distinct on (d.session_id) d.total_value_usd new_total,
			d.collateral_underlying new_collateral_underlying, d.profit_underlying as new_profit_underlying,
			d.session_id, block_num current_block
			FROM public.debts d WHERE block_num >= (SELECT min(id) FROM public.blocks WHERE timestamp > (extract(epoch from now())::bigint - $1)) 
            order by d.session_id, block_num DESC) t2
            ON t1.sid = t2.session_id
		JOIN public.credit_sessions cs ON cs.id = t2.session_id
		LEFT JOIN cm_prices ON cm_prices.credit_manager = cs.credit_manager;
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