


----区块表
create table blocks
(
    timestamp string,
    number bigint,
    hash string,
    parent_hash string,
    nonce string,
    sha3_uncles string,
    logs_bloom string,
    transactions_root string,
    state_root string,
    receipts_root string,
    miner symbol,
    difficulty long,
    total_difficulty long,
    size long,
    extra_data string,
    gas_limit long,
    gas_used long,
    transaction_count long
);


---交易表
create table token_transfers
(
  token_address symbol,
  from_address symbol,
  to_address symbol,
  value float,
  transaction_hash string,
  log_index long,
  block_timestamp string,
  block_number long,
  block_hash string
);

---ETH交易

---代币交易