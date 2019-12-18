# Tags
[![GoDoc](https://godoc.org/github.com/jobstoit/tags?status.svg)](https://godoc.org/github.com/jobstoit/tags)
[![Build Status](https://img.shields.io/endpoint.svg?url=https%3A%2F%2Factions-badge.atrox.dev%2Fjobstoit%2Ftags%2Fbadge&style=flat)](https://actions-badge.atrox.dev/jobstoit/tags/goto)

An interface package for parsing tags in structs simplifying tag based golang reflections

## Usage
Create your custom tag reader simply defining a parser:
```go
func main() {
	obj := Obj{}

	if err := CustomUnmarshaler(&obj); err != nil {
		panic(err)
	}

	if err := CustomScanner(&obj); err != nil {
		panic(err)
	}
}

type Obj struct {
	Val   bool   `customTag:"yourcase"`
	StVal string `customTag:"required"`
}

func CustomUnmarshaler(obj interface{}) error {
	return tags.Parse(obj, `customTag`, func(tagVal string) (string, error) {
		switch tagVal {
		case `yourcase`:
			return `true`, nil
		default:
			return ``, errCustom
		}
	})
}

func CustomScanner(obj interface{}) error {
	return tags.Scan(obj, `customTag`, func(tagVal string, value interface{}) error {
		v, ok := value.(string)
		if !ok {
			return nil
		}

		if tagVal == `required` && v == `` {
			return errRequired
		}

		return nil
	})
}

```
you get the values of a parent value in the string seperated by a dot:
```go
err := tags.Parse(obj, `tagname`, func(val string) (string, error) {
	allvalues := strings.Split(val, `.`)
	childVal := allvalues[len(allvalues)-1]
	...
```

## Install
Get this package using go get:
```bash
$ go get github.com/jobbitz/tags/...
```
