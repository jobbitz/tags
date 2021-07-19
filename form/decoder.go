// Copyright 2021 Job Stoit. All rights reserved.

package form

import (
	"net/http"

	"github.com/jobstoit/tags"
)

// Decoder can decode any forms into a given object
type Decoder struct {
	r *http.Request
}

// NewDecoder returns a new Decoder
func NewDecoder(r *http.Request) *Decoder {
	x := new(Decoder)

	x.r = r

	return x
}

// Decode parses the form values into the given object or returns an error
// if the Content-Type header is not set or if the parsing fails
func (x Decoder) Decode(obj interface{}) error {
	return tags.Parse(obj, tag, func(in string) (string, error) {
		return x.r.FormValue(in), nil
	})
}
