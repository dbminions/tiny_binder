package catalog

import "github.com/apache/arrow/go/v12/arrow"

type Catalog interface {
	CreateTable(tableDef *TableDef) error // TableDef related
	GetTable(tableName string) *TableDef

	AddColumn(tableName string, columnDef *ColumnDef) error // ColumnDef related
	GetColumn(tableName string, columnName string) *ColumnDef
}

var _ Catalog = &MockCatalog{}

type TableDef struct {
	TableName string
	Columns   []ColumnDef
}

type ColumnDef struct {
	ColumnName string
	ColumnType arrow.DataType
}
