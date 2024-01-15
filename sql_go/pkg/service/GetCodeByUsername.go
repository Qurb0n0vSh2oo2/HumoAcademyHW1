package service

import "fmt"

func (s *Service) GetCodeByUsername() {
	var username string

	fmt.Println("Введите имя пользователя")
	fmt.Scan(&username)



	id, err := s.repo.GetUserIdByName(username)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("ЭТО ТОЧНО ТЫ ИЛИ БЛЭФУЕШЬ?")
	otvet := "да"
	fmt.Scan(&otvet)
	

	getCode, err := s.repo.GetCodeByUserID(id)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Вот ваш код: ", getCode)



}