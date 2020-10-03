package rpg

import (
	"testing"
)

var myWeapon weapon = weapon{
	name:   "Weapon",
	damage: 5,
}

func TestWeapon(t *testing.T) {
	if myWeapon.name != "Weapon" {
		t.Error("The weapon name should have a name")
	}

	if myWeapon.damage != 5 {
		t.Error("The weapon name should have a damage")
	}
}
