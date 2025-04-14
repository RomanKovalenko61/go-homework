package main

import (
	"fmt"
	"slices"
	"strings"
)

const USDToEURO = 1.1
const USDToRUB = 85.0

func main()  {
	exchanger := []string{"USD", "RUB", "EURO"}
	fmt.Println("___ Конвертер валют ___")
	count, source, target := readData(exchanger)
	currency := converter(count, source, target)
	fmt.Printf("Конвертируем %v в %v Получаем в итоге =  %5.2f", source, target, currency)
}

func readData(exchanger []string) (float64, string, string) {
	countMoney := getCountMoney()
	source := choiceSource(exchanger)
	target := choiceTarget(source, exchanger)
	return countMoney, source, target
}

func getCountMoney() float64{
	fmt.Println("Введите количество валюты")
	var countMoney float64
	fmt.Scan(&countMoney)
	return countMoney
}

func choiceSource(exchanger []string) string {
	var source string
	for {
		fmt.Println("Выберете исходную валюту RUB - Рубли USD - Доллары EURO - Евро")
		fmt.Scan(&source)
		source = strings.ToUpper(source)
		if isExistsCurrency(source, exchanger) {
			break
		} else {
			fmt.Print("Вы ввели валюту которая отсутствует в списке")
		}
	}
	return source
}

func choiceTarget(source string, exchanger []string) string {
	var target string
	for {
		fmt.Println("Выберете в какую валюту конверитровать RUB - Рубли USD - Доллары EURO - Евро")
		fmt.Scan(&target)
		target = strings.ToUpper(target)
		if isExistsCurrency(target, exchanger) {
			if (source != target) {
				break
			} else {
				fmt.Println("Валюты совпадают")
			}
		} else {
			fmt.Print("Вы ввели валюту которая отсутствует в списке")
		}
	}
	return target
}

func isExistsCurrency(currency string, exchanger []string) bool {
	isExists := false
	if slices.Contains(exchanger, currency) {
			return true
		}
	return isExists
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