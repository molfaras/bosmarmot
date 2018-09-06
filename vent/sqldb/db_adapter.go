package sqldb

import (
	"database/sql"

	"github.com/monax/bosmarmot/vent/types"
)

// DBAdapter database access interface
type DBAdapter interface {
	Open(dbURL string) (*sql.DB, error)
	TypeMapping(sqlColumnType types.SQLColumnType) (string, error)
	CreateTableQuery(tableName string, columns []types.SQLTableColumn) string
	UpsertQuery(table types.SQLTable) types.UpsertQuery
	LastBlockIDQuery() string
	FindTableQuery(tableName string) string
	TableDefinitionQuery(tableName string) string
	AlterColumnQuery(tableName, columnName string, sqlColumnType types.SQLColumnType) string
	SelectRowQuery(tableName, fields, indexValue string) string
	SelectLogQuery() string
	InsertLogQuery() string
	ErrorEquals(err error, sqlErrorType types.SQLErrorType) bool
}
