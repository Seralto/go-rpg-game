package rpg

import (
	"fmt"
	"math/rand"
)

type Enemy struct {
	name     string
	hp       int
	strength int
}

func (e Enemy) attack(target *Hero) {
	roundDamage := e.damage()
	fmt.Printf("%s attacked %s, the damage was %d!\n", e.name, target.name, roundDamage)
	target.hp -= roundDamage
}

func (e Enemy) isAlive() bool {
	return e.hp > 0
}

func (e Enemy) String() string {
	if e.isAlive() {
		return fmt.Sprintf("%s: hp=%d, strength=%d", e.name, e.hp, e.strength)
	}

	return fmt.Sprintf("%v is dead!", e.name)
}

func (e Enemy) damage() int {
	min := e.strength / 2
	max := e.strength
	return rand.Intn(max-min+1) + min
}
