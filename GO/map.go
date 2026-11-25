package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/k0kubun/pp"
)

func main() {
	//Введите команду добавить лук
	//Вы хотите добавить лук

	//Введите команду удалить морковь
	//Вы хотите удалить морковь

	scanner := bufio.NewScanner(os.Stdin) //stdin входной поток из консоли

	for {
		fmt.Print("Введите команду: ")

		// ok := scanner.Scan()
		// if !ok {
		// 	fmt.Println("Ошибка!")
		// 	return
		// }

		if ok := scanner.Scan(); !ok {
			fmt.Println("Ошибка!")
			return
		}

		text := scanner.Text()

		fields := strings.Fields(text)

		if len(fields) == 0 {
			fmt.Println("Вы ничего не ввели!")
			return
		}

		pp.Println("Text: ", text)
		pp.Println("Слова: ", fields)

		fmt.Println("Команда: ", fields[0])

		cmd := strings.ToLower(fields[0])

		//strings.ReplaceAll(s, ",", " ") заменяет все запятые на пробелы перед тем, как разбивать строку на слова.
		//strings.Trim(s, ",:!?") обрезает указанные символы по краям строки.
		//strings.Split(s, ",") режет строку строго по указанному разделителю (например, по запятой), а не по пробелам.
		//strconv.Atoi(s) — превращает строку в целое число (int) и возвращает ошибку, если пользователь ввел буквы вместо цифр (например, "двадцать").
		//strconv.ParseFloat(s, 64) — превращает строку в дробное число (float64).
		//fmt.Printf("Привет %s! Тебе %d лет.", "Анна", 25)
		//%v если не уверены в формате данных
		//%T показывает тип переменной
		//%d для целых чисел
		//%s для строк
		//%f для чисел с плавающей точкой (%.2f) до 2 знаков после запятой

		if cmd == "выйти" {

			fmt.Println("До встречи!")
			return
		}
		if cmd == "добавить" {
			str := ""
			for i := 1; i < len(fields); i++ { //перебор идет со второго слова, тк первое слово это команда (добавить или удалить) пошли с первого индекса и собрали слова в строчку
				str += fields[i]

				if i != len(fields)-1 { //будем добавлять пробелы между словами кроме последнего слова
					str += " "
				}
			}
			fmt.Println("Вы хотите добавить: ", str)
		} else if cmd == "удалить" {
			str1 := ""
			for i := 1; i < len(fields); i++ {
				str1 += fields[i]

				if i != len(fields)-1 {
					str1 += " "
				}
			}
			fmt.Println("Вы хотите удалить: ", str1)
		} else if cmd == "help" {
			fmt.Println("Команда добавить (что нужно добавить)")
			fmt.Println("-- Эта команда позволяет добавить что-то")
			fmt.Println("Команда удалить (что нужно удалить)")
			fmt.Println("-- Эта команда позволяет удалить что-то")
		} else {
			fmt.Println("Неизвестная команда!")
		}
	}
}
