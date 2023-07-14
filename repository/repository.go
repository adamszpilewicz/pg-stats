package repository

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
)

// Stats represents the stats for a database
type Stats struct {
	DatabaseName                StringValue `json:"database_name"`
	NumberOfConnections         IntValue    `json:"number_of_connections"`
	TotalTransactionsCommitted  IntValue    `json:"total_transactions_committed"`
	TotalTransactionsRolledback IntValue    `json:"total_transactions_rolledback"`
	BlocksRead                  IntValue    `json:"blocks_read"`
	BlocksHit                   IntValue    `json:"blocks_hit"`
	TuplesReturned              IntValue    `json:"tuples_returned"`
	TuplesFetched               IntValue    `json:"tuples_fetched"`
	TuplesInserted              IntValue    `json:"tuples_inserted"`
	TuplesUpdated               IntValue    `json:"tuples_updated"`
	TuplesDeleted               IntValue    `json:"tuples_deleted"`
}

type StringValue struct {
	sql.NullString
}

func (v StringValue) MarshalJSON() ([]byte, error) {
	if v.Valid {
		return json.Marshal(v.String)
	}
	return json.Marshal(nil)
}

type IntValue struct {
	sql.NullInt64
}

func (v IntValue) MarshalJSON() ([]byte, error) {
	if v.Valid {
		return json.Marshal(v.Int64)
	}
	return json.Marshal(nil)
}

type Repository interface {
	// GetStats returns the stats for all databases
	GetStats() ([]Stats, error)
	// GetStatsByName returns the stats for the given database name
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
