package main

import (
	"fmt"
	"strings"
)

type currencyMap map[string]float64

func main()  {
	exchanger := currencyMap{"USD": 1, "RUB": 85.0, "EURO": 0.9}
	fmt.Println("___ Конвертер валют ___")
	countMoney, source, target := readData(&exchanger)
	currency := converter(countMoney, source, target, &exchanger)
	fmt.Printf("Конвертируем %v в %v Получаем в итоге =  %5.2f", source, target, currency)
}

func readData(exchanger *currencyMap) (float64, string, string) {
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

func choiceSource(exchanger *currencyMap) string {
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

func choiceTarget(source string, exchanger *currencyMap) string {
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

func isExistsCurrency(currency string, exchanger *currencyMap) bool {
	if _, key := (*exchanger)[currency]; key {
		return true		
	} else {
		return false
	}
}

func converter(countMoney float64, source string, target string, exchanger *currencyMap) float64 {
	toUSD := countMoney / (*exchanger)[source]
	to := (*exchanger)[target]
	return toUSD * to
}