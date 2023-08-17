package main

import "fmt"

type WeaponType int

const (
	Axe WeaponType = iota
	Sword
	WoodenStick
	Knife
)

func (s WeaponType) String() string {
	switch s {
	case Axe:
		return "AXE"
	case Sword:
		return "SWORD"
	case WoodenStick:
		return "WOODEN STICK"
	case Knife:
		return "KNIFE"
	default:
		panic("not found")
	}
}

func getDamage(weapon WeaponType) int {
	switch weapon {
	case Axe:
		return 100
	case Sword:
		return 90
	case WoodenStick:
		return 1
	case Knife:
		return 40
	default:
		panic("weapon not supported")
	}
}

func main() {
	// Here we can use "%s" in a type int because we have a method called .String() which is called automatically by compiler
	fmt.Printf("Weapon (%s): has %d damage\n", Axe, getDamage(Axe))
	fmt.Printf("Weapon (%s): has %d damage\n", Sword, getDamage(Sword))
	fmt.Printf("Weapon (%s): has %d damage\n", WoodenStick, getDamage(WoodenStick))
	fmt.Printf("Weapon (%s): has %d damage\n	", Knife, getDamage(Knife))
}
