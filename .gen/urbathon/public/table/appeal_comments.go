//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package table

import (
	"github.com/go-jet/jet/v2/postgres"
)

var AppealComments = newAppealCommentsTable("public", "appeal_comments", "")

type appealCommentsTable struct {
	postgres.Table

	// Columns
	ID       postgres.ColumnInteger
	AppealID postgres.ColumnInteger
	UserID   postgres.ColumnInteger
	Date     postgres.ColumnTimestampz
	Text     postgres.ColumnString

	AllColumns     postgres.ColumnList
	MutableColumns postgres.ColumnList
}

type AppealCommentsTable struct {
	appealCommentsTable

	EXCLUDED appealCommentsTable
}

// AS creates new AppealCommentsTable with assigned alias
func (a AppealCommentsTable) AS(alias string) *AppealCommentsTable {
	return newAppealCommentsTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new AppealCommentsTable with assigned schema name
func (a AppealCommentsTable) FromSchema(schemaName string) *AppealCommentsTable {
	return newAppealCommentsTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new AppealCommentsTable with assigned table prefix
func (a AppealCommentsTable) WithPrefix(prefix string) *AppealCommentsTable {
	return newAppealCommentsTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new AppealCommentsTable with assigned table suffix
func (a AppealCommentsTable) WithSuffix(suffix string) *AppealCommentsTable {
	return newAppealCommentsTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newAppealCommentsTable(schemaName, tableName, alias string) *AppealCommentsTable {
	return &AppealCommentsTable{
		appealCommentsTable: newAppealCommentsTableImpl(schemaName, tableName, alias),
		EXCLUDED:            newAppealCommentsTableImpl("", "excluded", ""),
	}
}

func newAppealCommentsTableImpl(schemaName, tableName, alias string) appealCommentsTable {
	var (
		IDColumn       = postgres.IntegerColumn("id")
		AppealIDColumn = postgres.IntegerColumn("appeal_id")
		UserIDColumn   = postgres.IntegerColumn("user_id")
		DateColumn     = postgres.TimestampzColumn("date")
		TextColumn     = postgres.StringColumn("text")
		allColumns     = postgres.ColumnList{IDColumn, AppealIDColumn, UserIDColumn, DateColumn, TextColumn}
		mutableColumns = postgres.ColumnList{AppealIDColumn, UserIDColumn, DateColumn, TextColumn}
	)

	return appealCommentsTable{
		Table: postgres.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		ID:       IDColumn,
		AppealID: AppealIDColumn,
		UserID:   UserIDColumn,
		Date:     DateColumn,
		Text:     TextColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}
