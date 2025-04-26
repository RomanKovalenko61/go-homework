package main

import (
	"app/hw-3/api"
	"app/hw-3/bins"
	"app/hw-3/file"
	"app/hw-3/storage"
	"encoding/json"
	"fmt"
)

func main() {
	fmt.Println("project 3-bin")
	storage := storage.GetStorageWithDb(file.NewJsonDb("storage/storage.json"))
	api.WorkWithApi()
Menu:
	for {
		choice := getMenu()
		switch choice {
		case 1:
			createBin(storage)
		case 2:
			readAllBins(storage)
		case 3:
			readAndPrint()
		default:
			break Menu
		}
	}
}

func getMenu() int {
	fmt.Println("Выберите действие: 1 - добавить бин в хранилище 2 - прочитать хранилище  3 - тест функции file.ReadFile 4-выход")
	var choice int
	fmt.Scan(&choice)
	return choice
}

func createBin(storage *storage.StorageWithDb) {
	fmt.Println("Создание бина")
	var id, name string
	fmt.Println("Введите id")
	fmt.Scan(&id)
	fmt.Println("Введите имя бина")
	fmt.Scan(&name)
	newBin := bins.NewBin(id, name)
	storage.AddBin(*newBin)
	storage.Save()
}

func readAllBins(db *storage.StorageWithDb) {
	for _, bin := range db.Storage.Bins {
		fmt.Println("bin: ", bin)
	}
}

func readAndPrint() {
	fmt.Println("Укажите путь к фаилу для чтения")
	var path string
	fmt.Scan(&path)
	data, err := file.ReadFile(path)
	if err != nil {
		fmt.Println("Чтение не удалось", err)
	}
	var storage storage.Storage
	err = json.Unmarshal(data, &storage)
	if err != nil {
		fmt.Println("Проблема Анмаршалл ", err)
	}
	fmt.Println(storage)
}
