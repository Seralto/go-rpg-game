package main

import (
	"fmt"
	"testing"
)

var myHero hero = hero{
	name:     "Hero",
	hp:       10,
	strength: 5,
	exp:      0,
}

var theEnemy enemy = enemy{
	name:     "Enemy",
	hp:       10,
	strength: 5,
}

var theWeapon weapon = weapon{
	name:   "Weapon",
	damage: 5,
}

func TestAttack(t *testing.T) {
	myHero.attack(&theEnemy)

	if theEnemy.hp == 10 {
		t.Errorf("Attack must have damage and decrease hp")
	}
}

func TestIsAlive(t *testing.T) {
	if myHero.isAlive() != true {
		t.Errorf("New hero must be alive")
	}
}

func TestEquip(t *testing.T) {
	myHero.equip(theWeapon)

	if myHero.weapon.name != "Weapon" || myHero.weapon.damage != 5 {
		t.Errorf("Expected hero to be equipped with the weapon")
	}
}

func TestString(t *testing.T) {
	r := fmt.Sprintf("%s", myHero)

	if r != "Hero: hp=10" {
		t.Errorf("Expected hero to be created")
	}
}

func TestDamage(t *testing.T) {
	r := myHero.damage()

	if r <= 0 {
		t.Errorf("Damage should be positive")
	}
}
