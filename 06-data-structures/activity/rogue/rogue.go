package main

import (
	"fmt"
	"math"
	"math/rand"
	"rogue/characters"
	"strings"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	hero := characters.NewHero()
	fmt.Println(hero)

	var response string
	for {
		for i := 1; ; i++ {
			goblin := hero.Step()
			if goblin != nil {
				fmt.Printf("step %v has goblin\n", i)
				fmt.Println(goblin)
				heroWins := hero.Fight(goblin)
				if heroWins {
					hero.RewardFight()
					fmt.Printf("Hero win\n%v\n", hero)
				} else {
					fmt.Print("Goblin win!\n")
					fmt.Printf("Hero dies!\n%v\n", hero)
					break
				}
			} else {
				fmt.Println("nothing on step", i)
			}
			if i%10 == 0 {
				hero.LevelUp()
				fmt.Println(hero)
				fmt.Print("visit potion shop? (y/n) ")
				fmt.Scan(&response)
				response = strings.ToLower(response)
				if response == "y" {
					var numPotions = 0
					fmt.Println("How many potions to buy?")
					fmt.Printf("Bag contains %v potions\n", hero.PotionCount())
					fmt.Printf("You currently have %v gold; 1 potion = 4 gold\n", hero.Gold)
					bagCapacity := len(hero.Potions) - hero.PotionCount()
					maxBuyable := int(math.Min(float64(hero.Gold/4), float64(bagCapacity)))
					fmt.Printf("(Enter number 0-%v): ", maxBuyable)
					fmt.Scan(&numPotions)
					for numPotions > maxBuyable {
						fmt.Print("Can't hold that many. Enter a lower amount: ")
						fmt.Scan(&numPotions)
					}
					fmt.Printf("Bought %v potions\n", numPotions)
					hero.FillPotions(numPotions)
				}
			}
		}
		fmt.Print("Would you like to play again? (y/n)")
		fmt.Scan(&response)
		response = strings.ToLower(response)
		if response != "y" {
			break
		}
		fmt.Println("Making a new character...")
		lastHerosGold := hero.Gold
		hero := characters.NewHero()
		hero.Gold = lastHerosGold
	}

	fmt.Printf("Hero had %v levels, %v gold, and killed %v goblins\n", hero.Level, hero.Gold, hero.GoblinsKilled)
}
