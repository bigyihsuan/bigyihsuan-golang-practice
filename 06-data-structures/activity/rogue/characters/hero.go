package characters

import (
	"fmt"
	"math/rand"
)

type Hero struct {
	MaxHealth     int
	CurrentHealth int
	AttackPower   int
	DefensePower  int
	Potions       [5]int
	Gold          int
	Level         int
	GoblinsKilled int
}

func NewHero() Hero {
	health := rand.Intn(11) + 20
	return Hero{
		MaxHealth:     health,
		CurrentHealth: health,
		AttackPower:   rand.Intn(3) + 1,
		DefensePower:  rand.Intn(5) + 1,
		Potions:       [5]int{2, 2, 2, 2, 2},
		Gold:          0,
		Level:         1,
		GoblinsKilled: 0,
	}
}

func (h *Hero) LevelUp() {
	h.Level += 1
	fmt.Printf("Leveled up to %v\n", h.Level)
}

func (h Hero) PotionCount() int {
	numPotions := 0
	for i := range h.Potions {
		if h.Potions[i] != 0 {
			numPotions += 1
		}
	}
	return numPotions
}

func (h *Hero) UsePotion() bool {
	if h.PotionCount() <= 0 {
		fmt.Println("Out of potions!")
		return false
	}
	for i := range h.Potions {
		if h.Potions[i] == 2 { // use the first potion
			h.Potions[i] = 0
			h.CurrentHealth += 2
			fmt.Printf("Used potion, %v potions left\n", h.PotionCount())
			break
		}
	}
	if h.CurrentHealth > h.MaxHealth {
		h.CurrentHealth = h.MaxHealth
	}
	return true
}

func (h *Hero) FillPotions(numPotions int) {
	numLeft := numPotions
	for i := range h.Potions {
		if h.Potions[i] <= 0 {
			h.Potions[i] = 2
			numLeft -= 1
		}
		if numLeft <= 0 {
			return
		}
	}
}

// randomly generate a goblin, about 1-in-5 chance?
func (h *Hero) Step() *Goblin {
	if rand.Intn(5) == 0 {
		goblin := NewGoblin()
		return &goblin
	}
	return nil
}

// for interface character
func (h *Hero) TakeDamage(attackPower int) {
	h.CurrentHealth -= attackPower
}

func (h *Hero) RewardFight() {
	h.Gold += 2
	h.GoblinsKilled += 1
}

// hero fights some given goblin. returns whether the hero won.
// hero will auto-use potions when health is low
func (h *Hero) Fight(g *Goblin) bool {
	for h.CurrentHealth > 0 && g.Health > 0 {
		g.TakeDamage(h.AttackPower) // hero attack
		h.TakeDamage(g.AttackPower) // goblin attack
		for h.CurrentHealth-2 <= 0 {
			couldUsePotion := h.UsePotion()
			if !couldUsePotion {
				break
			}
		}
	}
	return h.CurrentHealth > 0 && g.Health <= 0
}

// for interface Stringer
func (h Hero) String() string {
	return fmt.Sprintf("Hero: Lvl.%v, %v/%vHP, %vAP, %vDP, %vG, %v potions", h.Level, h.CurrentHealth, h.MaxHealth, h.AttackPower, h.DefensePower, h.Gold, h.PotionCount())
}
