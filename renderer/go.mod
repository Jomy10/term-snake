module jonaseveraert.be/snakeGameRenderer

go 1.18

replace jonaseveraert.be/snake => ../logic

require (
	jonaseveraert.be/snake v0.0.0-00010101000000-000000000000
	jonaseveraert.be/util v0.0.0-00010101000000-000000000000
)

require (
	github.com/eiannone/keyboard v0.0.0-20200508000154-caf4b762e807 // indirect
	golang.org/x/sys v0.0.0-20220412211240-33da011f77ad // indirect
)

replace jonaseveraert.be/util => ../util
