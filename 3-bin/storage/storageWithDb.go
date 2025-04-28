package storage

import (
	"app/hw-3/bins"
	"encoding/json"
	"fmt"
)

type Db interface {
	Write([]byte)
	Read() ([]byte, error)
}

type StorageWithDb struct {
	Storage
	Db Db
}

func GetStorageWithDb(db Db) *StorageWithDb {
	data, err := db.Read()
	if err != nil {
		fmt.Println("Не смог считать данные из DB")
		return &StorageWithDb{
			Storage: Storage{
				Bins: []bins.Bin{},
			},
			Db: db,
		}
	}
	var storage Storage
	err = json.Unmarshal(data, &storage)
	if err != nil {
		fmt.Println("Не удалось разобрать JSON")
		return &StorageWithDb{
			Storage: Storage{
				Bins: []bins.Bin{},
			},
			Db: db,
		}
	}
	return &StorageWithDb{
		Storage: storage,
		Db:      db,
	}
}

func (db *StorageWithDb) Save() {
	data, err := db.Storage.ToBytes()
	if err != nil {
		fmt.Println(err)
	}
	db.Db.Write(data)
}
