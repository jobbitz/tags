# Enum

Validates your enum values

## Usage
In this example we'll check a http request's body:
```go
func RequestBody(req *http.Request, dst interface{}) error {
	if err := json.NewDecoder(req.Body).Decode(dst); err != nil {
		return err
	}

	if err := validator.Validate(dst); err != nil {
		return err
	}
	
	return enum.Validate(dst)
}
```
