package main

import (
	"fmt"
	"time"
)

var database = make(map[string]int)

func AddClient() {
	var name string
	//   var age int

	fmt.Scan(&name)

	//   fmt.Scan(&age)

	//   database[name] = 2023 - age

	fmt.Println("________________")
	fmt.Println("Готово")
	fmt.Println("________________")

}

func TopUptheBalance() {
	var name string
	var AddBalance int
	fmt.Scan(&name)
	fmt.Scan(&AddBalance)
	database[name] += AddBalance
	fmt.Println("______________")
}

func ShowBalance() {

	// for key, val := range database {
	// 	fmt.Println(key, val)
	// }
	var name string
	fmt.Scan(&name)
	fmt.Println(database[name])
	fmt.Println("______________")
}

func MinusBalance() {
	var name string
	var MinusBalance int
	fmt.Scan(&name)
	fmt.Scan(&MinusBalance)
	database[name] -= MinusBalance
	fmt.Println("______________")
}

func ShowAllClients() {
	for key, val := range database {
		fmt.Println(key, val)
	}
	fmt.Println("______________")
}

func main() {

	for {
		var choice int
		fmt.Println("1. Добавить клиента")
		fmt.Println("2. Пополнить баланс клиента")
		fmt.Println("3. Показать баланс клиента")
		fmt.Println("4. Снять деньги клиента")
		fmt.Println("5. Показать всех клиентов")

		fmt.Scan(&choice)

		if choice == 1 {
			AddClient()
		} else if choice == 2 {
			TopUptheBalance()
		} else if choice == 3 {
			ShowBalance()
		} else if choice == 4 {
			MinusBalance()
		} else if choice == 5 {
			ShowAllClients()
		}

		time.Sleep(2 * time.Second)

	}

}
