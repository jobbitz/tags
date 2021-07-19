package tags

import (
	"errors"
	"testing"

	"git.fuyu.moe/Fuyu/assert"
	"github.com/jobstoit/strct"
)

var (
	// values
	sb = `true;false;true`
	in = `5`
	st = `check this out`

	errNotFound = errors.New(`a not found error`)
	errReq      = errors.New(`Missing arguments`)
)

// some test objects
type testObj struct {
	Name  string `tag:"name"`
	Slice []bool `tag:"truth"`
	Sub   subObj
}

type subObj struct {
	Ammount int `tag:"value"`
}

type notFound struct {
	SomeVal string `tag:"unknown"`
}

type notParseble struct {
	SomeVal bool `tag:"name"`
}

type scanOk struct {
	Values string `validate:"required"`
}

func testParser(val string) (string, error) {
	switch val {
	case `name`:
		return st, nil

	case `truth`:
		return sb, nil

	case `value`:
		return in, nil
	}

	return ``, errNotFound
}

func TestParse(t *testing.T) {
	as := assert.New(t)

	// Test custom error
	nfok := notFound{}
	err := Parse(&nfok, `tag`, testParser)
	as.Error(err)
	as.Eq(errNotFound, err)

	// test error on unparseble string
	npok := notParseble{}
	err = Parse(&npok, `tag`, testParser)
	as.Error(err)

	// test no pointer error message
	ok := testObj{}
	err = Parse(ok, `tag`, testParser)
	as.Error(err)
	as.Eq(strct.ErrNoPtr, err)

	// test right senaria
	ok.Sub.Ammount = 19 // shouldn't get overwritten
	err = Parse(&ok, `tag`, testParser)
	as.NoError(err)
	as.Eq(st, ok.Name)
	as.Eq(19, ok.Sub.Ammount)
	as.Eq(true, ok.Slice[0])
	as.Eq(false, ok.Slice[1])
	as.Eq(true, ok.Slice[2])
}

func TestParseHard(t *testing.T) {
	as := assert.New(t)

	obj := testObj{
		Name:  `different_name`,
		Slice: []bool{false, true, false},
		Sub:   subObj{99},
	}

	as.NoError(ParseHard(&obj, `tag`, testParser))
	as.Eq(st, obj.Name)
	as.Eq(5, obj.Sub.Ammount)
	as.Eq(true, obj.Slice[0])
	as.Eq(false, obj.Slice[1])
	as.Eq(true, obj.Slice[2])
}

func TestScan(t *testing.T) {
	scanner := func(tagval string, val interface{}) error {
		s, sok := val.(string)
		switch true {
		case sok:
			if tagval == `required` && s == `` {
				return errReq
			}

		default:
			return nil
		}

		return nil
	}

	as := assert.New(t)

	obj := scanOk{
		Values: `this is an active one`,
	}
	err := Scan(&obj, `validate`, scanner)
	as.NoError(err)

	obj.Values = ``
	err = Scan(&obj, `validate`, scanner)
	as.Error(err)
	as.Eq(errReq, err)
}
