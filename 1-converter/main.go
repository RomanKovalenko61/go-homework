package main

import (
	"fmt"
	"strings"
)

const USDToEURO = 1.1
const USDToRUB = 85.0

func main()  {
	fmt.Println("___ Конвертер валют ___")
	var count float64
	var source string
	var target string
	readData(&count, &source, &target)
	currency := converter(count, source, target)
	fmt.Printf("Получаем в итоге =  %5.2f %s", currency, target)
}

func readData(p_count *float64, source *string, target *string) {
	fmt.Println("Введите количество валюты")
	fmt.Scan(p_count)
	fmt.Println("Выберете исходную валюту RUB - Рубли USD - Доллары EURO - Евро")
	fmt.Scan(source)
	*source = strings.ToUpper(*source)
	fmt.Println("Выберете в какую валюту конверитровать RUB - Рубли USD - Доллары EURO - Евро")
	fmt.Scan(target)
	*target = strings.ToUpper(*target)
}

func converter(count float64, source string, target string) float64 {
	switch {
	case source == "RUB" && target == "USD":
		return count / USDToRUB
	case source == "RUB" && target == "EURO":
		return count / USDToRUB / USDToEURO
	case source == "USD" && target == "RUB":
		return count * USDToRUB
	case source == "USD" && target == "EURO":
		return count / USDToEURO
	case source == "EURO" && target == "RUB":
		return count * USDToRUB * USDToEURO 
	case source == "EURO" && target == "USD":
		return count * USDToEURO
	default: return count
	}
}