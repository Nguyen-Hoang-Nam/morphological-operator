package components

import (
	"fmt"
)

func drawHorizontalLine(horizontal rune, left rune, mid rune, right rune, xMax int) {
	fmt.Printf("%c%c", left, horizontal)

	x := 0
	for x < xMax {
		fmt.Printf("%c%c%c%c", horizontal, horizontal, mid, horizontal)
		x++
	}

	fmt.Printf("%c%c%c\n", horizontal, horizontal, right)
}

func DrawTable(table [][]rune, xMax int, yMax int) {
	vertical := '│'
	horizontal := '─'
	topMid := '┬'
	topLeft := '┌'
	topRight := '┐'
	bottomMid := '┴'
	bottomLeft := '└'
	bottomRight := '┘'
	midMid := '┼'
	midLeft := '├'
	midRight := '┤'

	// Top of table
	drawHorizontalLine(horizontal, topLeft, topMid, topRight, xMax)
	
	// Middle of table
	y := yMax
	for y > 0 {
		fmt.Printf("%c ", vertical)

		x := 0
		for x < xMax {
			fmt.Printf("%c %c ", table[y][x], vertical)
			x++
		}

		fmt.Printf("%c %c\n", table[y][xMax], vertical)
		y--

		drawHorizontalLine(horizontal, midLeft, midMid, midRight, xMax)
	}

	fmt.Printf("%c ", vertical)

	x := 0
	for x < xMax {
		fmt.Printf("%c %c ", table[y][x], vertical)
		x++
	}

	fmt.Printf("%c %c\n", table[y][xMax], vertical)
	y--

	// Bottom of table
	drawHorizontalLine(horizontal, bottomLeft, bottomMid, bottomRight, xMax)
}
