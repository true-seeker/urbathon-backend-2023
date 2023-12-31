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

var AppealPhotos = newAppealPhotosTable("public", "appeal_photos", "")

type appealPhotosTable struct {
	postgres.Table

	// Columns
	ID       postgres.ColumnInteger
	AppealID postgres.ColumnInteger
	URL      postgres.ColumnString

	AllColumns     postgres.ColumnList
	MutableColumns postgres.ColumnList
}

type AppealPhotosTable struct {
	appealPhotosTable

	EXCLUDED appealPhotosTable
}

// AS creates new AppealPhotosTable with assigned alias
func (a AppealPhotosTable) AS(alias string) *AppealPhotosTable {
	return newAppealPhotosTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new AppealPhotosTable with assigned schema name
func (a AppealPhotosTable) FromSchema(schemaName string) *AppealPhotosTable {
	return newAppealPhotosTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new AppealPhotosTable with assigned table prefix
func (a AppealPhotosTable) WithPrefix(prefix string) *AppealPhotosTable {
	return newAppealPhotosTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new AppealPhotosTable with assigned table suffix
func (a AppealPhotosTable) WithSuffix(suffix string) *AppealPhotosTable {
	return newAppealPhotosTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newAppealPhotosTable(schemaName, tableName, alias string) *AppealPhotosTable {
	return &AppealPhotosTable{
		appealPhotosTable: newAppealPhotosTableImpl(schemaName, tableName, alias),
		EXCLUDED:          newAppealPhotosTableImpl("", "excluded", ""),
	}
}

func newAppealPhotosTableImpl(schemaName, tableName, alias string) appealPhotosTable {
	var (
		IDColumn       = postgres.IntegerColumn("id")
		AppealIDColumn = postgres.IntegerColumn("appeal_id")
		URLColumn      = postgres.StringColumn("url")
		allColumns     = postgres.ColumnList{IDColumn, AppealIDColumn, URLColumn}
		mutableColumns = postgres.ColumnList{AppealIDColumn, URLColumn}
	)

	return appealPhotosTable{
		Table: postgres.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		ID:       IDColumn,
		AppealID: AppealIDColumn,
		URL:      URLColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}
