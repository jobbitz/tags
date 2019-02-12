package env

import (
	"os"

	"github.com/jobbitz/tags"
)

// Unmarshal reads the object and sets the properties to
func Unmarshal(obj interface{}) (err error) {
	return tags.Parse(obj, `env`, func(in string) (string, error) {
		return os.Getenv(in), nil
	})
}
