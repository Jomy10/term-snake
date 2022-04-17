package snake

type Direction int

const (
	Up    Direction = iota
	Down  Direction = iota
	Left  Direction = iota
	Right Direction = iota
	None  Direction = iota
)

func AreOpposite(dir1, dir2 Direction) bool {
	switch dir1 {
	case Left:
		if dir2 == Right {
			return true
		}
	case Right:
		if dir2 == Left {
			return true
		}
	case Up:
		if dir2 == Down {
			return true
		}
	case Down:
		if dir2 == Up {
			return true
		}
	}
	return false
}
