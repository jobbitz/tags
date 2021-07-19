package defaults

import (
	"testing"

	"git.fuyu.moe/Fuyu/assert"
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
	as.Eq(`John Locke`, ps.Name)
	as.Eq(`Picard`, ps.Surname)
	as.Eq(50, ps.Age)
	as.Eq(true, ps.Rank.Active)
	as.Eq(`Ensign`, ps.Rank.Name)
}
