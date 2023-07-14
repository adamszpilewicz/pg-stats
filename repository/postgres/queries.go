package postgres

var (
	queryStats = `SELECT datname AS database_name,
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
		FROM pg_stat_database`

	queryStatsName = `SELECT datname AS database_name,
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
		FROM pg_stat_database
		WHERE datname = $1`
)
