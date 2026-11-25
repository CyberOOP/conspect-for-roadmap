package methods

import (
	"fmt"
	"math/rand"
)

type Paypal struct{}

func NewPaypal() *Paypal {
	return &Paypal{}
}

func (c Paypal) Pay(rub int) int {
	fmt.Println("Оплата через банк!")
	fmt.Println("Размер оплаты: ", rub)

	return rand.Int()
}

func (c Paypal) CancelPay(id int) {
	fmt.Println("Отмена операции! ID: ", id)
}
