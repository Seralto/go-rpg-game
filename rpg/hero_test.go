package rpg

import (
	"fmt"
	"testing"
)

func TestAttack(t *testing.T) {
	enemy := Enemy{name: "Enemy", hp: 10, strength: 5}
	hero := Hero{name: "Hero", hp: 10, strength: 5}

	hero.attack(&enemy)

	if enemy.hp == 10 {
		t.Errorf("Attack must have damage and decrease hp")
	}
}

func TestIsAlive(t *testing.T) {
	hero := Hero{name: "Hero", hp: 10, strength: 5}

	if hero.isAlive() != true {
		t.Errorf("New hero must be alive")
	}
}

func TestEquip(t *testing.T) {
	weapon := Weapon{name: "Weapon", damage: 5}
	hero := Hero{name: "Hero", hp: 10, strength: 5}

	hero.equip(weapon)

	if hero.weapon.name != "Weapon" || hero.weapon.damage != 5 {
		t.Errorf("Expected hero to be equipped with the weapon")
	}
}

func TestString(t *testing.T) {
	hero := Hero{name: "Hero", hp: 10, strength: 5}
	expected := fmt.Sprintf("%s", hero)

	if expected != "Hero: hp=10" {
		t.Errorf("Expected hero to be created")
	}
}

func TestDamage(t *testing.T) {
	hero := Hero{name: "Hero", hp: 10, strength: 5}
	expected := hero.damage()

	if expected <= 0 {
		t.Errorf("Damage should be positive")
	}
}
