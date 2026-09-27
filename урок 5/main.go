package main

import "fmt"

func main() {
	// Задание 1
	schooling := 9
	fmt.Println(schooling)

	schooling = 10
	fmt.Println(schooling)

	// Задание 2
	name := "Vladislav"
	fmt.Println(name)

	name = "Amir"
	fmt.Println(name)

	// Задание 3
	steps := 0
	fmt.Println(steps)

	steps = 2000
	fmt.Println(steps)
	fmt.Println("Хорошая работа! Вы уже на пути к своей ежедневной цели")

	// Задание 4
	largeNumber := 5000001
	fmt.Println(largeNumber)

	// Задание 5
	const breakTime = 15
	fmt.Println(breakTime)

	// breakTime = 20
	// Нельзя изменить значение константы,
	// потому что const нельзя переопределить после объявления.
}