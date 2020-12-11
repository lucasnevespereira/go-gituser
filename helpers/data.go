package helpers

import (
	"fmt"
	"io/ioutil"
)

// GetDataFromJSON is a func that returns the data from a JSON File
func GetDataFromJSON(filename string) ([]byte, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
	}
	return data, nil
}
