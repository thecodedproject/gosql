package gosql

import (
	"errors"
	"io"
)

func WriteStatements(
	w io.Writer,
	statements []Statement,
) error {


	for _, s := range statements {

		t, ok := s.(CreateTable)
		if !ok {
			return errors.New("unknown statement type")
		}

		w.Write([]byte("create table " + t.Name + " (\n"))

		for _, f := range t.Fields {

			fullType, err := f.Type.FullType()
			if err != nil {
				return err
			}

			w.Write([]byte("  " + f.Name + " " + fullType))

			if f.PrimaryKey {
				w.Write([]byte(" primary key"))
			}

			if f.AutoIncrement {
				w.Write([]byte(" auto_increment"))
			}

			if f.NotNull {
				w.Write([]byte(" not null"))
			}

			w.Write([]byte(",\n"))
		}

		w.Write([]byte(");\n"))
	}

	return nil
}
