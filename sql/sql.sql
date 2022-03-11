CREATE DATABASE IF NOT EXISTS companion;

CREATE TABLE accesses (
    id SERIAL
    , uuid VARCHAR(64) NOT NULL
    , referrer VARCHAR(256) NOT NULL
    , cookie TEXT
    , user_agent TEXT NOT NULL
    , created_at TIMESTAMP WITHOUT TIME ZONE DEFAULT (NOW() AT TIME ZONE 'UTC')
    , os VARCHAR(64)
    , browser VARCHAR(64)
    , screen JSON
    , navigator JSON
) PARTITION BY RANGE (created_at);

CREATE TABLE IF NOT EXISTS accesses_20220101 PARTITION OF accesses
    FOR VALUES FROM ('2022-01-01') TO ('2022-01-02');