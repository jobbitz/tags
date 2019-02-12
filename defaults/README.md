# Defaults
This parser sets the given default values in the structs properties

## Usage
```go
func ReadConfig(rd io.Reader) Config {
	var conf Config

	if err := toml.Unmarshal(rd, conf); err != nil {
		panic(err)
	}

	if err := defaults.Parse(&conf); err != nil {
		panic(err)
	}

	return conf
}

type Config struct {
	Name 	string 	`default:"App"`
	Develop bool 	`default:"true"`
	DB 	DbConfig
}

type DbConfig struct {
	Driver 		 string `default:"postgres"`
	ConnectionString string
}
```

