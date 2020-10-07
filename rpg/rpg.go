package rpg

import (
	"fmt"
	"math/rand"
	"time"
)

var monsterKilled int

var monsters = []Enemy{
	Enemy{name: "Skull-in", hp: 20, strength: 5},
	Enemy{name: "Werebat", hp: 12, strength: 4},
	Enemy{name: "Deadtree", hp: 6, strength: 3},
}

// ToBattle starts the battle
func ToBattle() {
	fmt.Printf("To battle!!!\n\n")

	rand.Seed(time.Now().UnixNano())

	hero := Hero{
		name:     "Odin",
		hp:       30,
		strength: 6,
	}

	weapon := Weapon{
		name:   "Sword",
		damage: 3,
	}

	hero.equip(weapon)

	for hero.isAlive() {
		monster := getRandomMonster()

		for monster.isAlive() && hero.isAlive() {
			hero.attack(&monster)
			fmt.Printf("%s\n\n", monster)

			if !monster.isAlive() {
				fmt.Printf("-------------------------\n\n")
				continue
			}

			monster.attack(&hero)
			fmt.Printf("%v\n\n", hero)
		}

		monsterKilled++
	}

	fmt.Printf("%v have killed %v monsters!\n\n", hero.name, monsterKilled)
}

func getRandomMonster() Enemy {
	n := rand.Intn(len(monsters))
	monster := monsters[n]

	fmt.Printf("A %s appeared!!!\n", monster.name)
	fmt.Printf("%v\n\n", monster)

	return monster
}
