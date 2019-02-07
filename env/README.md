# Env
Get the machine's enviroment variables if they exist

## Usage
In this example we overwrite the toml settings with the enviroment settings
```go
func main() {
	c := config{}

	if err := toml.Unmarshal(file, &c); err != nil {
		panic(err)
	}
	
	if err := env.Unmarshal(&c); err != nil {
		panic(err)
	}

	setupWebContext(c)
}

type config struct {
	Cors []string `toml:"cors" env:"WEB_CORS"`
	DB   dbConfig `toml:"db"`
}

type dbConfig struct {
	Driver 	   string `toml:"driver" env:"DB_DRIVER"`
	Connection string `toml:"connection_string" env:"DB_STRING"`
}
```
