package main

import (
	"errors"
	"fmt"

	"github.com/k0kubun/pp"
)

type User struct {
	Name    string
	Balance int //$
}

func Pay(user *User, usd int) error { //при использовании errors мы понимаем что должно возвразаться ничего другого кроме ошибки, типо если возвращали бы строку то мы бы забыли например что должно было возвращаться
	if user.Balance-usd < 0 {
		//return("Недостаточно средств")
		return errors.New("Недостаточно средств!")
	}

	user.Balance -= usd
	return nil
}

func main() {
	user := User{
		Name:    "Гоша",
		Balance: 2,
	}

	pp.Println("Ваш счёт до: ", user)
	err := Pay(&user, 5)
	pp.Println("Ваш счёт после: ", user)

	if err != nil {
		fmt.Println("Оплаты не было! Причина: ", err.Error())
	} else {
		fmt.Println("Оплата была произведена!")
	}
}
