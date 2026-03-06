# Database Schema Overview

This document explains the main tables and their relationships in the Gearbox Third Eye database. The schema is designed to track liquidity pools, credit management, user sessions, and financial operations in the Gearbox protocol.

## Core Tables and Relationships

### 1. Blocks (`blocks`)
The foundation table that tracks blockchain blocks.
- `id` (PK): Block number
- `timestamp`: Block timestamp

**Relationships**: Referenced by most other tables for temporal tracking

### 2. Tokens (`tokens`)
Registry of all tokens in the system.
- `address` (PK): Token contract address
- `symbol`: Token symbol (e.g., USDC, DAI)
- `decimals`: Token decimal places

### 3. Pools System

#### `pools` (Main Pool Registry)
- `address` (PK): Pool contract address
- `underlying_token`: The base token for the pool (references `tokens.address`)
- `diesel_token`: LP token issued by the pool
- `is_weth`: Whether this is a WETH pool
- Pool configuration: `borrow_apy_bi`, `expected_liq_limit`, `withdraw_fee`

#### `pool_stats` (Pool Performance Metrics)
Historical statistics for each pool at each block.
- `pool` + `block_num` (Composite PK)
- **Foreign Keys**: 
  - `pool` → `pools.address`
  - `block_num` → `blocks.id`
- Metrics: `total_borrowed`, `total_profit`, `deposit_apy`, `available_liquidity`, etc.

#### `pool_ledger` (Pool Transaction History)
Records all pool-related transactions.
- `pool` + `block_num` + `log_id` (Composite PK)
- **Foreign Keys**: 
  - `pool` → `pools.address` 
  - `block_num` → `blocks.id`
- Transaction data: `event`, `user_address`, `amount`, `session_id`

### 4. Credit Management System

#### `credit_managers` (Credit Manager Registry)
- `address` (PK): Credit manager contract address
- `pool_address`: Associated pool (references `pools.address`)
- `underlying_token`: Base token for lending
- Configuration: `max_leverage`, `min_amount`, `max_amount`
- Aggregated metrics: `total_borrowed`, `total_profit`, `opened_accounts_count`

#### `credit_manager_stats` (Historical Credit Manager Metrics)
Time-series data for credit manager performance.
- **Foreign Keys**:
  - `credit_manager` → `credit_managers.address`
  - `block_num` → `blocks.id`
- Metrics mirror the aggregated fields in `credit_managers`

#### `allowed_tokens` (Credit Manager Token Permissions)
Defines which tokens are allowed as collateral for each credit manager.
- `credit_manager` + `token` + `block_num` (Composite PK)
- **Foreign Keys**:
  - `credit_manager` → `credit_managers.address`
  - `token` → `tokens.address`
- `liquiditythreshold`: Liquidation threshold for the token
- `disable_block`: When this token allowance was disabled

### 5. User Sessions and Operations

#### `credit_sessions` (Active/Closed User Sessions)
Tracks individual user borrowing sessions.
- `id` (PK): Unique session identifier
- `credit_manager`: Associated credit manager (references `credit_managers.address`)
- `borrower`: User's wallet address
- `account`: Credit account address
- Session lifecycle: `status`, `since`, `closed_at`
- Financial data: `initial_amount`, `borrowed_amount`, `balances` (JSONB)

#### `credit_session_snapshots` (Session State History)
Historical snapshots of session state at different blocks.
- **Foreign Keys**:
  - `session_id` → `credit_sessions.id`
  - `block_num` → `blocks.id`
- Financial metrics: `total_value`, `health_factor`, `collateral_usd`
- `balances` (JSONB): Token balances at snapshot time

#### `account_operations` (Transaction Details)
Detailed log of all operations performed within credit sessions.
- **Foreign Keys**:
  - `session_id` → `credit_sessions.id`
  - `block_num` → `blocks.id`
  - `main_action` → `account_operations.id` (self-reference for grouped operations)
- Operation data: `action`, `dapp`, `adapter_call`, `args` (JSONB), `transfers`

### 6. Financial Tracking

#### `debts` (Debt Calculations)
Periodic debt and health factor calculations for sessions.
- **Foreign Keys**:
  - `session_id` → `credit_sessions.id`
  - `block_num` → `blocks.id`
- Calculated metrics: `cal_health_factor`, `cal_total_value_bi`, `profit_underlying`

#### `price_feeds` (Token Price Data)
Historical token prices from oracles.
- **Foreign Keys**:
  - `token` → `tokens.address`
  - `block_num` → `blocks.id`
- Price data: `price`, `price_bi`, `price_in_usd`

#### `token_current_price` (Latest Prices)
Current/latest prices for each token by source.
- `token` + `price_source` (Composite PK)
- **Foreign Keys**: `token` → `tokens.address`

### 7. Liquidity Mining and Rewards

#### `diesel_balances` (User LP Token Holdings)
Tracks user holdings of diesel (LP) tokens.
- `diesel_sym` + `user_address` (Composite PK)
- Balance data: `balance`, `balance_bi`

#### `lm_rewards` (Liquidity Mining Rewards)
Accumulated liquidity mining rewards per user per pool.
- `pool` + `user_address` (Composite PK)
- **Foreign Keys**: `pool` → `pools.address`

#### `farm_v3` (Farming Contracts)
V3 farming contract configurations.
- `farm` (PK): Farm contract address
- **Foreign Keys**: `pool` → `pools.address`
- Farming parameters: `reward`, `period`, `end_ts`, `total_supply`

#### `user_lmdetails_v3` (User Farming Details)
User-specific farming data for V3 farms.
- `farm` + `account` (Composite PK)
- **Foreign Keys**: `farm` → `farm_v3.farm`

## Key Relationships Summary

1. **Pool Ecosystem**: `pools` ↔ `pool_stats` ↔ `pool_ledger` ↔ `credit_managers`
2. **Credit Flow**: `credit_managers` ↔ `credit_sessions` ↔ `credit_session_snapshots` ↔ `account_operations`
3. **Token Management**: `tokens` ↔ `allowed_tokens` ↔ `price_feeds`
4. **Temporal Tracking**: `blocks` is referenced by all time-series tables
5. **User Activity**: Users connect through `credit_sessions`, `diesel_balances`, and `lm_rewards`

## Data Flow
1. **Pool Operations**: Users interact with pools → recorded in `pool_ledger` → aggregated in `pool_stats`
2. **Credit Sessions**: Users open sessions → tracked in `credit_sessions` → operations in `account_operations` → snapshots in `credit_session_snapshots`
3. **Financial Tracking**: Prices from `price_feeds` → debt calculations in `debts` → health factors and liquidations
4. **Rewards**: LP token holdings in `diesel_balances` → farming participation in `user_lmdetails_v3` → rewards in `lm_rewards`