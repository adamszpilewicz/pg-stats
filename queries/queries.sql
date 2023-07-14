-- Query to check overall database statistics:
SELECT datname AS database_name,
       numbackends AS number_of_connections,
       xact_commit AS total_transactions_committed,
       xact_rollback AS total_transactions_rolledback,
       blks_read AS blocks_read,
       blks_hit AS blocks_hit,
       tup_returned AS tuples_returned,
       tup_fetched AS tuples_fetched,
       tup_inserted AS tuples_inserted,
       tup_updated AS tuples_updated,
       tup_deleted AS tuples_deleted
FROM pg_stat_database;

-- Query to view long running queries:
SELECT pid,
       now() - pg_stat_activity.query_start AS duration,
       query,
       state
FROM pg_stat_activity
WHERE (now() - pg_stat_activity.query_start) > interval '5 minutes';

-- Query to check size of all databases:
SELECT pg_database.datname AS database_name,
       pg_size_pretty(pg_database_size(pg_database.datname)) AS size
FROM pg_database;

-- Query to check table disk size:
SELECT relname AS table_name,
       pg_size_pretty(pg_total_relation_size(C.oid)) AS total_size
FROM pg_class C
         LEFT JOIN pg_namespace N ON (N.oid = C.relnamespace)
WHERE nspname NOT IN ('pg_catalog', 'information_schema')
  AND C.relkind <> 'i'
  AND nspname !~ '^pg_toast'
ORDER BY pg_total_relation_size(C.oid) DESC;
