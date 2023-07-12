package gosql_test

import (
	"bytes"
	"testing"

	"github.com/sebdah/goldie/v2"
	"github.com/stretchr/testify/require"

	"github.com/thecodedproject/gosql"
)

func TestWriteStatements(t *testing.T) {

	testCases := []struct{
		Name string
		Statements []gosql.Statement
		ExpectedErr error
	}{
		{
			Name: "no statements writes nothing",
		},
		{
			Name: "create single table with no fields",
			Statements: []gosql.Statement{
				gosql.CreateTable{
					Name: "some_model",
				},
			},
		},
		{
			Name: "create single table with some fields",
			Statements: []gosql.Statement{
				gosql.CreateTable{
					Name: "my_table",
					Fields: []gosql.Field{
						{
							Name: "id",
							Type: gosql.TypeInt{},
							PrimaryKey: true,
							AutoIncrement: true,
						},
						{
							Name: "some_datetime",
							Type: gosql.TypeDateTime{},
							NotNull: true,
						},
						{
							Name: "some_float",
							Type: gosql.TypeFloat{},
							NotNull: true,
						},
					},
				},
			},
		},
	}

	for _, test := range testCases {
		t.Run(test.Name, func(t *testing.T) {

			buffer := bytes.NewBuffer(nil)

			err := gosql.WriteStatements(
				buffer,
				test.Statements,
			)

			if test.ExpectedErr != nil {
				require.Equal(t, test.ExpectedErr, err)
				return
			}

			require.NoError(t, err)

			g := goldie.New(t)
			g.Assert(t, t.Name(), buffer.Bytes())
		})
	}
}


