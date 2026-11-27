package main

import "fmt"

func main() {
	defer func() {
		p := recover() //если была паника в программе, то в конце функции паника передается в "recover()" и оттуда проверяем, если не нил то выводится что была паника и выводится сама паника
		if p != nil {
			fmt.Println("Была паника: ", p)
		}
	}()

	a := 5
	b := 0
	c := a / b
	fmt.Println("a / b =", c)

}
