package enum

import (
	"fmt"
	"strings"

	"github.com/jobbitz/tags"
)

// Err gets trown if the enum is incorrect
type Err struct {
	valid  []string
	actual interface{}
}

// Error is the error implementation of enum error
func (err Err) Error() string {
	return fmt.Sprintf(`Invalid enum: %s, posible: %s`, err.actual, strings.Join(err.valid, `, `))
}

// Validate checks the values of the object if the enums are correct
func Validate(obj interface{}) error {
	return tags.Scan(obj, `enum`, func(tagval string, val interface{}) error {
		err := Err{strings.Split(tagval, `;`), fmt.Sprint(val)}
		for _, vld := range err.valid {
			if vld == err.actual {
				return nil
			}
		}
		return err
	})
}
