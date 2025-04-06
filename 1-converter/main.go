package main

import (
	"fmt"
)

const USDToEURO = 1.1
const USDToRUB = 85.0

func main()  {
	fmt.Println("___ Конвертер валют ___")
	var count float64
	fmt.Println("Введите количество рублей")
	fmt.Scan(&count)
	fmt.Println("Сумма в долларах = ", count / USDToRUB)
	fmt.Println("Сумма в евро = ", count / USDToRUB / USDToEURO)
}