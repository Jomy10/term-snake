package snake

type Player struct {
	Co         []Coordinate
	Cell       Cell
	HeadCell   Cell
	CurrentDir Direction
}

func (player *Player) Move(dir Direction, eaten bool) {
	if !AreOpposite(player.CurrentDir, dir) {
		player.CurrentDir = dir
	}
	firstCell := player.Co[0]
	var newCell Coordinate
	switch player.CurrentDir {
	case Up:
		newCell = Coordinate{firstCell.X, firstCell.Y - 1}
	case Down:
		newCell = Coordinate{firstCell.X, firstCell.Y + 1}
	case Left:
		newCell = Coordinate{firstCell.X - 1, firstCell.Y}
	case Right:
		newCell = Coordinate{firstCell.X + 1, firstCell.Y}
	}

	if eaten {
		player.Co = append([]Coordinate{newCell}, player.Co...) // prepend
	} else {
		player.Co = append([]Coordinate{newCell}, player.Co[:len(player.Co)-1]...) // prepend & remove last element
	}
}

func (player *Player) Collides(food *Food) bool {
	return player.Co[0] == food.Co
}

func (player *Player) WillCollide(game *Game, dir Direction) bool {
	x, y := player.Co[0].X, player.Co[0].Y
	switch dir {
	case Up:
		y -= 1
	case Down:
		y += 1
	case Left:
		x -= 1
	case Right:
		x += 1
	}

	return x < 0 || x >= game.Width || y < 0 || y >= game.Height
}
