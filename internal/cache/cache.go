package cache

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type Cache struct {
	C *os.File
}

func Default() (*Cache, error) {
	f, err := os.OpenFile("tmp.txt", os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		return nil, err
	}

	return &Cache{
		f,
	}, nil

}

func (c *Cache) Read() (*map[string][]byte, error) {
	// read data from file,
	rawFileData, err := io.ReadAll(c.C)
	if err != nil {
		fmt.Printf("%v", err)
		return nil, err
	}

	var fileData map[string][]byte = map[string][]byte{}

	json.Unmarshal(rawFileData, &fileData)

	return &fileData, nil

}

func (c *Cache) Update(dataPtr *map[string][]byte, url string, res *http.Response) ([]byte, error) {
	// retrieve body
	data, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	// retrieve headers
	headers := res.Header

	headersKey := fmt.Sprintf("%v-[headers]", url)
	bodyKey := fmt.Sprintf("%v-[body]", url)

	c.updateBody(dataPtr, bodyKey, data)
	c.updateHeaders(dataPtr, headersKey, headers)

	// write data to file
	rawFileData, err := json.Marshal(*dataPtr)
	if err != nil {
		return nil, err
	}

	if err = os.WriteFile("tmp.txt", rawFileData, 0777); err != nil {
		return nil, err
	}

	c.C.Sync()

	return data, nil
}
