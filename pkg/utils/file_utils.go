/*
Package utils provides utility functions for file operations.
*/

package utils

import (
	"encoding/json"
	"io/ioutil"
)

// ReadJSONFile reads a JSON file and unmarshals its content into the provided interface.
func ReadJSONFile(filepath string, v interface{}) error {
	data, err := ioutil.ReadFile(filepath)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, v)
}
