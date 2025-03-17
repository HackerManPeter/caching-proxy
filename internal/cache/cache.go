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

func Connect() (*Cache, error) {
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

	c.updateBody(dataPtr, GetBodyKey(url), data)

	if err := c.updateHeaders(dataPtr, GetHeaderKey(url), headers); err != nil {
		return nil, err
	}

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

func GetBodyKey(url string) string {
	return fmt.Sprintf("BODY_FOR_{%v}", url)
}

func GetHeaderKey(url string) string {
	return fmt.Sprintf("HEADERS_FOR_{%v}", url)
}
