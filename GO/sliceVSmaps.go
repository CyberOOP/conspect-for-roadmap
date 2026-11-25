package main

import (
	"fmt"
	"strconv"
	"time"
)

// slice VS map
// надо сохранять больше количество информации об оплатах
// нужно находить конкретный элемент информации об оплате, например нужно найти конкретный ID или найти не отмененную операцию
type PaymentInfo struct {
	ID        int
	name      string
	usd       int
	cancelled bool
}

type PaymentInfoSlice struct {
	s []PaymentInfo
}

func (s *PaymentInfoSlice) AddInfo(info PaymentInfo) {
	s.s = append(s.s, info)
}

func (s *PaymentInfoSlice) GetInfo(id int) PaymentInfo {
	for _, info := range s.s {
		if info.ID == id {
			return info
		}
	}

	return PaymentInfo{}
}

//---------------------------------------------------

type PaymentInfoMap struct {
	m map[int]PaymentInfo
}

func (m *PaymentInfoMap) AddInfo(info PaymentInfo) {
	m.m[info.ID] = info
}

func (m *PaymentInfoMap) GetInfo(id int) PaymentInfo {
	info, ok := m.m[id]
	if !ok {
		return PaymentInfo{}
	}

	return info
}

func main() {

	pSlice := PaymentInfoSlice{}
	pMap := PaymentInfoMap{
		m: make(map[int]PaymentInfo),
	}

	iterations := 10000000

	before := time.Now()

	for i := 0; i <= iterations; i++ {
		info := PaymentInfo{
			ID:   i,
			name: strconv.Itoa(i),
		}

		pSlice.AddInfo(info)
	}

	fmt.Println("slice add: ", time.Since(before))

	before = time.Now()

	for i := 0; i <= iterations; i++ {
		info := PaymentInfo{
			ID:   i,
			name: strconv.Itoa(i),
		}

		pMap.AddInfo(info)
	}

	fmt.Println("map add: ", time.Since(before))

	//------------------------------------

	before = time.Now()

	for i := 0; i <= iterations; i++ {
		info := pSlice.GetInfo(i)
		_ = info
	}

	fmt.Println("slice get: ", time.Since(before))

	before = time.Now()

	for i := 0; i <= iterations; i++ {
		info := pMap.GetInfo(i)
		_ = info
	}

	fmt.Println("map get: ", time.Since(before))
}
