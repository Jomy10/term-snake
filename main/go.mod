module jonaseveraert.be/snakeGame

go 1.18

replace jonaseveraert.be/renderer => ../renderer

replace jonaseveraert.be/snake => ../logic

require (
	golang.org/x/term v0.0.0-20220411215600-e5f449aeb171
	jonaseveraert.be/renderer v0.0.0-00010101000000-000000000000
	jonaseveraert.be/snake v0.0.0-00010101000000-000000000000
)

require (
	github.com/eiannone/keyboard v0.0.0-20200508000154-caf4b762e807 // indirect
	golang.org/x/sys v0.0.0-20220412211240-33da011f77ad // indirect
	jonaseveraert.be/util v0.0.0-00010101000000-000000000000 // indirect
)

replace jonaseveraert.be/util => ../util
