package tinydate

import (
	"database/sql/driver"
	"fmt"
)

func (td TinyDate) Value() (driver.Value, error) {
	return td.String(), nil
}

func (td *TinyDate) Scan(value interface{}) error {
	switch v := value.(type) {
	case string:
		if dt, err := Parse(iso8601Date, v); err != nil {
			return err
		} else {
			td.day = dt.day
			td.month = dt.month
			td.year = dt.year
			return nil
		}
	default:
		return fmt.Errorf("tinydate: scanning unsupported type: %T", value)
	}
}
