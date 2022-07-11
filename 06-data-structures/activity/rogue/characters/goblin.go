package characters

import (
	"fmt"
	"math/rand"
)

type Goblin struct {
	Health       int
	AttackPower  int
	DefensePower int
}

func NewGoblin() Goblin {
	return Goblin{
		Health:       rand.Intn(6) + 5,
		AttackPower:  rand.Intn(2) + 2,
		DefensePower: rand.Intn(2) + 1,
	}
}

// for interface character
func (g *Goblin) TakeDamage(attackPower int) {
	g.Health -= attackPower
}

// for interface fmt.Stringer
func (g Goblin) String() string {
	return fmt.Sprintf("Goblin: %vHP, %vAP, %vDP", g.Health, g.AttackPower, g.DefensePower)
}
