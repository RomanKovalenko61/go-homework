package storage

import (
	"app/hw-3/bins"
	"encoding/json"
	"fmt"
)

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

func (storage *Storage) ToBytes() ([]byte, error) {
	return json.Marshal(storage)
}
