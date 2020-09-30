package main

import (
	"fmt"
	"math/rand"
	"time"
)

var monsterKilled int

func main() {
	rand.Seed(time.Now().UnixNano())

	hero := hero{
		name:     "Conan",
		hp:       30,
		strength: 6,
	}

	weapon := weapon{
		name:   "Sword",
		damage: 3,
	}

	hero.equip(weapon)

	for hero.isAlive() {
		monster := getMoster()

		for monster.isAlive() && hero.isAlive() {
			hero.attack(&monster)
			fmt.Printf("%s\n\n", monster)

			if !monster.isAlive() {
				fmt.Printf("-------------------------\n\n")
			}

			if monster.isAlive() {
				monster.attack(&hero)
				fmt.Printf("%v\n\n", hero)
			}
		}

		monsterKilled++
	}

	fmt.Printf("%v have killed %v monsters!\n\n", hero.name, monsterKilled)
}

func getMoster() enemy {
	minHp := 8
	maxHp := 12
	minStrenght := 4
	maxStrenght := 6

	monster := enemy{
		name:     "Scorpion",
		hp:       rand.Intn(maxHp-minHp+1) + minHp,
		strength: rand.Intn(maxStrenght-minStrenght+1) + minStrenght,
	}

	fmt.Printf("A %s appeared!!!\n", monster.name)
	fmt.Printf("%v\n\n", monster)

	return monster
}
