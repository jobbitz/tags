// Copyright 2019 Job Stoit. All rights reserved.

// Package env is for parsing environment variables.
//
// Usage
//
// Create a struct that requires an environment variable and parse:
//   type dbConfig struct {
//        ConnectionString string `toml:"connection_string" env:"DB_CONNECTION_STRING"`
//        Driver           string `toml:"driver" env:"DB_DRIVER"`
//   }
//
//   func main() {
//        conf := new(dbConfig)
//
//        if err := env.Parse(conf); err != nil {
//            ...
//
package env

import (
	"os"

	"github.com/jobstoit/tags"
)

// Parse reads the object and sets the properties to
func Parse(obj interface{}) error {
	return tags.Parse(obj, `env`, func(in string) (string, error) {
		return os.Getenv(in), nil
	})
}
