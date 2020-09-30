package main

import (
	"fmt"
	"math/rand"
)

type enemy struct {
	name     string
	hp       int
	strength int
}

func (e enemy) attack(target *hero) {
	roundDamage := e.damage()
	fmt.Printf("%s attacked %s, the damage was %d!\n", e.name, target.name, roundDamage)
	target.hp -= roundDamage
}

func (e enemy) isAlive() bool {
	return e.hp > 0
}

func (e enemy) String() string {
	if e.isAlive() {
		return fmt.Sprintf("%s: hp=%d, strength=%d", e.name, e.hp, e.strength)
	}

	return fmt.Sprintf("%v is dead!", e.name)
}

func (e enemy) damage() int {
	min := e.strength / 2
	max := e.strength
	return rand.Intn(max-min+1) + min
}
