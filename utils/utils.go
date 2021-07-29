package utils

import (
	"ImyouboHomeKit/errors"
	"encoding/json"
)

func JsonUnmarshal(bytes []byte, target interface{}) (err error) {
	if err = json.Unmarshal(bytes, target); err != nil {
		return errors.ErrJsonUnmarshal(err, bytes)
	}
	return nil
}

func ContainStr(source []string, target string) bool {
	for _, elem := range source {
		if elem == target {
			return true
		}
	}
	return false
}