package methods

import (
	"fmt"
	"math/rand"
)

type Crypto struct{}

func NewCrypto() *Crypto {
	return &Crypto{}
}

func (c Crypto) Pay(rub int) int {
	fmt.Println("Оплата криптовалютой!")
	fmt.Println("Размер оплаты: ", rub)

	return rand.Int()
}

func (c Crypto) CancelPay(id int) {
	fmt.Println("Отмена криптооперации! ID: ", id)
}
