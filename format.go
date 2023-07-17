package tinydate

import (
	"fmt"
	"time"
)

// Parse a layout into a TinyDate, truncating any non-date information
func Parse(layout, value string) (TinyDate, error) {
	t, err := time.Parse(layout, value)
	if err != nil {
		return TinyDate{}, fmt.Errorf("tinydate Parse: %v", err)
	}
	td, err := FromTime(t)
	if err != nil {
		return TinyDate{}, fmt.Errorf("tinydate Parse: %v", err)
	}
	return td, nil
}

// FromString is shortcut for MustParse("2006-01-02", value)
func FromString(value string) TinyDate {
	return MustParse(iso8601Date, value)
}

// MustParse panic-if-error version of Parse
func MustParse(layout, value string) TinyDate {
	t, err := time.Parse(layout, value)
	if err != nil {
		panic(fmt.Errorf("tinydate MustParse: %v", err))
	}
	td, err := FromTime(t)
	if err != nil {
		panic(fmt.Errorf("tinydate MustParse: %v", err))
	}
	return td
}

// ParseInLocation parses a layout into a TinyDate including a location.
// The input is converted into UTC
func ParseInLocation(layout, value string, loc *time.Location) (TinyDate, error) {
	t, err := time.ParseInLocation(layout, value, loc)
	if err != nil {
		return TinyDate{}, fmt.Errorf("tinydate ParseInLocation: %v", err)
	}
	td, err := FromTime(t)
	if err != nil {
		return TinyDate{}, fmt.Errorf("tinydate ParseInLocation: %v", err)
	}
	return td, nil
}

// AppendFormat is like Format but appends the textual representation to b and returns the extended buffer
func (td TinyDate) AppendFormat(b []byte, layout string) []byte {
	t := td.ToTime()
	return t.AppendFormat(b, layout)
}

// Format returns a formatted date, as specified by the standard time library
// https://golang.org/src/time/format.go?s=16029:16071#L485
func (td TinyDate) Format(layout string) string {
	t := td.ToTime()
	return t.Format(layout)
}
