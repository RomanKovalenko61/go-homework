package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

var actions = map[string]func([]int) float64{
	"AVG": avg,
	"SUM": sum,
	"MED": med,
}

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
		temp, err := strconv.Atoi(value)
		if err != nil {
			log.Println("Can't convert to int ", value)
		} else {
			numbers = append(numbers, temp)
		}
	}
	return numbers
}

func calc(operation string, array []int) float64 {
	op := actions[operation]
	if op != nil {
		return op(array)
	} else {
		log.Fatal("Unkown_operation ", operation)
		return 0
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
	length := len(array)
	if length%2 != 0 {
		return float64(array[length/2])
	}
	sum := array[length/2] + array[length/2-1]
	return float64(sum) / 2.0
}
