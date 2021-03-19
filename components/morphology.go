package components

func Dilation(original []Point, kernel []Point) []Point {
	result := make([]Point, 0)

	o := 0
	for o < len(original) {
		k := 0
		for k < len(kernel) {
			temp := Plus(original[o], kernel[k])
			result = append(result, temp)
			k++
		}
		o++
	}

	return result
}

func Erosion(original []Point, kernel []Point) []Point {
	result := make([]Point, 0)

	o := 0
	for o < len(original) {
		temps := make([]Point, 0)
		valid := true

		k := 0
		for k < len(kernel) {
			temp := Plus(original[o], kernel[k])
			if !PointInTable(original, temp) {
				valid = false
				break
			}

			temps = append(temps, temp)
			k++
		}

		if valid {
			result = append(result, temps...)
		}
		o++
	}

	return result
}

func Open(original []Point, kernel []Point) []Point {
	return Dilation(Erosion(original, kernel), kernel)
}

func Close(original []Point, kernel []Point) []Point {
	return Erosion(Dilation(original, kernel), kernel)
}
