package gosql_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/thecodedproject/gosql"
)

func TestParseType(t *testing.T) {

	testCases := []struct{
		Type string
		Expected any
		ExpectedErr string
	}{
		{
			Type: "some unknown type",
			ExpectedErr: "parse error",
		},
		{
			Type: "char(",
			ExpectedErr: "parse error",
		},
		{
			Type: "datetime)(",
			ExpectedErr: "parse error",
		},
		{
			Type: "datetime(hello)",
			ExpectedErr: "parse error",
		},
		{
			Type: "char(1)fdskl",
			ExpectedErr: "parse error",
		},
		{
			Type: "bigint",
			Expected: gosql.TypeBigInt{},
		},
		{
			Type: "bit",
			Expected: gosql.TypeBit{},
		},
		{
			Type: "char",
			Expected: gosql.TypeChar{},
		},
		{
			Type: "char(123)",
			Expected: gosql.TypeChar{123},
		},
		{
			Type: "char(255)",
			Expected: gosql.TypeChar{255},
		},
		{
			Type: "char(1,2)",
			ExpectedErr: "parse error",
		},
		{
			Type: "datetime",
			Expected: gosql.TypeDateTime{},
		},
		{
			Type: "datetime(3)",
			Expected: gosql.TypeDateTime{3},
		},
		{
			Type: "datetime(10)",
			Expected: gosql.TypeDateTime{10},
		},
		{
			Type: "datetime(1,2)",
			ExpectedErr: "parse error",
		},
		{
			Type: "decimal",
			ExpectedErr: "parse error",
		},
		{
			Type: "decimal(123)",
			ExpectedErr: "parse error",
		},
		{
			Type: "decimal(18,9)",
			Expected: gosql.TypeDecimal{P: 18, S: 9},
		},
		{
			Type: "decimal(10, 20)",
			Expected: gosql.TypeDecimal{P: 10, S: 20},
		},
		{
			Type: "decimal(1,2,3)",
			ExpectedErr: "parse error",
		},
		{
			Type: "double",
			Expected: gosql.TypeDouble{},
		},
		{
			Type: "float",
			Expected: gosql.TypeFloat{},
		},
		{
			Type: "int",
			Expected: gosql.TypeInt{},
		},
		{
			Type: "smallint",
			Expected: gosql.TypeSmallInt{},
		},
		{
			Type: "tinyint",
			Expected: gosql.TypeTinyInt{},
		},
		{
			Type: "varchar",
			ExpectedErr: "parse error",
		},
		{
			Type: "varchar(123)",
			Expected: gosql.TypeVarChar{123},
		},
		{
			Type: "varchar(255)",
			Expected: gosql.TypeVarChar{255},
		},
		{
			Type: "varchar(1,2)",
			ExpectedErr: "parse error",
		},
	}

	for _, test := range testCases {
		t.Run(test.Type, func(t *testing.T) {

			actual, err := gosql.ParseType(test.Type)

			if test.ExpectedErr != "" {
				require.ErrorContains(t, err, test.ExpectedErr)
				return
			}

			require.NoError(t, err)

			require.Equal(t, test.Expected, actual)
		})
	}
}
