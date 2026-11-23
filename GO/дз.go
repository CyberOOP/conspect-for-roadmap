package main

import "fmt"

type Account struct {
	owner   string
	balance float64
}

func NewAccount(
	owner string,
	balance float64,
) *Account {
	if owner == "" {
		return &Account{}
	}

	if balance < 0 {
		return &Account{}
	}

	return &Account{
		owner:   owner,
		balance: balance,
	}
}

func (u *Account) Deposit(ammount float64) string {
	if ammount <= 0 {
		return "Пополнить отрицательный баланс нельзя!"
	}

	u.balance += ammount
	return "Пополнение успешно!"
}

func (u *Account) Withdraw(ammount float64) bool {
	if ammount <= 0 || u.balance < ammount {
		return false
	}

	u.balance -= ammount
	return true
}

func (a Account) info() {
	fmt.Println("Имя владельца: ", a.owner)
	fmt.Println("Ваш баланс: ", a.balance)
}

func main() {
	acc := NewAccount(
		"Сергей",
		100.0,
	)

	fmt.Println("Баланс изначально: ", acc.balance)
	fmt.Println("Баланс после депозита +100: ", acc.Deposit(100.0))
	fmt.Println("Баланс после депозита на -100: ", acc.Deposit(-100.0))
	fmt.Println("Баланс после снятия 50: ", acc.Withdraw(50))
	fmt.Println("Баланс после снятия 200:", acc.Withdraw(200))
	fmt.Println("Итоговый баланс: ", acc.balance)
}
