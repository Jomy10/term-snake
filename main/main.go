package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"golang.org/x/term"
	"jonaseveraert.be/renderer"
	"jonaseveraert.be/snake"
)

func initialize(width, height int) snake.Game {
	// gameMap := make([][]string, height)
	// for i := range gameMap {
	// 	gameMap[i] = make([]string, width)
	// }
	// for x := 0; x < width; x++ {
	// 	for y := 0; y < height; y++ {
	// 		fmt.Println(x, y)
	// 		//gameMap[y][x] = " "
	// 	}
	// }
	return snake.Game{
		snake.Player{
			Co:       []snake.Coordinate{snake.Coordinate{width / 2, height / 2}, snake.Coordinate{width/2 + 1, height / 2}},
			Cell:     "=",
			HeadCell: "^",
		},
		snake.Food{
			snake.RandCo(width-1, height-1),
			"•",
		},
		// gameMap,
		snake.Size{width, height},
	}
}

func run(game *snake.Game, dirCh chan snake.Direction) {
	renderer.Render(game)
	// renderer.DrawMap(game.GameMap)
	game.Player.CurrentDir = snake.Direction(rand.Intn(3))
	for {
		moveDir := game.Player.CurrentDir
		select {
		case dir := <-dirCh:
			moveDir = dir
		default:
		}

		eaten := false
		if game.Player.Collides(&game.Food) {
			// random new location
			game.Food = snake.RandomFoodLocation(game, game.Food.Cell)
			eaten = true
		}
		if game.Player.WillCollide(game, moveDir) {
			fmt.Println("GAME OVER")
			break
		}
		game.Player.Move(moveDir, eaten)

		renderer.Render(game)
		// renderer.DrawMap(game.GameMap)
		time.Sleep(1 * time.Second)
	}
}

func main() {
	// Get terminal size
	if !term.IsTerminal(0) {
		fmt.Println("Snake cannot run outside of a terminal")
		os.Exit(1)
	}
	// width, height := terminalSize()

	//game := initialize(width-2, height-2)
	game := initialize(5, 5)
	channel := make(chan snake.Direction, 1)
	go snake.ListenDirection(channel)
	run(&game, channel)
}

func terminalSize() (int, int) {
	width, height, err := term.GetSize(0)
	if err != nil {
		width = 25
		height = 10
	}
	return width, height
}
