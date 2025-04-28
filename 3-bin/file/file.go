package file

import (
	"fmt"
	"os"
	"strings"
)

type JsonDb struct {
	filename string
}

func NewJsonDb(name string) *JsonDb {
	return &JsonDb{
		filename: name,
	}
}

func (db *JsonDb) Write(content []byte) {
	err := os.WriteFile(db.filename, content, 0744)
	if err != nil {
		fmt.Println(err)
	}
}

func (db *JsonDb) Read() ([]byte, error) {
	data, err := os.ReadFile(db.filename)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return data, nil
}

func ReadFile(path string) ([]byte, error) {
	if !strings.HasSuffix(path, ".json") {
		fmt.Println("Это не JSON файл")
		return make([]byte, 0), fmt.Errorf("IS_NOT_JSON_FROM_PATH: %s", path)
	}
	bytes, err := os.ReadFile(path)
	if err != nil {
		return make([]byte, 0), fmt.Errorf("ERROR_READ_FILE: %s", path)
	}
	return bytes, nil
}
