package main

import (
	"github.com/k0kubun/pp"
)

type User struct {
	Name    string
	Rating  float64
	Premium bool
}

func main() {

	userArray := [3]User{
		User{
			Name:    "Даня",
			Rating:  5.5,
			Premium: true,
		},

		User{
			Name:    "Вася",
			Rating:  7.5,
			Premium: true,
		},

		User{
			Name:    "Костя",
			Rating:  7.7,
			Premium: false,
		},
	}

	pp.Println("---До---")
	/*for i := 0; i < len(userArray); i++ { перебор старым способом
		pp.Println(userArray[i])
	}*/

	for index, user := range userArray { //пробегись по массиву и для каждого элемента выполни данную итерацию, на каждую итерацию клади индекс элемента и его значение
		pp.Println(index, user)
	}

	/*если мы хотим получить только значения элементов массива

	for _, user := range userArray {
		pp.Println(user)
	}
	fmt.Println("")
	*/

	/*
		for i := 0; i < len(userArray); i++ {
			if userArray[i].Premium {
				userArray[i].Rating += 1.0
			}
		}
	*/

	//чтобы работать не с копиям (тк мы создаем две новые переменные в них копируется значение), то в основном массиве ничего не изменится
	//поэтому делаем так

	for index, user := range userArray {
		if user.Premium {
			userArray[index].Rating += 1.0

		}
	}

	pp.Println("---После---")
	for _, u := range userArray {
		pp.Println(u)
	}
	//все у кого есть Premium, прибавляем единицу
}
