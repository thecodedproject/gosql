package gosql

import (

)

type Statement interface{}

type CreateTable struct {
	Name string
	Fields []Field
}

type Field struct {
	Name string
	Type Type

	// TODO think of a nicer interface for these attributes
	PrimaryKey bool
	AutoIncrement bool
	NotNull bool
}
