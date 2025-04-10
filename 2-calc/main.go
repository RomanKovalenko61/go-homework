package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("___ Работа с слайсами ___")
	operation := getOperation()
	input := getInputString()
	numbers := mapToSliceInt(input)
	result := calc(operation, numbers)
	fmt.Printf("Операция: %v результат: %.2f", operation, result)
}

func getOperation() string {
	fmt.Println("Введите операцию AVG - среднее, SUM - сумма, MED - медиана")
	var op string
	fmt.Scanln(&op)
	op = strings.ToUpper(op)
	return op
}

func getInputString() string {
	fmt.Println("Введите значения через запятую для выполнения операции. В конце нажмите Enter")
	var input string
	fmt.Scanln(&input)
	return input
}

func mapToSliceInt(input string) []int {
	strings_arr := strings.Split(input, ",")
	var numbers []int
	for _, value := range strings_arr {
		temp, _ := strconv.Atoi(value)
		numbers = append(numbers, temp)  
	}
	return numbers
}

func calc(operation string, array []int) float64 {
	switch operation {
	case "AVG": 
		return avg(array)
	case "SUM": 
		return sum(array)
	case "MED": 
		sort.Ints(array)
		return med(array)
	default: panic("Unknown operation")
	}
}

func sum(array []int) float64 {
	var sum int
	for _, value := range array {
		sum += value
	}
	return float64(sum)
}

func avg(array []int) float64 {
	return sum(array) / float64(len(array))
}

func med(array []int) float64 {
	len := len(array)
	if len % 2 != 0 {
		return float64(array[len / 2])
	}
	sum := array[len / 2] + array[len/ 2 - 1]
	return float64(sum) / 2.0
}