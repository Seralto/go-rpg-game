package rpg

import (
	"fmt"
	"testing"
)

func TestEnemyAttack(t *testing.T) {
	enemy := Enemy{name: "Enemy", hp: 10, strength: 5}
	hero := Hero{name: "Hero", hp: 10, strength: 5}

	enemy.attack(&hero)

	if hero.hp == 10 {
		t.Errorf("Attack must have damage and decrease hp")
	}
}

func TestEnemyIsAlive(t *testing.T) {
	enemy := Enemy{name: "Enemy", hp: 10, strength: 5}

	if enemy.isAlive() != true {
		t.Errorf("New enemy must be alive")
	}
}

func TestEnemyString(t *testing.T) {
	enemy := Enemy{name: "Enemy", hp: 10, strength: 5}
	expected := "Enemy: hp=10, strength=5"
	resp := fmt.Sprintf("%s", enemy)

	if resp != expected {
		t.Errorf("Error, expected %s, got %s", expected, resp)
	}
}

func TestEnemyDamage(t *testing.T) {
	enemy := Enemy{name: "Enemy", hp: 10, strength: 5}
	resp := enemy.damage()

	if resp <= 0 {
		t.Errorf("Damage should be positive")
	}
}
