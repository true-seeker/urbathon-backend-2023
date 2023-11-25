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

var OrganizationAppealCategory = newOrganizationAppealCategoryTable("public", "organization_appeal_category", "")

type organizationAppealCategoryTable struct {
	postgres.Table

	// Columns
	ID               postgres.ColumnInteger
	OrganizationID   postgres.ColumnInteger
	AppealCategoryID postgres.ColumnInteger

	AllColumns     postgres.ColumnList
	MutableColumns postgres.ColumnList
}

type OrganizationAppealCategoryTable struct {
	organizationAppealCategoryTable

	EXCLUDED organizationAppealCategoryTable
}

// AS creates new OrganizationAppealCategoryTable with assigned alias
func (a OrganizationAppealCategoryTable) AS(alias string) *OrganizationAppealCategoryTable {
	return newOrganizationAppealCategoryTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new OrganizationAppealCategoryTable with assigned schema name
func (a OrganizationAppealCategoryTable) FromSchema(schemaName string) *OrganizationAppealCategoryTable {
	return newOrganizationAppealCategoryTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new OrganizationAppealCategoryTable with assigned table prefix
func (a OrganizationAppealCategoryTable) WithPrefix(prefix string) *OrganizationAppealCategoryTable {
	return newOrganizationAppealCategoryTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new OrganizationAppealCategoryTable with assigned table suffix
func (a OrganizationAppealCategoryTable) WithSuffix(suffix string) *OrganizationAppealCategoryTable {
	return newOrganizationAppealCategoryTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newOrganizationAppealCategoryTable(schemaName, tableName, alias string) *OrganizationAppealCategoryTable {
	return &OrganizationAppealCategoryTable{
		organizationAppealCategoryTable: newOrganizationAppealCategoryTableImpl(schemaName, tableName, alias),
		EXCLUDED:                        newOrganizationAppealCategoryTableImpl("", "excluded", ""),
	}
}

func newOrganizationAppealCategoryTableImpl(schemaName, tableName, alias string) organizationAppealCategoryTable {
	var (
		IDColumn               = postgres.IntegerColumn("id")
		OrganizationIDColumn   = postgres.IntegerColumn("organization_id")
		AppealCategoryIDColumn = postgres.IntegerColumn("appeal_category_id")
		allColumns             = postgres.ColumnList{IDColumn, OrganizationIDColumn, AppealCategoryIDColumn}
		mutableColumns         = postgres.ColumnList{OrganizationIDColumn, AppealCategoryIDColumn}
	)

	return organizationAppealCategoryTable{
		Table: postgres.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		ID:               IDColumn,
		OrganizationID:   OrganizationIDColumn,
		AppealCategoryID: AppealCategoryIDColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}
