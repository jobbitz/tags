// Copyright 2020 Job Stoit. All rights reserved.

// Package flags is for parsing execution flags.
//
// Usage
//
// Create a struct that requires a flag and parse:
//   type Config struct {
//       ConnectionString string   `flag:"cs"`
//       Driver           string   `flag:"driver"`
//       Args             []string `flag:"*"`
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

var mainSet Set

// Parse returns the
func Parse(obj interface{}) error {
	if !mainSet.parsed {
		mainSet.args = os.Args[1:]
	}
	return mainSet.Parse(obj)
}

// Parsed returns true if any flags have been parsed
func Parsed() bool {
	return mainSet.Parsed()
}

// Args returns the remaining arguments
func Args() []string {
	return mainSet.Args()
}

// Set contains a new set of flags with the given arguments
type Set struct {
	args    []string
	argPtrs []*reflect.Value
	parsed  bool
}

// New creates a new FlagSet
func New(args []string) *Set {
	x := new(Set)
	x.args = args
	return x
}

// Parse gets the tag and adds it to the property if set.
// Please note that flags overwrite previous values
func (x *Set) Parse(obj interface{}) error {
	if !x.parsed {
		x.parsed = true
	}

	defer x.parseArgsToPtr()

	return strct.Scan(obj, func(field reflect.StructField, value *reflect.Value) error {
		tagVal := field.Tag.Get(`flag`)
		if tagVal == `` {
			return nil
		}

		if tagVal == `*` { // parse args
			x.argPtrs = append(x.argPtrs, value)
			return nil
		}

		match := x.matchArg(tagVal, value.Kind() == reflect.Bool)
		if match == `` {
			return nil
		}

		return strct.ParseHard(match, value)
	})
}

// Parsed returns if any flags have been parsed
func (x Set) Parsed() bool {
	return x.parsed
}

// Args returns the remainder of the arguments
func (x Set) Args() []string {
	return x.args
}

func (x *Set) parseArgsToPtr() {
	for _, ptr := range x.argPtrs {
		if ptr != nil &&
			ptr.Kind() == reflect.Slice {
			ptr.Set(reflect.ValueOf(x.args))
		}
	}
}

func (x *Set) matchArg(fl string, isBool bool) string {
	reg := regexp.MustCompile(`[\-]{1,2}` + fl)
	var key, value string

	for i, arg := range x.args {
		if reg.MatchString(arg) {
			key = arg
			if len(x.args) > i {
				value = x.args[i+1]
				break
			}
		}
	}

	if isBool && !regexp.MustCompile(`^(true|false|0|1|TRUE|FALSE)$`).MatchString(value) {
		x.args = remove(x.args, key)
		return `true`
	}

	if value != `` {
		x.args = remove(x.args, key, value)
	}

	return value
}

func remove(ssl []string, items ...string) []string {
	if len(items) < 1 {
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
