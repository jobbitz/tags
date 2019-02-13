package tags

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"time"
)

var (
	// ErrPtr returns if the given object is not a pointer
	ErrPtr = fmt.Errorf(`Given object is not a pointer or a struct`) // nolint: staticcheck
)

// Parse parses the given tag values using the given parser
func Parse(obj interface{}, tagname string, parser func(string) (string, error)) error {
	return parse(obj, tagname, parser, false)
}

// ParseHard parses the giver tag values using the given parser and overrides the previous values
func ParseHard(obj interface{}, tagname string, parser func(string) (string, error)) error {
	return parse(obj, tagname, parser, true)
}

func parse(obj interface{}, tagname string, parser func(string) (string, error), override bool) error {
	return exec(obj, tagname, func(tag string, f *reflect.Value) error {
		value, err := parser(tag)
		if err != nil {
			return err
		}
		return decode(f, value, override)
	})
}

// Scan gives the given tag value and the value in that stuct property to the given scanner
func Scan(obj interface{}, tagname string, scanner func(string, interface{}) error) error {
	scanFunc := func(tag string, f *reflect.Value) error {
		return scanner(tag, f.Interface())
	}
	return exec(obj, tagname, scanFunc)
}

func exec(obj interface{}, tagname string, action func(string, *reflect.Value) error) error { // nolint: gocyclo
	rv := reflect.ValueOf(obj)
	if rv.Kind() != reflect.Ptr || rv.IsNil() {
		return ErrPtr
	}

	rv = rv.Elem()
	if rv.Kind() != reflect.Struct {
		return ErrPtr
	}

	t := rv.Type()
	for i := 0; i < rv.NumField(); i++ {
		f := rv.Field(i)
		switch f.Kind() {
		case reflect.Ptr:
			if f.Elem().Kind() != reflect.Struct {
				break
			}

			f = f.Elem()
			fallthrough

		case reflect.Struct:
			if !f.Addr().CanInterface() {
				continue
			}

			if err := exec(f.Addr().Interface(), tagname, action); err != nil {
				return err
			}
		}

		if !f.CanSet() {
			continue
		}

		tag := t.Field(i).Tag.Get(tagname)
		if tag == `` {
			continue
		}

		if err := action(tag, &f); err != nil {
			return err
		}

	}

	return nil
}

func decode(f *reflect.Value, tv string, override bool) error { // nolint: gocyclo
	if tv == `` {
		return nil
	}
	currentVal := fmt.Sprint(f.Interface())
	fmt.Println(currentVal)
	switch f.Kind() {
	case reflect.Bool:
		if currentVal != `false` && !override {
			return nil
		}

		v, err := strconv.ParseBool(tv)
		if err != nil {
			return err
		}
		f.SetBool(v)

	case reflect.Float32, reflect.Float64:
		if currentVal != `0` && !override { // nolint: goconst
			return nil
		}

		bits := f.Type().Bits()
		v, err := strconv.ParseFloat(tv, bits)
		if err != nil {
			return err
		}
		f.SetFloat(v)

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		if t := f.Type(); t.PkgPath() == `time` && t.Name() == `Duration` {
			v, err := time.ParseDuration(tv)
			if err != nil {
				return err
			}
			f.SetInt(int64(v))
		} else {
			if currentVal != `0` && !override { // nolint: goconst
				return nil
			}

			bits := f.Type().Bits()
			v, err := strconv.ParseInt(tv, 0, bits)
			if err != nil {
				return err
			}
			f.SetInt(v)
		}

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		if currentVal != `0` && !override { // nolint: goconst
			return nil
		}

		bits := f.Type().Bits()
		v, err := strconv.ParseUint(tv, 0, bits)
		if err != nil {
			return err
		}
		f.SetUint(v)

	case reflect.String:
		if currentVal != `` && !override {
			return nil
		}
		f.SetString(tv)

	case reflect.Slice:
		if currentVal != `[]` && !override {
			return nil
		}

		parts := strings.Split(tv, `;`)
		slice := reflect.MakeSlice(f.Type(), len(parts), len(parts))
		for i, part := range parts {
			part = strings.TrimSpace(part)
			e := slice.Index(i)
			if err := decode(&e, part, override); err != nil {
				return err
			}
		}
		f.Set(slice)
	}

	return nil
}
