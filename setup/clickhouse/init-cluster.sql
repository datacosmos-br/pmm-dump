CREATE DATABASE IF NOT EXISTS pmm;

-- Plain MergeTree per shard: a single-replica test cluster does not need
-- ReplicatedMergeTree (which would require a ClickHouse Keeper/ZooKeeper
-- quorum). pmm-dump reads through the Distributed table below, so the
-- multi-shard read path is still exercised end to end.
CREATE TABLE IF NOT EXISTS pmm.metrics_local
(
    period_start DateTime,
    queryid      String,
    service_name String,
    value        Float64
) ENGINE = MergeTree()
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
