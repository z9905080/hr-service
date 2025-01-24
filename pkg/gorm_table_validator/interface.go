package gorm_table_validator

// InfPo is an interface that defines the method to get the table name and convert to entity.
// T is the entity type.
type InfPo[T any] interface {
	TableName() string
	ToEntity() T
}

// InfUpdatePo is an interface that defines the method to update the table.
// T is the entity type.
type InfUpdatePo[T any] interface {
	InfPo[T]
	ToUpdateMap() map[string]interface{}
}

// InfListPo is an interface that defines the method to get the sort key list.
// T is the entity type.
type InfListPo[T any] interface {
	InfPo[T]
	GetSortKeyList() []string
}
