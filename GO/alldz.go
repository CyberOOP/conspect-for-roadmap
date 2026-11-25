package homework

import (
	"fmt"

	"github.com/k0kubun/pp"
)

type Account struct {
	Owner   string
	Balance int
	History []string
}

type Transaction struct {
	From   string
	To     string
	Amount int
}

func Transfer(accounts map[string]*Account, tx Transaction) {
	defer func() {
		fmt.Println("---Операция обработана---")
	}()

	sender, ok := accounts[tx.From]
	if !ok {
		pp.Println("Ошибка! Счет не найден!")
		return
	}

	receiver, ok := accounts[tx.To]
	if !ok {
		pp.Println("Ошибка! Счет не найден!")
		return
	}

	pp.Println("Пользователи существуют! Проводим операцию...")

	if tx.Amount > 0 {
		if sender.Balance >= tx.Amount {
			sender.Balance -= tx.Amount
			receiver.Balance += tx.Amount
			sender.History = append(sender.History, fmt.Sprintf("Перевод отправлен на сумму: %d", tx.Amount))
			receiver.History = append(receiver.History, fmt.Sprintf("Пополнение на сумму: %d", tx.Amount))
		} else {
			pp.Println("Ошибка! Недостаточно средств!")
		}
	}
}

func Bank() {
	accounts := map[string]*Account{
		"Алиса": {Owner: "Алиса", Balance: 500, History: []string{}},
		"Даня":  {Owner: "Даня", Balance: 1000, History: []string{}},
		"Вадим": {Owner: "Вадим", Balance: 1500, History: []string{}},
	}

	txs := []Transaction{
		{From: "Алиса", To: "Даня", Amount: 600},
		{From: "Даня", To: "Алиса", Amount: 700},
		{From: "Вадим", To: "Даня", Amount: 1000},
	}

	for _, tx := range txs {
		Transfer(accounts, tx)
	}

	pp.Println(accounts)

}
