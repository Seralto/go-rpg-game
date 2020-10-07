package rpg

import (
	"testing"
)

func TestWeapon(t *testing.T) {
	weapon := Weapon{name: "Weapon", damage: 5}

	if weapon.name != "Weapon" {
		t.Error("The weapon should have a name")
	}

	if weapon.damage != 5 {
		t.Error("The weapon should have a damage")
	}
}
