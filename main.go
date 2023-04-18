package main

import (
	"fmt"
)

func main() {
	fmt.Println("Введите число, а я узнаю, четное оно или нет:")
	var number int
	fmt.Scan(&number)

	if number%2 == 0 {
		fmt.Println("Число чётное")
	} else {
		fmt.Println("Число нечётное")
	}
}
