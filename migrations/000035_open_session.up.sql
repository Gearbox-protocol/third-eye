update account_operations  SET args = args || jsonb_build_object('userFunds',args->>'amount') where action like 'OpenCreditAccount%';