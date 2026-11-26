//добавить обработку ошибок (errors)

package rental

import (
	"fmt"
	"math/rand"
)

type TS interface {
	Summa(minutes int) int
	Start()
	Stop()
}

type User struct {
	Name      string
	Balance   int
	HistoryID []int
}

type CardProkat struct {
	Name       string
	Minutes    int
	Summa      int
	StopProkat bool
}

type Manager struct {
	ts        TS
	MapUser   map[string]*User
	MapProkat map[int]*CardProkat
}

func NewManager(ts TS) *Manager {
	return &Manager{
		ts:        ts,
		MapUser:   make(map[string]*User),
		MapProkat: make(map[int]*CardProkat),
	}
}

func (m *Manager) AddUser(name string, balance int) {
	m.MapUser[name] = &User{
		Name:    name,
		Balance: balance,
	}
}

func (m *Manager) StartRent(username string, minutes int) int {
	user, ok := m.MapUser[username]
	if !ok {
		fmt.Println("Пользователя нет в базе данных!")
		m.AddUser(username, 0)
		fmt.Println("Пополните баланс и берите поездку! Удачи вам!)")
		return 0
	}
	fmt.Println("Происходит расчёт цены...")
	summa := m.ts.Summa(minutes)
	if user.Balance < summa {
		fmt.Println("Недостаточно средств!")
		return 0
	} else {
		user.Balance -= summa
		m.ts.Start()
		id := rand.Int()
		m.MapProkat[id] = &CardProkat{
			Name:       username,
			Minutes:    minutes,
			Summa:      summa,
			StopProkat: false,
		}
		fmt.Println("Деньги списаны! Поездка началась!")

		return id
	}
}

func (m *Manager) StopRent(id int) {
	card, ok := m.MapProkat[id]
	if !ok {
		fmt.Println("Такой поездки не существует!")
		return
	}

	if m.MapProkat[id].StopProkat {
		fmt.Println("Поездка уже завершена!")
		return
	}

	m.ts.Stop()
	card.StopProkat = true

	m.MapUser[card.Name].HistoryID = append(m.MapUser[card.Name].HistoryID, id)
}

func (m *Manager) GetInfo(id int) CardProkat {
	card, ok := m.MapProkat[id]
	if !ok {
		fmt.Println("Такой поездки не существует!")
		return CardProkat{}
	}

	return *card
}
