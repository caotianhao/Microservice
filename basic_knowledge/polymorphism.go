// polymorphism 多态
package main

import "fmt"

type Hero interface {
	FindWeapon()
	FindPerl()
	ShowInfo()
}

type LowLevel struct {
	Atk, Def int
}

type HighLevel struct {
	Atk, Def int
}

func (hero *LowLevel) FindWeapon() {
	hero.Atk += 50
}

func (hero *HighLevel) FindWeapon() {
	hero.Atk += 500
}

func (hero *LowLevel) FindPerl() {
	hero.Def += 50
}

func (hero *HighLevel) FindPerl() {
	hero.Def += 500
}

func (hero *LowLevel) ShowInfo() {
	fmt.Printf("atk=%d def=%d\n", hero.Atk, hero.Def)
}

func (hero *HighLevel) ShowInfo() {
	fmt.Printf("atk=%d def=%d\n", hero.Atk, hero.Def)
}

// 多态：传入不同参数得到不同结果
func changeAttack(r Hero) {
	r.FindWeapon()
}

func main() {
	nvWa := &LowLevel{Atk: 30, Def: 20}
	nvWa.ShowInfo()
	nvWa.FindWeapon()
	nvWa.FindWeapon()
	nvWa.FindPerl()
	nvWa.ShowInfo()

	panGu := &HighLevel{Atk: 90, Def: 70}
	panGu.ShowInfo()
	changeAttack(panGu)
	panGu.ShowInfo()
}
