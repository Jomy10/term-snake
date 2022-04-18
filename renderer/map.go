package renderer

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"jonaseveraert.be/snake"
	"jonaseveraert.be/util"
)

var __Clear func()

func init() {
	// Unix
	__Clear = func() {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

type CoPair = util.Pair[snake.Coordinate, string]

func Render(game *snake.Game) {
	// Clear map
	for j, row := range game.GameMap {
		for i := range row {
			game.GameMap[j][i] = " "
		}
	}

	// Draw to map
	game.GameMap[game.Food.Co.Y][game.Food.Co.X] = string(game.Food.Cell)
	for i, co := range game.Player.Co {
		cellSprite := string(game.Player.Cell)
		if i == 0 {
			cellSprite = string(game.Player.HeadCell)
		}
		game.GameMap[co.Y][co.X] = cellSprite
	}

}

func DrawMap(_map [][]string, drawBefore func(), cellWidth int) {
	__Clear()

	drawBefore()

	fmt.Println("\x1b[47m  " + strings.Repeat(" ", (len(_map[0]))*cellWidth) + "  \x1b[0m")
	for _, row := range _map {
		fmt.Print("\x1b[47m  \x1b[0m")
		for _, cell := range row {
			for i := 0; i < cellWidth; i++ {
				fmt.Print(cell)
			}
		}
		fmt.Print("\x1b[47m  \x1b[0m\n")
	}
	fmt.Println("\x1b[47m  " + strings.Repeat(" ", (len(_map[0]))*cellWidth) + "  \x1b[0m")
}
