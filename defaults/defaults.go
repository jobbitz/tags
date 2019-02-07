package defaults

import "github.com/jobbitz/tags"

// Parse parses the given defaults from the default tag of a struct property
func Parse(obj interface{}) error {
	parser := func(in string) (string, error) {
		return in, nil
	}
	return tags.Parse(obj, `default`, parser)
}
