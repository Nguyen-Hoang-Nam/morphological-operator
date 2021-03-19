package components

func PointInTable(table []Point, point Point) bool {
	i := 0
	for i < len(table) {
		if Compare(table[i], point) {
			return true
		}
		i++
	}

	return false
}
