package filex

import (
	"encoding/json"
	"os"
)

func JsonGet[T any](file string, entrys *[]T) error {
	openFile, err := os.OpenFile(file, os.O_RDONLY, 0666)
	if err != nil {
		return err
	}
	defer openFile.Close()

	decoder := json.NewDecoder(openFile)
	for {
		var entry T
		if err := decoder.Decode(&entry); err != nil {
			break
		}
		*entrys = append(*entrys, entry)
	}

	return nil
}

func JsonInsert[T any](file string, entry T) error {
	// T to json
	jsonData, err := json.Marshal(entry)
	if err != nil {
		return err
	}
	// open file
	openFile, err := os.OpenFile(file, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return err
	}
	defer openFile.Close()
	// write json
	if _, err := openFile.Write(jsonData); err != nil {
		return err
	}
	return nil
}
