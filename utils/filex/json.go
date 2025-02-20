package filex

import (
	"encoding/json"
)

func JsonGet[T any](file string, entries *[]T) error {
	f, err := FileOpen(file)
	if err != nil {
		return err
	}
	defer f.Close()

	// 检查文件是否为空
	stat, err := f.Stat()
	if err != nil {
		return err
	}
	if stat.Size() == 0 {
		*entries = []T{}
		return nil
	}

	decoder := json.NewDecoder(f)
	if err := decoder.Decode(entries); err != nil {
		return err
	}
	return nil
}

func JsonSet[T any](file string, entries []T) error {
	err := FileDelete(file)
	if err != nil {
		return err
	}
	f, err := FileOpen(file)
	if err != nil {
		return err
	}
	defer f.Close()

	encoder := json.NewEncoder(f)
	if err := encoder.Encode(entries); err != nil {
		return err
	}
	return nil
}

func JsonInsert[T any](file string, entry T) error {
	var entries []T
	if err := JsonGet(file, &entries); err != nil {
		return err
	}
	entries = append(entries, entry)
	return JsonSet(file, entries)
}
