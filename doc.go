/*
Package tags a simplified golang reflect for your tags.

Usage

Create your custom parser:
	func ParseEnv(obj interface{}) error {
		return tags.Parse(obj, `env`, func(tagval string) (string, error) {
			envVar := os.Env(tagval)
			if envVar == `` {
				return ``, fmt.Errorf(`enviroment variable "%s" not found`, tagval)
			}
			return envVar, nil
		})
	}

Creating your own scanner:
	func ValidateEnum(obj interface{}) error {
		return tags.Scan(obj, `enum`, func(tagval string, propVal interface{}) error {
			validEnums := strings.Split(tagval, `;`)
			actual := fmt.Sprint(propVal)
			for _, enum := range validEnums {
				if actual == enum {
					return nil
				}
			}
			return fmt.Errorf(`Not a valid enum`)
		})
	}

*/
package tags
