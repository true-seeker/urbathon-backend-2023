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

var Tko = newTkoTable("public", "tko", "")

type tkoTable struct {
	postgres.Table

	// Columns
	ID        postgres.ColumnInteger
	Address   postgres.ColumnString
	Latitude  postgres.ColumnFloat
	Longitude postgres.ColumnFloat
	Type      postgres.ColumnString

	AllColumns     postgres.ColumnList
	MutableColumns postgres.ColumnList
}

type TkoTable struct {
	tkoTable

	EXCLUDED tkoTable
}

// AS creates new TkoTable with assigned alias
func (a TkoTable) AS(alias string) *TkoTable {
	return newTkoTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new TkoTable with assigned schema name
func (a TkoTable) FromSchema(schemaName string) *TkoTable {
	return newTkoTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new TkoTable with assigned table prefix
func (a TkoTable) WithPrefix(prefix string) *TkoTable {
	return newTkoTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new TkoTable with assigned table suffix
func (a TkoTable) WithSuffix(suffix string) *TkoTable {
	return newTkoTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newTkoTable(schemaName, tableName, alias string) *TkoTable {
	return &TkoTable{
		tkoTable: newTkoTableImpl(schemaName, tableName, alias),
		EXCLUDED: newTkoTableImpl("", "excluded", ""),
	}
}

func newTkoTableImpl(schemaName, tableName, alias string) tkoTable {
	var (
		IDColumn        = postgres.IntegerColumn("id")
		AddressColumn   = postgres.StringColumn("address")
		LatitudeColumn  = postgres.FloatColumn("latitude")
		LongitudeColumn = postgres.FloatColumn("longitude")
		TypeColumn      = postgres.StringColumn("type")
		allColumns      = postgres.ColumnList{IDColumn, AddressColumn, LatitudeColumn, LongitudeColumn, TypeColumn}
		mutableColumns  = postgres.ColumnList{AddressColumn, LatitudeColumn, LongitudeColumn, TypeColumn}
	)

	return tkoTable{
		Table: postgres.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		ID:        IDColumn,
		Address:   AddressColumn,
		Latitude:  LatitudeColumn,
		Longitude: LongitudeColumn,
		Type:      TypeColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}
