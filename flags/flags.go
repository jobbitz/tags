// Copyright 2020 Job Stoit. All rights reserved.

package flags

import (
	"os"
	"reflect"
	"regexp"
	"strings"

	"github.com/jobstoit/strct"
)

// Unmarshal gets the tag and adds it to the property if set.
// Please note that flags overwrite previous values
func Unmarshal(obj interface{}) error {
	return strct.Scan(obj, func(field reflect.StructField, value *reflect.Value) error {
		tagVal := field.Tag.Get(`flag`)
		if tagVal == `` {
			return nil
		}

		match := matchArg(tagVal, value.Kind() == reflect.Bool)
		if match == `` {
			return nil
		}

		return strct.ParseHard(match, value)
	})
}

func matchArg(fl string, isBool bool) string {
	args := strings.Join(os.Args[1:], ` `)
	if !regexp.MustCompile(`[\-]{1,2}` + fl).MatchString(args) {
		return ``
	}

	if isBool {
		return `true`
	}

	match := regexp.MustCompile(`[\-]{1,2}` + fl + `(\=|\ )(\w+)`).
		FindStringSubmatch(args)
	if len(match) > 2 {
		return match[2]
	}
	return ``
}
