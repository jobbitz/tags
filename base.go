// Copyright 2019 Job Stoit

// Package tags a simplified golang reflect for your tags.
//
// Usage
//
// Create your custom parser:
// 	func ParseEnv(obj interface{}) error {
// 		return tags.Parse(obj, `env`, func(tagval string) (string, error) {
// 			envVar := os.Env(tagval)
// 			if envVar == `` {
// 				return ``, fmt.Errorf(`enviroment variable "%s" not found`, tagval)
// 			}
// 			return envVar, nil
// 		})
// 	}
//
// Creating your own scanner:
// 	func ValidateEnum(obj interface{}) error {
// 		return tags.Scan(obj, `enum`, func(tagval string, propVal interface{}) error {
// 			validEnums := strings.Split(tagval, `;`)
// 			actual := fmt.Sprint(propVal)
// 			for _, enum := range validEnums {
// 				if actual == enum {
// 					return nil
// 				}
// 			}
// 			return fmt.Errorf(`Not a valid enum`)
// 		})
// 	}
//
package tags

import (
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

		val, err := parser(tagVal)
		if err != nil {
			return err
		}

		if override {
			return strct.ParseHard(val, value)
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
