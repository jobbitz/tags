/*
Package env is for parsing enviroment variables.

Usage

Create a struct that requires an enviroment variable and parse:
	type dbConfig struct {
		ConnectionString string `env:"DB_URL"`
		Driver 		 string `env:"DB_DRIVER"`
	}

Than create use the parser:
	main() {
		config := dbConfig{}

		env.Unmarshal(&config)
	...
*/
package env
