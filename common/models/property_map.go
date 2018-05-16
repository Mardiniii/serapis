package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

// PropertyMap maps JSONB properties with Postgres
type PropertyMap map[string]interface{}

// Value transforms map to valid database format
func (p PropertyMap) Value() (driver.Value, error) {
	j, err := json.Marshal(p)
	return j, err
}

// Scan transforms database format to Golang map
func (p *PropertyMap) Scan(src interface{}) error {
	source, ok := src.([]byte)
	if !ok {
		return errors.New("Type assertion .([]byte) failed")
	}

	var i interface{}
	err := json.Unmarshal(source, &i)
	if err != nil {
		return err
	}

	*p, ok = i.(map[string]interface{})
	if !ok {
		return errors.New("Type assertion .(map[string]interface{}) failed")
	}

	return nil
}
