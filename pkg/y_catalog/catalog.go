package catalog

type Catalog interface {
	CreateTable(tableDef *TableDef) error // TableDef related
	GetTable(tableName string) *TableDef

	AddColumn(tableName string, columnDef *ColumnDef) error // ColumnDef related
	GetColumn(tableName string, columnName string) *ColumnDef
}

var _ Catalog = &MockCatalog{}

type TableDef struct {
	catalog    Catalog
	id         uint32
	name       string
	cols       []*ColumnDef
	colsByID   map[uint32]*ColumnDef
	colsByName map[string]*ColumnDef

	indexes        []*Index
	indexesByName  map[string]*Index
	indexesByColID map[uint32][]*Index

	primaryIndex    *Index
	autoIncrementPK bool
	maxPK           int64

	maxColID   uint32
	maxIndexID uint32
}

type Index struct {
	table    *TableDef
	id       uint32
	unique   bool
	cols     []*ColumnDef
	colsByID map[uint32]*ColumnDef
}

type ColumnDef struct {
	table         *TableDef
	id            uint32
	colName       string
	colType       SQLValueType
	maxLen        int
	autoIncrement bool
	notNull       bool
}

type SQLValueType = string

const (
	IntegerType   SQLValueType = "INTEGER"
	BooleanType   SQLValueType = "BOOLEAN"
	VarcharType   SQLValueType = "VARCHAR"
	UUIDType      SQLValueType = "UUID"
	BLOBType      SQLValueType = "BLOB"
	Float64Type   SQLValueType = "FLOAT"
	TimestampType SQLValueType = "TIMESTAMP"
	AnyType       SQLValueType = "ANY"
)

func IsNumericType(t SQLValueType) bool {
	return t == IntegerType || t == Float64Type
}
