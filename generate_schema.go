package gosql

import (
)

func GenerateSchema(
	filepath string,
	s []Statement,
) error {

	writer, err := CreatePathAndOpen(filepath)
	if err != nil {
		return err
	}
	err = WriteStatements(writer, s)
	if err != nil {
		return err
	}

	return nil
}
