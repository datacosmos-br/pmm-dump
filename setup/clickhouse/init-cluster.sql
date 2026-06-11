CREATE DATABASE IF NOT EXISTS pmm;

CREATE TABLE IF NOT EXISTS pmm.metrics_local
(
    period_start DateTime,
    queryid      String,
    service_name String,
    value        Float64
) ENGINE = ReplicatedMergeTree('/clickhouse/tables/{shard}/pmm.metrics', '{replica}')
ORDER BY (period_start, queryid);

CREATE TABLE IF NOT EXISTS pmm.metrics AS pmm.metrics_local
ENGINE = Distributed('test_cluster', 'pmm', 'metrics_local', rand());

INSERT INTO pmm.metrics_local (period_start, queryid, service_name, value)
SELECT
    toDateTime('2026-06-01 00:00:00') + number * 60,
    toString(number % 100),
    'test-service',
    rand() / 4294967295
FROM numbers(1000);
