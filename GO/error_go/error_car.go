package main

import (
	"errors"
	"fmt"
	"math/rand"
)

type Car struct {
	Armor int
}

func (c *Car) Gas() (int, error) { //если вернем чисто ноль км/ч пользователь не поймет из за чего он не смог разогнаться, поэтому через errors добавляем в return ошибку тем самым возвращая число разгона и ошибку (int, error)
	if c.Armor-10 <= 0 {
		return 0, errors.New("Ваша прочность очень мала! Машина сломается при разгоне!")
	}

	kmch := rand.Intn(150)

	c.Armor -= 10

	return kmch, nil //если объявили две возвращаемые переменные то возвращаем их через запятую (если объявляли два возвращения то возвращаем два значения)
}

func main() {
	car := Car{
		20,
	}

	for {
		fmt.Println("Armor до: ", car.Armor)
		gas, err := car.Gas()
		fmt.Println("Armor после: ", car.Armor)
		if err != nil {
			fmt.Println("Ошибка! Причина:", err.Error())
			fmt.Println("Разгон: ", gas)
			break
		}
		fmt.Println("Разгон: ", gas)
		fmt.Println("")
	}

	/*
		└──╼ $go run car.go
		Armor до:  20
		Armor после:  10
		Разгон:  45

		Armor до:  10
		Armor после:  10
		Ошибка! Причина: Ваша прочность очень мала! Машина сломается при разгоне!
		Разгон:  0
	*/
}
