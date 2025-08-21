-- Query to get token balance for credit sessions at the end of each day
-- Token: 0x7a4EffD87C2f3C55CA251080b1343b605f327E3a
-- Credit Managers: 0xD665774088c7936B65be0cBCF83AEBe87CB10DE7, 0x4bF7481fDf7b67A8206254Badc15480A55bB25ab

WITH daily_end_blocks AS (
    -- Get the last block of each day since block 21944825
    SELECT 
        DATE(to_timestamp(timestamp)) as day,
        MAX(id) as last_block_of_day
    FROM blocks 
    WHERE id >= 21944825
    GROUP BY DATE(to_timestamp(timestamp))
),
target_credit_sessions AS (
    -- Get credit sessions for the specified credit managers
    SELECT id, since, closed_at
    FROM credit_sessions 
    WHERE credit_manager IN ('0xD665774088c7936B65be0cBCF83AEBe87CB10DE7', '0x4bF7481fDf7b67A8206254Badc15480A55bB25ab')
),
valid_snapshots AS (
    -- Get snapshots that exist at the end of each day for our target sessions
    SELECT DISTINCT ON (css.session_id, deb.day)
        deb.day,
        deb.last_block_of_day,
        css.session_id,
        css.block_num,
        css.balances->'0x7a4EffD87C2f3C55CA251080b1343b605f327E3a'->>'F' as token_balance
    FROM daily_end_blocks deb
    JOIN credit_session_snapshots css ON css.block_num <= deb.last_block_of_day
    JOIN target_credit_sessions tcs ON tcs.id = css.session_id
    WHERE css.balances->'0x7a4EffD87C2f3C55CA251080b1343b605f327E3a'->>'F' IS NOT NULL
        AND tcs.since <= deb.last_block_of_day  -- Session must have started by end of day
        AND (tcs.closed_at = 0 OR tcs.closed_at > deb.last_block_of_day)  -- Session must be open at end of day
    ORDER BY css.session_id, deb.day, css.block_num DESC
)
SELECT 
    day,
    last_block_of_day,
    session_id,
    block_num as snapshot_block,
    token_balance
FROM valid_snapshots
WHERE token_balance IS NOT NULL
    AND token_balance != '0'
ORDER BY day DESC, session_id;

-- Alternative query if you want aggregated data per day
/*
SELECT 
    day,
    last_block_of_day,
    COUNT(DISTINCT session_id) as active_sessions,
    SUM(CAST(token_balance AS NUMERIC)) / POWER(10, 18) as total_token_balance,
    AVG(CAST(token_balance AS NUMERIC)) / POWER(10, 18) as avg_token_balance
FROM valid_snapshots
WHERE token_balance IS NOT NULL
    AND token_balance != '0'
GROUP BY day, last_block_of_day
ORDER BY day DESC;
*/