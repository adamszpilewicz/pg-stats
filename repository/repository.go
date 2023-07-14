package repository

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
)

// Stats is an example structure for holding table stats.
// You can define this based on your actual requirement.
type Stats struct {
	DatabaseName                sql.NullString
	NumberOfConnections         sql.NullInt64
	TotalTransactionsCommitted  sql.NullInt64
	TotalTransactionsRolledback sql.NullInt64
	BlocksRead                  sql.NullInt64
	BlocksHit                   sql.NullInt64
	TuplesReturned              sql.NullInt64
	TuplesFetched               sql.NullInt64
	TuplesInserted              sql.NullInt64
	TuplesUpdated               sql.NullInt64
	TuplesDeleted               sql.NullInt64
}

type Repository interface {
	// GetStats returns the stats for the given table
	GetStats() ([]Stats, error)
	// GetStatsByName returns the stats for the given table
	GetStatsByName(ctx *gin.Context, dbName string) (Stats, error)
}

func (s Stats) String() string {
	return fmt.Sprintf(
		"DatabaseName: %v, NumberOfConnections: %v, TotalTransactionsCommitted: %v, "+
			"TotalTransactionsRolledback: %v, BlocksRead: %v, BlocksHit: %v, TuplesReturned: %v, "+
			"TuplesFetched: %v, TuplesInserted: %v, TuplesUpdated: %v, TuplesDeleted: %v\n",
		s.DatabaseName.String,
		s.NumberOfConnections.Int64,
		s.TotalTransactionsCommitted.Int64,
		s.TotalTransactionsRolledback.Int64,
		s.BlocksRead.Int64,
		s.BlocksHit.Int64,
		s.TuplesReturned.Int64,
		s.TuplesFetched.Int64,
		s.TuplesInserted.Int64,
		s.TuplesUpdated.Int64,
		s.TuplesDeleted.Int64,
	)
}
