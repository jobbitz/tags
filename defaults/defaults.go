package defaults

import "github.com/jobbitz/tags"

// Parse parses the given defaults from the default tag of a struct property
func Parse(obj interface{}) error {
	return tags.Parse(obj, `default`, func(in string) (string, error) {
		return in, nil
	})
}
