package main

import (
	"fmt"
	"math/rand"
	"time"

	keyboard "github.com/eiannone/keyboard"
)

func main() {

	var grid = [20][20]int{}
	fmt.Println(grid)

	p1 := Player{name: "Helder Salvador", p: Position{positionX: 0, positionY: 0}, LifePoints: 10, dmgMin: 5, dmgMax: 10}
	p2 := Player{name: "Ayoub Fakir", p: Position{positionX: 1, positionY: 1}, LifePoints: 10, dmgMin: 5, dmgMax: 10}

	attack(&p2, &p1)
	attack(&p2, &p1)
	attack(&p1, &p2)

	char, _, err := keyboard.GetSingleKey()
	if err != nil {
		panic(err)
	}
	fmt.Printf("You pressed: %q\r\n", char)

	for true {
		switch char {
		case 'z':
			fmt.Println("z")
			break
		case 's':
			fmt.Println("s")
			break
		case 'q':
			fmt.Println("s")
			break
		case 'd':
			fmt.Println("s")
			break
		default:
			break
		}
	}
}

type Position struct {
	positionX int
	positionY int
}

type Player struct {
	p          Position
	name       string
	LifePoints int
	dmgMin     int
	dmgMax     int
}

func newPlayer(_name string, _position Position) *Player {
	p := Player{name: _name, p: _position, dmgMin: 5, dmgMax: 10}
	return &p
}

func newPosition(x int, y int) *Position {
	s := Position{positionX: x, positionY: y}
	return &s
}

func attack(from *Player, to *Player) *Player {
	rand.Seed(time.Now().UnixNano())
	if from.LifePoints > 0 {
		if to.LifePoints > 0 {
			fmt.Printf("WOUAH ! %s a frappé %s !\n", from.name, to.name)

			jet := rand.Intn((from.dmgMax - from.dmgMin + 1)) + from.dmgMin
			lastingLifePoints := to.LifePoints - jet
			to.LifePoints = lastingLifePoints
			fmt.Printf("Dégats : %d \n", jet)

			return to
		} else {
			fmt.Printf("MAIS... %s est mort !\n", to.name)
			return to
		}

	} else {
		fmt.Printf("MAIS... %s est mort, il ne peut plus attaqué !\n", from.name)
		return to
	}
}
