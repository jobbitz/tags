/*
Package defaults sets default values in your struct.

Usage

Create your struct that needs defaults
	type Config struct {
		Name 	string 	`default:"App"`
		Develop bool 	`default:"true"`
		DB 	DbConfig
	}

	type DbConfig struct {
		Driver 		 string `defaults:"postgress"`
		ConnectionString string
	}

And than read it:
	func ReadConfig(rd io.Reader) Config {
		var conf Config

		if err := toml.Unmarshal(&conf) {
			panic(err)
		}

		if err := defaults.Parse(&conf) {
			panic(err)
		}

		return config
	}
*/
package defaults
