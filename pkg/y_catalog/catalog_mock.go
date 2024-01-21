package catalog

type MockCatalog struct {
	tables       []*TableDef
	tablesByID   map[uint32]*TableDef
	tablesByName map[string]*TableDef

	maxTableID uint32 // The maxTableID variable is used to assign unique ids to new tables as they are created.
}

func NewMockCatalog() Catalog {
	ctlg := MockCatalog{
		tablesByID:   make(map[uint32]*TableDef),
		tablesByName: make(map[string]*TableDef),
	}

	pgTypeTable := &TableDef{
		catalog: &ctlg,
		name:    "pg_type",
		cols: []*ColumnDef{
			{
				colName: "oid",
				colType: VarcharType,
				maxLen:  10,
			},
			{
				colName: "typbasetype",
				colType: VarcharType,
				maxLen:  10,
			},
			{
				colName: "typname",
				colType: VarcharType,
				maxLen:  50,
			},
		},
	}

	pgTypeTable.colsByName = make(map[string]*ColumnDef, len(pgTypeTable.cols))

	for _, col := range pgTypeTable.cols {
		pgTypeTable.colsByName[col.colName] = col
	}

	pgTypeTable.indexes = []*Index{
		{
			unique: true,
			cols: []*ColumnDef{
				pgTypeTable.colsByName["oid"],
			},
			colsByID: map[uint32]*ColumnDef{
				0: pgTypeTable.colsByName["oid"],
			},
		},
	}

	pgTypeTable.primaryIndex = pgTypeTable.indexes[0]
	ctlg.tablesByName[pgTypeTable.name] = pgTypeTable

	return &ctlg
}

func (c *MockCatalog) CreateTable(tableDef *TableDef) error {
	//TODO implement me
	panic("implement me")
}

func (c *MockCatalog) GetTable(tableName string) *TableDef {
	return c.tablesByName[tableName]
}

func (c *MockCatalog) AddColumn(tableName string, columnDef *ColumnDef) error {
	//TODO implement me
	panic("implement me")
}

func (c *MockCatalog) GetColumn(tableName string, columnName string) *ColumnDef {
	//TODO implement me
	panic("implement me")
}
