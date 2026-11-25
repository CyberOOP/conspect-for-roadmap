package homework

import (
	"fmt"
	"strings"
)

func Sss() {
	// var name string
	// fmt.Print("Введите имя: ")
	// fmt.Scan(&name)
	// fmt.Printf("Привет, %s!", name)

	//для ввода пользователем строки из нескольких слов создаем сканер
	// scanner := bufio.NewScanner(os.Stdin)
	// fmt.Print("Введите фразу: ")
	//ждем пока пользователь нажмет enter
	// if scanner.Scan() {
	// 	text := scanner.Text() //Забираем из переменной scanner введенную строку
	// 	fmt.Println("Ты написал: ", text)
	// }

	// scanner := bufio.NewScanner(os.Stdin)
	// fmt.Print("Введите число: ")
	// if scanner.Scan() {
	// 	text := scanner.Text() Забираем из переменной scanner введенную строку
	// 	fmt.Println("Ты написал: ", text)

	// 	num, err := strconv.Atoi(text) проверяется ошибка, если ошибка не равна пустоте те ошибка не пустая, то выводится ошибка
	// 		return

	// 	}
	// 	fmt.Println("Число: ", num)
	// }

	text := "добавить зеленый лук"
	fields := strings.Fields(text) //strings.Fields решеит строку по пробелам и возвращает слайс
	if len(fields) == 0 {
		fmt.Println("Вы ничего не ввели!")
		return
	}
}
