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

var PollOptions = newPollOptionsTable("public", "poll_options", "")

type pollOptionsTable struct {
	postgres.Table

	// Columns
	ID     postgres.ColumnInteger
	PollID postgres.ColumnInteger
	Title  postgres.ColumnString

	AllColumns     postgres.ColumnList
	MutableColumns postgres.ColumnList
}

type PollOptionsTable struct {
	pollOptionsTable

	EXCLUDED pollOptionsTable
}

// AS creates new PollOptionsTable with assigned alias
func (a PollOptionsTable) AS(alias string) *PollOptionsTable {
	return newPollOptionsTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new PollOptionsTable with assigned schema name
func (a PollOptionsTable) FromSchema(schemaName string) *PollOptionsTable {
	return newPollOptionsTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new PollOptionsTable with assigned table prefix
func (a PollOptionsTable) WithPrefix(prefix string) *PollOptionsTable {
	return newPollOptionsTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new PollOptionsTable with assigned table suffix
func (a PollOptionsTable) WithSuffix(suffix string) *PollOptionsTable {
	return newPollOptionsTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newPollOptionsTable(schemaName, tableName, alias string) *PollOptionsTable {
	return &PollOptionsTable{
		pollOptionsTable: newPollOptionsTableImpl(schemaName, tableName, alias),
		EXCLUDED:         newPollOptionsTableImpl("", "excluded", ""),
	}
}

func newPollOptionsTableImpl(schemaName, tableName, alias string) pollOptionsTable {
	var (
		IDColumn       = postgres.IntegerColumn("id")
		PollIDColumn   = postgres.IntegerColumn("poll_id")
		TitleColumn    = postgres.StringColumn("title")
		allColumns     = postgres.ColumnList{IDColumn, PollIDColumn, TitleColumn}
		mutableColumns = postgres.ColumnList{PollIDColumn, TitleColumn}
	)

	return pollOptionsTable{
		Table: postgres.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		ID:     IDColumn,
		PollID: PollIDColumn,
		Title:  TitleColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}
