package main

import "fmt"

type Notifier interface {
	Send(msg string)
}

type Email struct {
	Address string
}

type Telegram struct {
	Username string
}

func (e Email) Send(msg string) {
	fmt.Println("Отправка на почту ", e.Address, ":", msg)
}

func (t Telegram) Send(msg string) {
	fmt.Println("Отправка в ТГ ", t.Username, ":", msg)
}

func NotifyAll(notifyers []Notifier, text string) {
	for _, n := range notifyers {
		n.Send(text)
	}
}

func main() {
	list := []Notifier{Email{
		Address: "Ivan@gmail.ru",
	}, Telegram{
		Username: "@ivan_golang",
	}}

	NotifyAll(list, "Ваш заказ готов!")

	fmt.Println(list)
}
