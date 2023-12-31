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

var AppealCategories = newAppealCategoriesTable("public", "appeal_categories", "")

type appealCategoriesTable struct {
	postgres.Table

	// Columns
	ID    postgres.ColumnInteger
	Title postgres.ColumnString

	AllColumns     postgres.ColumnList
	MutableColumns postgres.ColumnList
}

type AppealCategoriesTable struct {
	appealCategoriesTable

	EXCLUDED appealCategoriesTable
}

// AS creates new AppealCategoriesTable with assigned alias
func (a AppealCategoriesTable) AS(alias string) *AppealCategoriesTable {
	return newAppealCategoriesTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new AppealCategoriesTable with assigned schema name
func (a AppealCategoriesTable) FromSchema(schemaName string) *AppealCategoriesTable {
	return newAppealCategoriesTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new AppealCategoriesTable with assigned table prefix
func (a AppealCategoriesTable) WithPrefix(prefix string) *AppealCategoriesTable {
	return newAppealCategoriesTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new AppealCategoriesTable with assigned table suffix
func (a AppealCategoriesTable) WithSuffix(suffix string) *AppealCategoriesTable {
	return newAppealCategoriesTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newAppealCategoriesTable(schemaName, tableName, alias string) *AppealCategoriesTable {
	return &AppealCategoriesTable{
		appealCategoriesTable: newAppealCategoriesTableImpl(schemaName, tableName, alias),
		EXCLUDED:              newAppealCategoriesTableImpl("", "excluded", ""),
	}
}

func newAppealCategoriesTableImpl(schemaName, tableName, alias string) appealCategoriesTable {
	var (
		IDColumn       = postgres.IntegerColumn("id")
		TitleColumn    = postgres.StringColumn("title")
		allColumns     = postgres.ColumnList{IDColumn, TitleColumn}
		mutableColumns = postgres.ColumnList{TitleColumn}
	)

	return appealCategoriesTable{
		Table: postgres.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		ID:    IDColumn,
		Title: TitleColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}
