package main

import (
	"fmt"
	"math"
	"math/rand"
	"strings"
	"time"
)

type EnumCardValue int
type EnumCardSuit string

type Hand []Card
type Deck []Card

type Card struct {
	Value EnumCardValue
	Suit  EnumCardSuit
}

type CardValue struct {
	Score int
	Name  string
}

const (
	cv_ace EnumCardValue = iota + 1
	cv_two
	cv_three
	cv_four
	cv_five
	cv_six
	cv_seven
	cv_eight
	cv_nine
	cv_ten
	cv_jack
	cv_queen
	cv_king
)

const (
	cs_spade   EnumCardSuit = "♠"
	cs_club                 = "♥"
	cs_heart                = "♣"
	cs_diamond              = "♦"
)

var cvs = map[EnumCardValue]CardValue{
	cv_ace:   {1, "A"},
	cv_two:   {2, "2"},
	cv_three: {3, "3"},
	cv_four:  {4, "4"},
	cv_five:  {5, "5"},
	cv_six:   {6, "6"},
	cv_seven: {7, "7"},
	cv_eight: {8, "8"},
	cv_nine:  {9, "9"},
	cv_ten:   {10, "10"},
	cv_jack:  {10, "J"},
	cv_queen: {10, "Q"},
	cv_king:  {10, "K"},
}

var suits = []EnumCardSuit{cs_spade, cs_club, cs_heart, cs_diamond}

var values = []EnumCardValue{cv_ace, cv_two, cv_three, cv_four, cv_five, cv_six, cv_seven, cv_eight, cv_nine, cv_ten, cv_jack, cv_queen, cv_king}

// for fmt.Stringer
func (c Card) String() string {
	return fmt.Sprintf("%v%v", cvs[c.Value].Name, c.Suit)
}

func GenerateDeck() Deck {
	var deck Deck
	for s := 0; s < len(suits); s++ {
		for v := 0; v < len(values); v++ {
			deck = append(deck, Card{Value: values[v], Suit: suits[s]})
		}
	}
	return deck
}

func (d Deck) Shuffle() {
	if len(d) == 52 {
		rand.Shuffle(len(d), func(i, j int) {
			d[i], d[j] = d[j], d[i]
		})
	}
}

// draw and remove a card from the deck taking from the front of the deck. Returns the card drawn
func (d *Deck) Draw1() Card {
	card := (*d)[0]
	*d = append((*d)[1:len((*d))], (*d)[1:]...)[:len((*d))-1]
	return card
}

// returns the score of the hand
func (h Hand) Score() int {
	score_ace1 := 0
	score_ace11 := 0
	// special case: one of the cards is an ace, and one of the cards is 10+
	// return 21
	if h.containsCard(cvs[cv_ace]) && any(h, cvs[cv_ten], cvs[cv_jack], cvs[cv_queen], cvs[cv_king]) {
		return 21
	}

	for _, card := range h {
		// ace is both 1 and 11, pick the score that will give the minimum score
		if card.Value == cv_ace {
			score_ace11 += 11
		} else {
			score_ace11 += cvs[card.Value].Score
		}
		score_ace1 += cvs[card.Value].Score
	}
	return int(math.Min(float64(score_ace1), float64(score_ace11)))
}

func (h Hand) IsBust() bool {
	return h.Score() > 21
}

func (h Hand) IsBlackJack() bool {
	return h.Score() == 21
}

func (h Hand) containsCard(c CardValue) bool {
	for _, card := range h {
		if cvs[card.Value] == c {
			return true
		}
	}
	return false
}

func any(h Hand, elements ...CardValue) bool {
	for _, e := range elements {
		if h.containsCard(e) {
			return true
		}
	}
	return false
}

type Play interface {
	Hit(*Deck)
	Stand()
}

type Player struct {
	Hand     *Hand
	IsDealer bool
}

// for Play
func (p *Player) Hit(d *Deck) {
	card := d.Draw1()
	*(p.Hand) = append(*(p.Hand), card)
}

// for Play
func (p Player) Stand() {
	fmt.Println("standing")
}

// for fmt.Stringer
func (p Player) String() string {
	return fmt.Sprintf("%v", p.Hand)
}

func DoBlackjackRound(players ...Player) {
	deck := GenerateDeck()
	deck.Shuffle()
	fmt.Println(deck)

	// init the round:
	// 		all players get dealt to twice in a circle
	for i := 0; i < 2*len(players); i++ {
		players[i%len(players)].Hit(&deck)
	}
	for _, p := range players {
		if p.IsDealer {
			fmt.Printf("dealer: %v = %v\n", p, p.Hand.Score())
		} else {
			fmt.Printf("player: %v = %v\n", p, p.Hand.Score())
		}
	}
	fmt.Println(deck)

	for { // until someone busts or blackjacks, go through each player for hit and stand
		var gameIsDone bool = false
		for _, p := range players {
			if gameIsDone {
				fmt.Printf("other player has hand %v, score %v\n", p.Hand, p.Hand.Score())
				break
			}
			if !p.IsDealer {
				// ask for input
				fmt.Print("hit or stand? (hit/stand) ")
				var response string
				for response != "hit" && response != "stand" {
					fmt.Scan(&response)
					response = strings.ToLower(response)
					switch response {
					case "hit":
						p.Hit(&deck)
						break
					case "stand":
						p.Stand()
						break
					default:
						fmt.Print("unknown action, try again (hit/stand): ")
					}
				}
			} else {
				// dealer must hit on <= 17
				if p.Hand.Score() <= 17 {
					p.Hit(&deck)
				} else {
					p.Stand()
				}
			}
			if p.IsDealer {
				fmt.Printf("post-turn dealer: %v = %v\n", p, p.Hand.Score())
			} else {
				fmt.Printf("post-turn player: %v = %v\n", p, p.Hand.Score())
			}
			// check for busts and blackjack
			var name string = "Player"
			if p.IsDealer {
				name = "Dealer"
			}
			if p.Hand.IsBust() {
				fmt.Printf("%s busts out with hand %v, score %v\n", name, p.Hand, p.Hand.Score())
				gameIsDone = true
				break
			}
			if p.Hand.IsBlackJack() {
				fmt.Printf("%s has BlackJack with hand %v, score %v\n", name, p.Hand, p.Hand.Score())
				gameIsDone = true
				break
			}
		}
		if gameIsDone {
			break
		}
	}
}

func main() {
	rand.Seed(time.Now().Unix())

	dealer := Player{Hand: new(Hand), IsDealer: true}
	player := Player{Hand: new(Hand), IsDealer: false}

	DoBlackjackRound(dealer, player)

	fmt.Println(dealer, player)
}
