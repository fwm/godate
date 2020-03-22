package tinydate

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	td, err := New(2020, 04, 3)
	assert.Nil(t, err)
	assert.Equal(t, TinyDate{year: 2020, month: 3, day: 2}, td)
}

func TestNewManyMonths(t *testing.T) {
	td, err := New(2020, 13, 3)
	assert.Nil(t, err)
	assert.Equal(t, TinyDate{year: 2021, month: 0, day: 2}, td)
}

func TestNewManyDays(t *testing.T) {
	td, err := New(2020, 1, 32)
	assert.Nil(t, err)
	assert.Equal(t, TinyDate{year: 2020, month: 1, day: 0}, td)
}

func TestUnix(t *testing.T) {
	td := Unix(1, 0)
	assert.Equal(t, TinyDate{year: 1970, month: 0, day: 0}, td)

	td = Unix(86400, 0)
	assert.Equal(t, TinyDate{year: 1970, month: 0, day: 1}, td)

	td = Unix(2678400, 0)
	assert.Equal(t, TinyDate{year: 1970, month: 1, day: 0}, td)

	td = Unix(31536000, 0)
	assert.Equal(t, TinyDate{year: 1971, month: 0, day: 0}, td)
}
