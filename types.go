package gosql

import (
	"strconv"
)

type Type interface{
	FullType() (string, error)
}

type TypeBigInt struct{}

func (t TypeBigInt) FullType() (string, error) {
	return "bigint", nil
}

type TypeBit struct{}

func (t TypeBit) FullType() (string, error) {
	return "bit", nil
}

type TypeChar struct{
	N int64
}

func (t TypeChar) FullType() (string, error) {
	return "char(" + strconv.FormatInt(t.N,10) + ")", nil
}

type TypeDateTime struct{}

func (t TypeDateTime) FullType() (string, error) {
	return "datetime", nil
}

type TypeDecimal struct{
	P int64
	S int64
}

func (t TypeDecimal) FullType() (string, error) {
	return "decimcal(" + strconv.FormatInt(t.P,10) + "," + strconv.FormatInt(t.S,10) + ")", nil
}

type TypeDouble struct{}

func (t TypeDouble) FullType() (string, error) {
	return "double", nil
}

type TypeFloat struct{}

func (t TypeFloat) FullType() (string, error) {
	return "float", nil
}

type TypeInt struct{}

func (t TypeInt) FullType() (string, error) {
	return "int", nil
}

type TypeSmallInt struct{}

func (t TypeSmallInt) FullType() (string, error) {
	return "smallint", nil
}

type TypeTinyInt struct{}

func (t TypeTinyInt) FullType() (string, error) {
	return "tinyint", nil
}

type TypeVarChar struct{
	N int64
}

func (t TypeVarChar) FullType() (string, error) {
	return "varchar(" + strconv.FormatInt(t.N,10) + ")", nil
}
