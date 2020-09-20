package models

import (
	"database/sql/driver"
	"fmt"
	"strings"
)

type Tags []string

func (t Tags) Value() (driver.Value, error) {
	return strings.Join(t, ","), nil
}

func (t *Tags) Scan(value interface{}) error {
	s, ok := value.(string)
	if !ok {
		return fmt.Errorf("invalid type for tags")
	}
	tags := strings.Split(s, ",")
	for _, tag := range tags {
		*t = append(*t, tag)
	}
	return nil
}
