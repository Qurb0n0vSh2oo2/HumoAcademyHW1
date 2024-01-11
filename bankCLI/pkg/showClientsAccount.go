package pkg

import "fmt"

func ShowClientsAccount() {
	var name string

	fmt.Println("Введите имя клиента")

	fmt.Scan(&name)

	balance := 0
	has := false

	

	if !has {
		fmt.Println("Ошибка!, данного пользователя нет в нашей бд")
		return
	}

	fmt.Printf("Баланс счета %v является %v \n", name, balance)
}
