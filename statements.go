package gosql

import (

)

type Statement interface{}

type CreateTable struct {
	Name string
}
