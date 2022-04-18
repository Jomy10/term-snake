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
	gameMap := make([][]string, height)
	for i := range gameMap {
		gameMap[i] = make([]string, width)
	}
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			gameMap[y][x] = " "
		}
	}
	return snake.Game{
		snake.Player{
			Co:       []snake.Coordinate{snake.Coordinate{width / 2, height / 2}, snake.Coordinate{width/2 + 1, height / 2}},
			Cell:     "\x1b[42m \x1b[0m",
			HeadCell: "\x1b[42m \x1b[0m",
		},
		snake.Food{
			snake.RandCo(width-1, height-1),
			"\x1b[41m \x1b[0m",
		},
		gameMap,
		snake.Size{width, height},
	}
}

var score uint64 = 0

func run(game *snake.Game, dirCh chan snake.Direction) {
	renderer.Render(game)
	renderer.DrawMap(game.GameMap, drawScore, 2)
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
			score += 1
		}
		if game.Player.WillCollide(game, moveDir) {
			fmt.Println("GAME OVER")
			break
		}
		game.Player.Move(moveDir, eaten)

		renderer.Render(game)
		renderer.DrawMap(game.GameMap, drawScore, 2)
		// TODO: use time elapsed instead of just sleep
		time.Sleep(500 * time.Millisecond)
	}
}

func drawScore() {
	fmt.Println(fmt.Sprintf(
		"\x1b[47;30m  Score: %d  \x1b[0m",
		// "        "+strings.Repeat(" ", CountDigits(score)),
		score,
	))
}

func main() {
	// Get terminal size
	if !term.IsTerminal(0) {
		fmt.Println("Snake cannot run outside of a terminal")
		os.Exit(1)
	}
	width, height := terminalSize()

	width = width/2 - 1
	height = height - 1

	game := initialize(width, height)
	channel := make(chan snake.Direction, 1)
	go snake.ListenDirection(channel)
	run(&game, channel)
}

func terminalSize() (int, int) {
	width, height, err := term.GetSize(0)
	if err != nil {
		width = 25
		height = 10
	} else {
		width -= 2
		height -= 3
	}
	return width, height
}
