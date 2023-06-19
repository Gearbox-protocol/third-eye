alter table credit_sessions add close_transfers jsonb;
update credit_sessions set close_transfers=transfers FROM 
    (select id, jsonb_object(ARRAY_AGG(k), ARRAY_AGG(v->>'F')) transfers FROM 
        (select id, k,v from credit_sessions, jsonb_each(balances) arr(k,v) where version=2 and status=1) tmp group by id) tmp2
        where tmp2.id=credit_sessions.id;