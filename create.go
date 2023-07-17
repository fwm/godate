package tinydate

import (
	"time"
)

// New creates a new TinyDate, similar to the time.Time package,
// months are specifed from 1-12, and days are specified 1-31, depending on the month
func New(year, month, day int) (TinyDate, error) {
	// use the time package's validation
	t := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
	return FromTime(t)
}

// MustNew creates a new TinyDate just like New, but panics when error occurs
// (and of course returns only TinyDate so that it can be inlined)
// it will panic when too big year is specified, however we just don't care
func MustNew(year, month, day int) TinyDate {
	if td, err := New(year, month, day); err != nil {
		panic(err)
	} else {
		return td
	}
}

// Now returns the current date
func Now() TinyDate {
	td, _ := FromTime(time.Now())
	return td
}

// Unix creates a tinydate from seconds and nanoseconds. As usual, this is truncated to
// the nearest day
func Unix(sec int64, nsec int64) TinyDate {
	td, _ := FromTime(time.Unix(sec, nsec))
	return td
}
