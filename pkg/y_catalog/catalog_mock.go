package catalog

import "github.com/apache/arrow/go/v12/arrow"

type MockCatalog struct {
}

func (c *MockCatalog) CreateTable(tableDef *TableDef) error {
	//TODO implement me
	panic("implement me")
}

func (c *MockCatalog) GetTable(tableName string) *TableDef {
	return &TableDef{
		TableName: tableName,
		Columns: []ColumnDef{
			{
				ColumnName: "c1",
				ColumnType: arrow.PrimitiveTypes.Int64,
			},
			{
				ColumnName: "c2",
				ColumnType: arrow.PrimitiveTypes.Int64,
			},
		},
	}
}

func (c *MockCatalog) AddColumn(tableName string, columnDef *ColumnDef) error {
	//TODO implement me
	panic("implement me")
}

func (c *MockCatalog) GetColumn(tableName string, columnName string) *ColumnDef {
	//TODO implement me
	panic("implement me")
}
