package rpg

import (
	"fmt"
	"math/rand"
)

type hero struct {
	name     string
	hp       int
	strength int
	exp      int
	weapon   weapon
}

func (h hero) attack(target *enemy) {
	roundDamage := h.damage()
	fmt.Printf("%s attacked the %s. The damage was %d!\n", h.name, target.name, roundDamage)
	target.hp -= roundDamage
}

func (h hero) isAlive() bool {
	return h.hp > 0
}

func (h *hero) equip(w weapon) {
	h.weapon = w
}

func (h hero) String() string {
	if h.isAlive() {
		return fmt.Sprintf("%s: hp=%d", h.name, h.hp)
	}

	return fmt.Sprintf("\n%s is dead.\nThe battle is over!", h.name)
}

func (h hero) damage() int {
	roundStrenght := rand.Intn(h.strength)
	if h.weapon.damage != 0 {
		return h.weapon.damage + roundStrenght
	}

	return roundStrenght
}
