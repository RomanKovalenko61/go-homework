package main

import (
	"app/hw-3/bins"
	"app/hw-3/file"
	"app/hw-3/storage"
	"encoding/json"
	"fmt"
)

func main() {
	fmt.Println("project 3-bin")
	storage := storage.ReadFromJSON()
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

func createBin(storage *storage.Storage) {
	fmt.Println("Создание бина")
	var id, name string
	fmt.Println("Введите id")
	fmt.Scan(&id)
	fmt.Println("Введите имя бина")
	fmt.Scan(&name)
	newBin := bins.NewBin(id, name)
	storage.AddBin(*newBin)
	storage.WriteToJSON()
}

func readAllBins(storage *storage.Storage) {
	for _, bin := range storage.Bins {
		fmt.Println("bin: ", bin)
	}
}

func readAndPrint() {
	fmt.Println("Укажите путь к фаилу для чтения")
	var path string
	fmt.Scan(&path)
	data, err := file.ReadFile(path)
	if err != nil {
		fmt.Println(err)
	}
	var storage storage.Storage
	json.Unmarshal(data, &storage)
	fmt.Println(storage)
}
