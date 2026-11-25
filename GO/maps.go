package homework

import (
	"github.com/k0kubun/pp"
)

type Product struct {
	name  string
	price int
	count int
}

func NewProduct(
	Name string,
	Price int,
	Count int,
) *Product {
	if Name == "" {
		pp.Println("Ошибка! Пустая строка!")
		return nil
	}
	if Price < 0 {
		pp.Println("Ошибка! Цена не может быть меньше нуля!")
		return nil
	}
	if Count < 0 {
		pp.Println("Ошибка! Количество не может быть меньше нуля!")
		return nil
	}

	return &Product{
		name:  Name,
		price: Price,
		count: Count,
	}
}

func UpdateStock(
	warehouse map[string]*Product, //все операции проводим с мапой warehouse где ключом является строка а возвращается то что лежит в ключе в структуре
	name string, //название товара
	addCount int, //сколько штук привезли
	newPrice int, //новая цена товарва
) {
	prod, ok := warehouse[name] //поиск товара на складе
	if ok {                     //если товар найден, те true
		prod.count += addCount //к изначальному количеству товара добавляется то количество которое мы передали
		if newPrice > 0 {      //если новая цена больше нуля
			prod.price = newPrice // присваиваем старой цене новую
		}
	} else {
		//товара нет
		newProd := NewProduct(name, newPrice, addCount) //через конструктор создаем товар
		if newProd != nil {                             //если товар не пустой, то в мапу присваивается этот пустой товар
			warehouse[name] = newProd //по типу ages["Аня"] = 26, те ключу Аня стринг присваивается возраст 26 значение
		}
	}
}

func Maps() {
	warehouse := make(map[string]*Product)

	UpdateStock(warehouse, "Яблоко", 10, 120)
	UpdateStock(warehouse, "Яблоко", 5, 200)
	UpdateStock(warehouse, "Груша", 15, 150)

	pp.Println(warehouse)
}
