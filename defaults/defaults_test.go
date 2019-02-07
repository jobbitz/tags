package defaults

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type person struct {
	Name    string `default:"John Locke"`
	Surname string `default:"Picard"`
	Age     int    `default:"50"`
	Rank    rank
}

type rank struct {
	Active bool   `default:"true"`
	Name   string `default:"Ensign"`
}

func TestDefault(t *testing.T) {
	as := assert.New(t)
	ps := person{}

	as.NoError(Parse(&ps))
	as.Equal(`John Locke`, ps.Name)
	as.Equal(`Picard`, ps.Surname)
	as.Equal(50, ps.Age)
	as.Equal(true, ps.Rank.Active)
	as.Equal(`Ensign`, ps.Rank.Name)
}
