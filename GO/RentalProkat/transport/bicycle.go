package transport

import "fmt"

type bicycle struct{}

func NewBicycle() *bicycle {
	return &bicycle{}
}

func (b *bicycle) Summa(minutes int) int {
	sum := minutes * 8
	return sum
}

func (b *bicycle) Start() {
	fmt.Println("Я велосипед! Крути педали!")
	fmt.Println("Поездка началась! Приятной поездки!")
}

func (b *bicycle) Stop() {
	fmt.Println("Я велосипед! Я останавливаюсь!")
	fmt.Println("Спасибо за использование велосипедов нашей компании!")
}
