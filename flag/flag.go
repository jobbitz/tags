package flag

import (
	"errors"
	f "flag"
	"reflect"

	"github.com/jobbitz/strct"
)

// Unmarshal reads the object and parses the flag value on to that object
func Unmarshal(obj interface{}) (err error) {
	if !f.Parsed() {
		f.Parse()
	}
	return strct.Scan(obj, func(field reflect.StructField, value *reflect.Value) error {
		tagVal := field.Tag.Get(`flag`)

		switch value.Kind() {
		case reflect.Bool:
			x := f.Bool(tagVal, false, ``)
			value.SetBool(*x)

		case reflect.Float32, reflect.Float64:
			x := f.Float64(tagVal, 0.0, ``)
			value.SetFloat(*x)

		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			x := f.Int64(tagVal, 0, ``)
			value.SetInt(*x)

		case reflect.String:
			x := f.String(tagVal, ``, ``)
			value.SetString(*x)

		default:
			return errors.New(`value not a parseble flag`)
		}
		return nil
	})
}
