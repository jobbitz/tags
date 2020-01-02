// Copyright 2020 Job Stoit. All rights reserved.

// Package flags is for parsing execution flags.
//
// Usage
//
// Create a struct that requires a flag and parse:
//   type Config struct {
//       ConnectionString string   `flag:"cs"`
//       Driver           string   `flag:"driver"`
//       Args             []string `flag:"*"
//   }
//
//   func main() {
//       conf := new(Config)
//
//       if err := flags.Parse(conf); err != nil {
//           ...
// The '*' can be used to parse in all the args
package flags

import (
	"os"
	"reflect"
	"regexp"

	"github.com/jobstoit/strct"
)

var (
	args    []string
	argPtrs []*reflect.Value
	parsed  bool
)

// Parse gets the tag and adds it to the property if set.
// Please note that flags overwrite previous values
func Parse(obj interface{}) error {
	if !parsed {
		args = os.Args[1:]
		parsed = true
	}

	defer parseArgsToPtr()

	return strct.Scan(obj, func(field reflect.StructField, value *reflect.Value) error {
		tagVal := field.Tag.Get(`flag`)
		if tagVal == `` {
			return nil
		}

		if tagVal == `*` { // parse args
			argPtrs = append(argPtrs, value)
			return nil
		}

		match := matchArg(tagVal, value.Kind() == reflect.Bool)
		if match == `` {
			return nil
		}

		return strct.ParseHard(match, value)
	})
}

// Parsed returns if any flags have been parsed
func Parsed() bool {
	return parsed
}

// Args returns the remainder of the arguments
func Args() []string {
	return args
}

func parseArgsToPtr() {
	for _, ptr := range argPtrs {
		if ptr != nil &&
			ptr.Kind() == reflect.Slice {
			ptr.Set(reflect.ValueOf(args))
		}
	}
}

func matchArg(fl string, isBool bool) string {
	reg := regexp.MustCompile(`[\-]{1,2}` + fl)
	var key, value string

	for i, arg := range args {
		if reg.MatchString(arg) {
			key = arg
			if i+1 < len(args) {
				value = args[i+1]
				break
			}
		}
	}

	if isBool && !regexp.MustCompile(`^(true|false|0|1|TRUE|FALSE)$`).MatchString(value) {
		args = remove(args, key)
		return `true`
	}

	if value != `` {
		args = remove(args, key, value)
	}

	return value
}

func remove(ssl []string, items ...string) []string {
	if len(ssl) < 1 || len(items) < 1 {
		return ssl
	}
	for i, s := range ssl {
		if s == items[0] {
			ssl = append(ssl[:i], ssl[i+1:]...)
			break
		}
	}
	return remove(ssl, items[1:]...)
}
