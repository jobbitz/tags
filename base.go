package tags

import (
	"fmt"
	"reflect"

	"github.com/jobbitz/strct"
)

// Parse parses the given tag values using the given parser
func Parse(obj interface{}, tagname string, parser func(string) (string, error)) error {
	return parse(obj, tagname, parser, false)
}

// ParseHard parses the giver tag values using the given parser and overrides the previous values
func ParseHard(obj interface{}, tagname string, parser func(string) (string, error)) error {
	return parse(obj, tagname, parser, true)
}

func parse(obj interface{}, tagname string, parser func(string) (string, error), override bool) error {
	return strct.Scan(obj, func(field reflect.StructField, value *reflect.Value) error {
		tagVal := field.Tag.Get(tagname)
		if tagVal == `` {
			return nil
		}

		sv := fmt.Sprint(value.Interface())
		if !override && !(sv == `false` || sv == `0` || sv == `[]` || sv == ``) {
			return nil
		}

		val, err := parser(tagVal)
		if err != nil {
			return err
		}

		return strct.Parse(val, value)
	})
}

// Scan gives the given tag value and the value in that stuct property to the given scanner
func Scan(obj interface{}, tagname string, scanner func(string, interface{}) error) error {
	return strct.Scan(obj, func(field reflect.StructField, value *reflect.Value) error {
		tagval := field.Tag.Get(tagname)
		if tagval == `` {
			return nil
		}
		return scanner(tagval, value.Interface())
	})
}
