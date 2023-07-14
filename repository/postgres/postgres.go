package postgres

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4/pgxpool"
	"stats-pg/repository"
)

type PostgresRepository struct {
	conn *pgxpool.Pool
}

// NewPostgresRepository creates a new instance of PostgresRepository.
func NewPostgresRepository(connString string) (*PostgresRepository, error) {
	conn, err := pgxpool.Connect(context.Background(), connString)
	if err != nil {
		return nil, err
	}
	return &PostgresRepository{
		conn: conn,
	}, nil
}

// GetStats returns the stats
func (r *PostgresRepository) GetStats() ([]repository.Stats, error) {
	rows, err := r.conn.Query(context.Background(), queryStats)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var statsList []repository.Stats
	for rows.Next() {
		var stats repository.Stats
		err = rows.Scan(&stats.DatabaseName, &stats.NumberOfConnections, &stats.TotalTransactionsCommitted,
			&stats.TotalTransactionsRolledback, &stats.BlocksRead, &stats.BlocksHit, &stats.TuplesReturned,
			&stats.TuplesFetched, &stats.TuplesInserted, &stats.TuplesUpdated, &stats.TuplesDeleted)
		if err != nil {
			return nil, err
		}
		statsList = append(statsList, stats)
	}

	return statsList, nil
}

func (r *PostgresRepository) Close() {
	r.conn.Close()
}

func (r *PostgresRepository) GetStatsByName(ctx *gin.Context, dbName string) (repository.Stats, error) {
	var stats repository.Stats
	//var ctx = context.Background()
	err := r.conn.QueryRow(ctx, queryStatsName, dbName).Scan(&stats.DatabaseName, &stats.NumberOfConnections,
		&stats.TotalTransactionsCommitted, &stats.TotalTransactionsRolledback, &stats.BlocksRead,
		&stats.BlocksHit, &stats.TuplesReturned, &stats.TuplesFetched, &stats.TuplesInserted,
		&stats.TuplesUpdated, &stats.TuplesDeleted)
	if err != nil {
		return stats, err
	}
	return stats, nil
}
