package transport

import "fmt"

type samokat struct{}

func NewSamokat() *samokat {
	return &samokat{}
}

func (s *samokat) Summa(minutes int) int {
	sum := minutes * 10
	return sum
}

func (b *samokat) Start() {
	fmt.Println("Я  самокат! Не сбей пешеходов!")
	fmt.Println("Поездка началась! Приятной поездки!")
}

func (b *samokat) Stop() {
	fmt.Println("Я самокат! Я останавливаюсь!")
	fmt.Println("Спасибо за использование самокатов нашей компании!")
}
