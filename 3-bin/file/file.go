package file

import (
	"fmt"
	"os"
	"strings"
)

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
