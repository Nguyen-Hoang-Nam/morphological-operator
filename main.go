package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
	"regexp"

	"github.com/Nguyen-Hoang-Nam/morphological-operators/components"
)

func showTable(points []components.Point) {
	if len(points) == 0 {
		fmt.Println("Empty")
	} else {
		xMax := points[0].X
		yMax := points[0].Y

		i := 0
		for i < len(points) {
			if xMax < points[i].X {
				xMax = points[i].X
			}

			if yMax < points[i].Y {
				yMax = points[i].Y
			}
			i++
		}

		table := make([][]rune, 0)
		x := 0
		y := 0
		for y <= yMax {
			row := make([]rune, 0)
			x = 0
			for x <= xMax {
				row = append(row, ' ')
				x++
			}
			table = append(table, row)
			y++
		}

		i = 0
		for i < len(points) {
			table[points[i].Y][points[i].X] = '■'
			i++
		}

		if table[0][0] == ' ' {
			table[0][0] = '✖'
		} else {
			table[0][0] = '☒'
		}

		components.DrawTable(table, xMax, yMax)		
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	variables := make(map[string][]components.Point)

	for true {
		fmt.Print("> ")

		command, err := reader.ReadString('\n')
		if err != nil {
			panic(err)
		}

		command = strings.TrimSuffix(command, "\n")
		if command == "exit" {
			fmt.Println("Bye")
			break
		} else {
			tokens := strings.Split(command, " ")
			
			if len(tokens) == 1 {
				points := variables[tokens[0]]
				showTable(points)
			} else if len(tokens) == 3 {
				if tokens[1] == "=" {
					listPointPattern := regexp.MustCompile(`(\d+,\d+)`)
					listPoint := listPointPattern.FindAllString(tokens[2], -1)

					var points []components.Point
					i := 0
					for i < len(listPoint) {
						point := components.StringToPoint(listPoint[i])
						points = append(points, point)
						i++
					}

					variables[tokens[0]] = points
					fmt.Println("True")
				} else if tokens[1] == "dilation" {
					points := components.Dilation(variables[tokens[0]], variables[tokens[2]])
					showTable(points)
				} else if tokens[1] == "erosion" {
					points := components.Erosion(variables[tokens[0]], variables[tokens[2]])
					showTable(points)
				} else if tokens[1] == "open" {
					points := components.Open(variables[tokens[0]], variables[tokens[2]])
					showTable(points)
				} else if tokens[1] == "close" {
					points := components.Close(variables[tokens[0]], variables[tokens[2]])
					showTable(points)
				}
			} else {
				fmt.Println("Error")
			}
		}
	}
}
