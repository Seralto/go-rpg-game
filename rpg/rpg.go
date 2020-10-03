package rpg

import (
	"fmt"
	"math/rand"
	"time"
)

var monsterKilled int

var monsters = []enemy{
	enemy{name: "Skull-in", hp: 20, strength: 5},
	enemy{name: "Werebat", hp: 12, strength: 4},
	enemy{name: "Deadtree", hp: 6, strength: 3},
}

// ToBattle starts the battle
func ToBattle() {
	fmt.Printf("To battle!!!\n\n")

	rand.Seed(time.Now().UnixNano())

	hero := hero{
		name:     "Odin",
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
	n := rand.Intn(len(monsters))
	monster := monsters[n]

	fmt.Printf("A %s appeared!!!\n", monster.name)
	fmt.Printf("%v\n\n", monster)

	return monster
}
