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

func matchArg(fl string, isBool bool) (r string) {
	reg := regexp.MustCompile(`[\-]{1,2}` + fl)
	var in int

	for i, arg := range args {
		if reg.MatchString(arg) {
			if i+1 < len(args) {
				r = args[i+1]
				in = i
				break
			}
		}
	}

	if isBool && !regexp.MustCompile(`^(true|false|0|1|TRUE|FALSE)$`).MatchString(r) {
		args = remove(args, in)
		return `true`
	}

	if r != `` {
		args = remove(args, in)
		args = remove(args, in)
	}

	return
}

func remove(ssl []string, index int) []string {
	if len(ssl) < 1 {
		return ssl
	}
	return append(ssl[:index], ssl[index+1:]...)
}
