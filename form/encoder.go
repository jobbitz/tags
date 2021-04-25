// Copyright 2021 Job Stoit. All rights reserved.

package form

import (
	"fmt"
	"net/http"

	"github.com/jobstoit/tags"
)

// Encoder will write form values into a http.Request
type Encoder struct {
	r *http.Request
}

// NewEncoder returns a new Encoder and sets the content type header to
// the given request
func NewEncoder(r *http.Request) *Encoder {
	x := new(Encoder)

	x.r = r
	x.r.Header.Add(contentTypeKey, contentTypeVal)

	return x
}

// Encode writes the values of the given object into the request form
func (x Encoder) Encode(obj interface{}) error {
	return tags.Scan(obj, tag, func(key string, val interface{}) error {
		x.r.Form.Add(key, fmt.Sprint(val))
		return nil
	})
}
