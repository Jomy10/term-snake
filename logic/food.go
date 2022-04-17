package snake

import "math/rand"

type Food struct {
	Co Coordinate
	Cell
}

func RandomFoodLocation(game *Game, cell Cell) Food {
	randX := rand.Intn(game.Width - 1)
	randY := rand.Intn(game.Height - 1)

	// Make sure food is not in player
	for {
		collides := checkCollision(&game.Player, randX, randY)
		if collides {
			randX = rand.Intn(game.Width - 1)
			randY = rand.Intn(game.Height - 1)
		} else {
			break
		}
	}

	return Food{Coordinate{randX, randY}, cell}
}

func checkCollision(player *Player, x, y int) bool {
	collidesX := false
	collidesY := false
	for _, snakeBit := range player.Co {
		if snakeBit.X == x {
			if collidesY {
				return true
			} else {
				collidesX = true
			}
		}
		if snakeBit.Y == y {
			if collidesX {
				return true
			} else {
				collidesY = true
			}
		}
	}
	return false
}
