package gosql

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

func ParseType(t string) (Type, error) {

	r, err := regexp.Compile(`^[a-z]+($)|(\((\d+)(,\s*\d+)*\)$)`)
	if err != nil {
		return nil, err
	}

	if !r.MatchString(t) {
		return nil, errors.New(
			"gosql parse error '" + t + "': does not match type regexp",
		)
	}

	typeName := strings.Split(t, "(")[0]

	switch typeName {
	case "bigint": return TypeBigInt{}, nil
	case "bit": return TypeBit{}, nil
	case "char": return parseChar(t)
	case "datetime": return parseDatetime(t)
	case "decimal": return parseDecimal(t)
	case "double": return TypeDouble{}, nil
	case "float": return TypeFloat{}, nil
	case "int": return TypeInt{}, nil
	case "smallint": return TypeSmallInt{}, nil
	case "tinyint": return TypeTinyInt{}, nil
	case "varchar": return parseVarchar(t)
	}


	return nil, errors.New(
		"gosql parse error '" + t + "': unknown type",
	)
}

func parseChar(t string) (Type, error) {

	n, err := values(t)
	if err != nil {
		return nil, err
	}

	if len(n) == 0 {
		return TypeChar{}, nil
	}

	if len(n) > 1 {
		return nil, errors.New(
			"gosql parse error '" + t + "': too many args for char",
		)
	}

	return TypeChar{N: n[0]}, nil
}

func parseDatetime(t string) (Type, error) {

	n, err := values(t)
	if err != nil {
		return nil, err
	}

	if len(n) == 0 {
		return TypeDateTime{}, nil
	}

	if len(n) > 1 {
		return nil, errors.New(
			"gosql parse error '" + t + "': too many args for datetime",
		)
	}

	return TypeDateTime{N: n[0]}, nil
}

func parseDecimal(t string) (Type, error) {

	n, err := values(t)
	if err != nil {
		return nil, err
	}

	if len(n) != 2 {
		return nil, errors.New(
			"gosql parse error '" + t + "': wrong num args for decimal",
		)
	}

	return TypeDecimal{P: n[0], S: n[1]}, nil
}

func parseVarchar(t string) (Type, error) {

	n, err := values(t)
	if err != nil {
		return nil, err
	}

	if len(n) != 1 {
		return nil, errors.New(
			"gosql parse error '" + t + "': wrong num args for varchar",
		)
	}

	return TypeVarChar{N: n[0]}, nil
}

func values(t string) ([]int64, error) {

	i := strings.IndexByte(t, '(') + 1
	j := strings.IndexByte(t, ')')

	if i == -1 || j == -1 {
		return nil, nil
	}

	vs := strings.Split(t[i:j], ",")

	ret := make([]int64, 0, len(vs))

	for _, val := range vs {
		val = strings.TrimSpace(val)
		n, err := strconv.ParseInt(val, 10, 64)
		if err != nil {
			return nil, errors.Wrap(
				err,
				"gosql parse error '" + t + "'",
			)
		}

		ret = append(ret, n)
	}

	return ret, nil
}
