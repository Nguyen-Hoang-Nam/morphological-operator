package components

import (
	"strings"
	"strconv"
)

type Point struct {
	X int
	Y int
}

func Plus(p1, p2 Point) Point {
	var r Point

	r.X = p1.X + p2.X
	r.Y = p1.Y + p2.Y

	return r
}

func Compare(p1, p2 Point) bool {
	return p1.X == p2.X && p1.Y == p2.Y
}

func StringToPoint(str string) Point {
	var r Point

	coor := strings.Split(str, ",")
	r.X, _ = strconv.Atoi(coor[0])
	r.Y, _ = strconv.Atoi(coor[1])

	return r
}
