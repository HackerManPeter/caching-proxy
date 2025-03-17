package cache

import (
	"encoding/json"
	"net/http"
)

func (c *Cache) updateHeaders(dataPtr *map[string][]byte, key string, headers http.Header) error {
	// get actual data
	data := *dataPtr

	// convert headers to byte slice
	value, err := json.Marshal(headers)
	if err != nil {
		return err
	}

	// assign headers to map key
	data[key] = value

	return nil

}

func (c *Cache) updateBody(dataPtr *map[string][]byte, key string, value []byte) {
	data := *dataPtr
	data[key] = value

}
