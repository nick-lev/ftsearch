// helpers for full text search - indexing tool
package index

import (
	"encoding/json"
	"os"
)

type Data map[string]map[string]int

func Load(path string, data Data) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	err = json.NewDecoder(file).Decode(&data)
	if err != nil {
		return err
	}
	return nil
}

func Save(path string, data Data) error {
	file, err := os.Create(path + ".tmp")
	if err != nil {
		return err
	}
	err = json.NewEncoder(file).Encode(data)
	if err != nil {
		return err
	}
	err = os.Rename(path+".tmp", path)
	if err != nil {
		return err
	}
	return nil
}
