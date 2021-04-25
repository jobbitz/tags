// Copyright 2021 Job Stoit. All rights reserved.

// Package forms is for parsing http.Request form data
package form

import (
	"fmt"
	"net/http"
)

const contentType = `multipart/form-data`

// ErrNoHeaders is returned if the request doesnt contain the right content type for parsing the form
var ErrNoHeaders = fmt.Errorf("error request doesnt contain the Content-Type: '%s'", contentType)

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
