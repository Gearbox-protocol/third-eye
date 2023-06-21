alter table credit_sessions add close_transfers jsonb;
update credit_sessions set close_transfers=transfers FROM 
    (select id, jsonb_object(ARRAY_AGG(k), ARRAY_AGG(v->>'F')) transfers FROM 
        (select id, k,v from credit_sessions, jsonb_each(balances) arr(k,v) where version=2 and status=1) tmp group by id) tmp2
        where tmp2.id=credit_sessions.id;

update credit_sessions cs set  balances=css.balances from  credit_session_snapshots css where status=1 and version = 2 and  css.block_num = cs.closed_at-1 and cs.id= css.session_id;