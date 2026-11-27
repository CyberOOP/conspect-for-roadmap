package main

import (
	"study/methods"
	"study/pay"

	"github.com/k0kubun/pp"
)

func main() {
	method := methods.NewPaypal()
	payments := pay.NewPayment(method)

	payments.Pay("Бургер", 100)

	idgame := payments.Pay("Игра", 20)

	payments.CancelPay(idgame)

	allinfo := payments.AllInfo()

	pp.Println(allinfo)

	gameinfo := payments.Info(idgame)

	pp.Println(gameinfo)
}
