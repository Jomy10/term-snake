package snake

import (
	"math/rand"
	"time"
)

type Game struct {
	Player
	Food
	GameMap [][]string
	Size
}

type Coordinate struct {
	X, Y int
}

type BoolOrEqual int

const (
	False BoolOrEqual = iota
	True  BoolOrEqual = iota
	Equal BoolOrEqual = iota
)

func (co1 Coordinate) IsFurtherFrom(co2 Coordinate) BoolOrEqual {
	if co1.Y > co2.Y {
		return True
	} else if co1.Y == co2.Y && co1.X == co2.X {
		return Equal
	} else if co1.X > co2.X && co1.Y == co2.Y {
		return True
	} else {
		return False
	}
}

type Size struct {
	Width, Height int
}

func init() {
	rand.Seed(time.Now().Unix())
}

func RandCo(maxX, maxY int) Coordinate {
	return Coordinate{
		X: rand.Intn(maxX),
		Y: rand.Intn(maxY),
	}
}

// The cell sprite
type Cell string
