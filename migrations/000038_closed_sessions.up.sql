create table closed_trading_sessions (
    session_id varchar(80),
    data jsonb,
    PRIMARY KEY (session_id));