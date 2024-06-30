- update field in json body
```
update account_operations  SET args = args || jsonb_build_object('userFunds',args->>'amount') where action like 'OpenCreditAccount%';
```

- serialize json into key/value and build again.
```
update credit_sessions set close_transfers=transfers FROM 
    (select id, jsonb_object(ARRAY_AGG(k), ARRAY_AGG(v->>'F')) transfers FROM 
        (select id, k,v from credit_sessions, jsonb_each(balances) arr(k,v) where version=2 and status=1) tmp group by id) tmp2
        where tmp2.id=credit_sessions.id;
```


```
update last_snaps  SET state = state::jsonb || jsonb_build_object('HrlyRateStartTs',1719493200) ;
```