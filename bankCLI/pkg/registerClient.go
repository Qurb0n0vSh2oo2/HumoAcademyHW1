package pkg

import (
	"bankCLI/pkg/models"
	"fmt"
)

func RegisterClient() {
	var name string
	var age int
	var phoneNumber string
	var cityName string
	var city models.City
	var has bool

	fmt.Println("Введите имя клиента")

	fmt.Scan(&name)

	fmt.Println("Введите возраст клиента")

	fmt.Scan(&age)

	fmt.Println("Введите номер телефона клиента")

	fmt.Scan(&phoneNumber)

	fmt.Println("Введите город клиента")

	fmt.Scan(&cityName)

	for _, v := range Cities {
		if v.City == cityName {
			city = v
			has = true
		}

	}

	if !has {
		fmt.Println("Город не найден в нашей бд")
		return
	}

	Clients = append(Clients, &models.Clients{
		Name:        name,
		Balance:     0,
		Age:         age,
		PhoneNumber: phoneNumber,
		City:        city,
	})

	fmt.Println("________________")
	fmt.Println("Готово")
	fmt.Println("________________")

}
