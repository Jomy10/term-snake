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
	// Linux
	__Clear = func() {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

type CoPair = util.Pair[snake.Coordinate, string]

func Render(game *snake.Game) {
	// Collect sprites
	allSprites := []CoPair{}
	for i, cellCo := range game.Player.Co {
		if i == 0 {
			allSprites = append(allSprites, CoPair{cellCo, string(game.Player.HeadCell)})
		} else {
			allSprites = append(allSprites, CoPair{cellCo, string(game.Player.Cell)})
		}
	}

	if game.Player.Co[0].IsFurtherFrom(game.Food.Co) != snake.Equal {
		allSprites = append(allSprites, CoPair{game.Food.Co, string(game.Food.Cell)})
	}

	// fmt.Println(game.Player.Co[0].IsFurtherFrom(game.Food.Co))

	// Sort sprites
	sortedSprites := sort(&allSprites)

	// make surte it draws until the corner (insert empty sprite)
	if (sortedSprites[len(sortedSprites)-1].Left != snake.Coordinate{game.Width, game.Height}) {
		sortedSprites = append(sortedSprites, CoPair{snake.Coordinate{game.Width, game.Height}, " "})
	}

	fmt.Println(sortedSprites)

	// Draw sprites
	currentPos := snake.Coordinate{0, 0}
	fmt.Println(strings.Repeat("=", game.Width))
	for _, sprite := range sortedSprites {
		isFurther := sprite.Left.IsFurtherFrom(currentPos)
		if isFurther == snake.True {
			if sprite.Left.Y == currentPos.Y {
				diffX := sprite.Left.X - currentPos.X
				fmt.Print(strings.Repeat(" ", diffX))
				currentPos.X += diffX + 1
				fmt.Print(sprite.Right)
			} else {
				xToNewLine := game.Width - currentPos.X
				fmt.Print(strings.Repeat(" ", xToNewLine))
				fmt.Print("|\n")
				currentPos.Y += 1
				currentPos.X = 0

				for {
					if sprite.Left.Y == currentPos.Y {
						diffX := sprite.Left.X - currentPos.X
						fmt.Print(strings.Repeat(" ", diffX))
						currentPos.X += diffX + 1
						fmt.Print(sprite.Right)
						break
					} else {
						fmt.Print(strings.Repeat(" ", game.Width))
						fmt.Print("|\n")
						currentPos.Y += 1
					}
				}
			}
		} else if isFurther == snake.Equal {
			fmt.Print(sprite.Right)
		} else {
			panic(fmt.Sprintf("current: %s, sprite: %s", currentPos, sprite.Left))
			currentPos = sprite.Left
		}
	}
	fmt.Println(strings.Repeat("=", game.Width))
	/*
		// Clear map
		// for j, row := range game.GameMap {
		// 	for i := range row {
		// 		game.GameMap[j][i] = " "
		// 	}
		// }

		// Draw to map
		// game.GameMap[game.Food.Co.Y][game.Food.Co.X] = string(game.Food.Cell)
		// for i, co := range game.Player.Co {
		// 	cellSprite := string(game.Player.Cell)
		// 	if i == 0 {
		// 		cellSprite = string(game.Player.HeadCell)
		// 	}
		// 	game.GameMap[co.Y][co.X] = cellSprite
		// }
	*/

}

func DrawMap(_map [][]string) {
	__Clear()

	fmt.Println("-" + strings.Repeat("-", len(_map[0])) + "-")
	for _, row := range _map {
		fmt.Print("|")
		for _, cell := range row {
			fmt.Print(cell)
		}
		fmt.Print("|\n")
	}
	fmt.Println("-" + strings.Repeat("-", len(_map[0])) + "-")
}

func sort(pairs *[]CoPair) []CoPair {
	var sorted []CoPair = nil
	for i := 0; i < len(*pairs); i++ {
		if sorted == nil {
			sorted = []CoPair{(*pairs)[i]}
			continue
		}

		inList := false
		for idx, sortedPair := range sorted {
			fmt.Println((*pairs)[i].Left.IsFurtherFrom(sortedPair.Left), (*pairs)[i], sortedPair)
			if (*pairs)[i].Left.IsFurtherFrom(sortedPair.Left) == snake.False {
				var before []CoPair
				var after []CoPair
				if len(sorted) != 1 {
					before = sorted[:idx]
					after = sorted[idx+1:]
				} else {
					before = []CoPair{}
					after = sorted
				}
				insertVal := (*pairs)[i]
				sorted = append(append(before, insertVal), after...)
				inList = true
				break
			}
		}
		if !inList {
			sorted = append(sorted, (*pairs)[i])
		}
	}
	return sorted
}
