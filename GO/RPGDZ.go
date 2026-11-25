package homework

import "fmt"

type Weapon struct {
	name       string //название оружия
	damage     int    //урон оружия
	durability int    //прочность оружия
	IsBroken   bool   //флаг поломки оружия
}

func NewWeapon(
	name string,
	damage int,
	durability int,
) *Weapon {
	IsBroken := false
	if name == "" {
		name = "Ржавый кинжал"
		fmt.Println("Имя по умолчанию: ", name)
	}
	if damage <= 0 {
		damage = 5
		fmt.Println("Установлен минимальный урон: ", damage)
	}
	if durability > 100 {
		durability = 100
		fmt.Println("Прочность превышена! Установлена прочность: ", durability)
	} else if durability <= 0 {
		durability = 0
		IsBroken = true
		fmt.Println("Оружие повреждено!")
		fmt.Println("Установлено минимальное значение прочности! Меньше нуля или ноль!", durability)
	}

	return &Weapon{
		name:       name,
		damage:     damage,
		durability: durability,
		IsBroken:   IsBroken,
	}

}

func (w *Weapon) Attack() int {
	defer func() {
		if w.durability <= 0 && !w.IsBroken {
			w.durability = 0
			w.IsBroken = true
			fmt.Println("Внимание! ", w.name, "Сломалось после удара!")
		}
	}()

	if w.IsBroken == true {
		fmt.Println("Оружие сломано! Удар невозможен!")
		return 0
	}
	w.durability -= 15
	return w.damage
}

func (w *Weapon) Repair(points int) {
	if points <= 0 {
		return
	}
	w.durability += points
	if w.durability > 100 {
		w.durability = 100
	}
	if w.durability > 0 {
		w.IsBroken = false
	}
}

func boba() {
	weapon := NewWeapon(
		"Экскалибур",
		30,
		40,
	)

	fmt.Println("Weapon: ", weapon)
	totalDamage := 0
	for round := 1; round <= 6; round++ {
		fmt.Println("---Раунд", round, "---")
		if round == 4 {
			weapon.Repair(35)
			fmt.Println("Меч отремонтирован!")
		}

		dmg := weapon.Attack()
		totalDamage += dmg
		fmt.Println("Итоговый урон за раунд: ", totalDamage)
	}
	fmt.Println("Урон за 6 раундов: ", totalDamage)
}
