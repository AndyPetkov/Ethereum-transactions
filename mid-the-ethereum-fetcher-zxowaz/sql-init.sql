CREATE TABLE IF NOT EXISTS transactions (
    transactionHash BYTEA PRIMARY KEY,
    transactionStatus integer,
    blockHash bytea,
    blockNumber integer,
    to_ bytea,
    from_ bytea,
    contractAddress bytea,
    logsCount integer,
    input bytea,
    value_ bytea
);

