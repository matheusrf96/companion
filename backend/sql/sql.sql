CREATE DATABASE IF NOT EXISTS companion;

CREATE TABLE accesses (
    id SERIAL
    , uuid VARCHAR(64) NOT NULL
    , source_id INT
    , ecommerce_id INT
    , utm_source VARCHAR(64)
    , utm_medium VARCHAR(64)
    , tags VARCHAR(32)[]
    , created_at TIMESTAMP WITHOUT TIME ZONE DEFAULT (NOW() AT TIME ZONE 'UTC')
    , referrer VARCHAR(256) NOT NULL
    , cookie TEXT
    , user_agent TEXT NOT NULL
    , query VARCHAR(256)
    , device VARCHAR(64)
    , os VARCHAR(64)
    , browser VARCHAR(64)
    , screen JSON
    , navigator JSON
) PARTITION BY RANGE (created_at);

CREATE TABLE IF NOT EXISTS accesses_20220101 PARTITION OF accesses
    FOR VALUES FROM ('2022-01-01') TO ('2022-01-02');