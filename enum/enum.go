package enum

import (
	"fmt"
	"strings"

	"github.com/jobbitz/tags"
)

// Err gets trown if the enum is incorrect
type Err struct {
	Valid  []string
	Actual interface{}
}

// Error is the error implementation of enum error
func (err Err) Error() string {
	return fmt.Sprintf(`Invalid enum: %s, posible: %s`, err.Actual, strings.Join(err.Valid, `, `))
}

// Validate checks the values of the object if the enums are correct
func Validate(obj interface{}) error {
	return tags.Scan(obj, `enum`, func(tagval string, val interface{}) error {
		err := Err{strings.Split(tagval, `;`), fmt.Sprint(val)}
		for _, vld := range err.Valid {
			if vld == err.Actual {
				return nil
			}
		}
		return err
	})
}
