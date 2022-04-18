package utils

import (
	"bytes"
	"encoding/json"
)

func CreateResponse[T any](res T) ([]byte, error) {
	var b = new(bytes.Buffer)
	err := json.NewEncoder(b).Encode(&res)
	if err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}
