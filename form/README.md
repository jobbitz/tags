# Forms
This encoder and decoder will parse values from and to your http.Request form data.

## Usage
```go
func MyHandler(w http.ResponseWriter, r *http.Request) {
	var body requestBody
	if err := form.Decode(r, &body); err != nil {
		w.WriteHeader(http.StatusUnprocessibleEntity)
		return
	}

	...
}

type requestBody struct {
	Name       string `form:"name"`
	Age        int    `form:"age"`
	NewsLetter bool   `form:"news_letter"`
}
```


