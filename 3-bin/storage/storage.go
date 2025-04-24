package storage

import (
	"app/hw-3/bins"
	"encoding/json"
	"fmt"
	"os"
)

const STORAGE string = "storage/storage.json"

type Storage struct {
	Bins []bins.Bin `json:"bins"`
}

func (storage *Storage) AddBin(bin bins.Bin) {
	storage.Bins = append(storage.Bins, bin)
}

func (storage *Storage) GetBin(id string) (bins.Bin, error) {
	for _, bin := range storage.Bins {
		if bin.Id == id {
			return bin, nil
		}
	}
	return bins.Bin{}, fmt.Errorf("NOT_FOUND_BIN_WITH_ID_%s", id)
}

func (storage *Storage) WriteToJSON() {
	data, err := json.Marshal(storage)
	if err != nil {
		fmt.Println("Не удалось преобразовать в JSON")
	}
	err = os.WriteFile(STORAGE, data, 0744)
	if err != nil {
		fmt.Println(err)
	}
}

func ReadFromJSON() *Storage {
	file, err := os.ReadFile(STORAGE)
	if err != nil {
		return emptyStorage()
	}
	var storage Storage
	err = json.Unmarshal(file, &storage)
	if err != nil {
		fmt.Println("Не удалось разобрать JSON файл")
		return emptyStorage()
	}
	return &storage
}

func emptyStorage() *Storage {
	return &Storage{
		Bins: []bins.Bin{},
	}
}
