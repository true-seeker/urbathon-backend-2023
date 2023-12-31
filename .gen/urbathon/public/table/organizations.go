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

var Organizations = newOrganizationsTable("public", "organizations", "")

type organizationsTable struct {
	postgres.Table

	// Columns
	ID      postgres.ColumnInteger
	Name    postgres.ColumnString
	Inn     postgres.ColumnString
	Address postgres.ColumnString
	Phone   postgres.ColumnString

	AllColumns     postgres.ColumnList
	MutableColumns postgres.ColumnList
}

type OrganizationsTable struct {
	organizationsTable

	EXCLUDED organizationsTable
}

// AS creates new OrganizationsTable with assigned alias
func (a OrganizationsTable) AS(alias string) *OrganizationsTable {
	return newOrganizationsTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new OrganizationsTable with assigned schema name
func (a OrganizationsTable) FromSchema(schemaName string) *OrganizationsTable {
	return newOrganizationsTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new OrganizationsTable with assigned table prefix
func (a OrganizationsTable) WithPrefix(prefix string) *OrganizationsTable {
	return newOrganizationsTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new OrganizationsTable with assigned table suffix
func (a OrganizationsTable) WithSuffix(suffix string) *OrganizationsTable {
	return newOrganizationsTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newOrganizationsTable(schemaName, tableName, alias string) *OrganizationsTable {
	return &OrganizationsTable{
		organizationsTable: newOrganizationsTableImpl(schemaName, tableName, alias),
		EXCLUDED:           newOrganizationsTableImpl("", "excluded", ""),
	}
}

func newOrganizationsTableImpl(schemaName, tableName, alias string) organizationsTable {
	var (
		IDColumn       = postgres.IntegerColumn("id")
		NameColumn     = postgres.StringColumn("name")
		InnColumn      = postgres.StringColumn("inn")
		AddressColumn  = postgres.StringColumn("address")
		PhoneColumn    = postgres.StringColumn("phone")
		allColumns     = postgres.ColumnList{IDColumn, NameColumn, InnColumn, AddressColumn, PhoneColumn}
		mutableColumns = postgres.ColumnList{NameColumn, InnColumn, AddressColumn, PhoneColumn}
	)

	return organizationsTable{
		Table: postgres.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		ID:      IDColumn,
		Name:    NameColumn,
		Inn:     InnColumn,
		Address: AddressColumn,
		Phone:   PhoneColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}
