// Copyright 2021 Job Stoit. All rights reserved.

// Package forms is for parsing http.Request form data
//
// Usage:
//   func MyHandler(w http.ResponseWriter, r *http.Request) {
//       var body reqBody
//       if err := form.Decode(r, &body); err != nil {
//            w.WriterHeader(http.StatusUnprocessibleEnitity)
//            return
//       }
//   }
//
//   type reqBody struct {
//       Name string `form:"name"`
//       Age  int    `form:"age"`
//   }
package form

import (
	"fmt"
	"net/http"
)

const (
	contentTypeVal = `multipart/form-data`
	contentTypeKey = `Content-Type`

	tag = `form`
)

// ErrNoHeaders is returned if the request doesnt contain the right content type for parsing the form
var ErrNoHeaders = fmt.Errorf("error request doesnt contain the Content-Type: '%s'", contentTypeVal)

// Decode parses the form values into the given object or returns an error
// if the Content-Type header is not set or if the parsing fails
func Decode(r *http.Request, obj interface{}) error {
	return NewDecoder(r).Decode(obj)
}

// Encoder writes the given data to the the form and adds the Content-Type
// header to the request
func Encode(r *http.Request, obj interface{}) error {
	return NewEncoder(r).Encode(obj)
}
