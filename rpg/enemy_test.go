package rpg

import (
	"fmt"
	"testing"
)

var aHero hero = hero{
	name:     "Hero",
	hp:       10,
	strength: 5,
}

var anEnemy enemy = enemy{
	name:     "Enemy",
	hp:       10,
	strength: 5,
}

func TestEnemyAttack(t *testing.T) {
	anEnemy.attack(&aHero)

	if aHero.hp == 10 {
		t.Errorf("Attack must have damage and decrease hp")
	}
}

func TestEnemyIsAlive(t *testing.T) {
	if anEnemy.isAlive() != true {
		t.Errorf("New enemy must be alive")
	}
}

func TestEnemyString(t *testing.T) {
	r := fmt.Sprintf("%s", anEnemy)

	if r != "Enemy: hp=10, strength=5" {
		t.Errorf("Expected hero")
	}
}

func TestEnemyDamage(t *testing.T) {
	r := anEnemy.damage()

	if r <= 0 {
		t.Errorf("Damage should be positive")
	}
}
