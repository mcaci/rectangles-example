package rectangles

func parseEdges(in []string) []struct{ x, y int } {
	edges := make([]struct{ x, y int }, 0)
	for i, line := range in {
		for j, c := range line {
			if c != '+' {
				continue
			}
			edges = append(edges, struct{ x, y int }{x: i, y: j})
		}
	}
	return edges
}

func sameX(a, b struct{ x, y int }) bool { return a.x == b.x }
func sameY(a, b struct{ x, y int }) bool { return a.y == b.y }
func isHorizontalRect(a, b, c, d struct{ x, y int }) bool {
	return sameX(a, b) && sameX(c, d) && sameY(a, c) && sameY(b, d)
}

func xLine(in []string) func(a, b struct{ x, y int }) []byte {
	return func(a, b struct{ x, y int }) []byte {
		return []byte(in[a.x][a.y : b.y+1])
	}
}

func yLine(in []string) func(a, b struct{ x, y int }) []byte {
	return func(a, b struct{ x, y int }) []byte {
		line := make([]byte, b.x-a.x+1)
		for i := range line {
			line[i] = in[i+a.x][a.y]
		}
		return line
	}
}

func lineFilled(line []byte, lineElems ...byte) bool {
nextLine:
	for _, l := range line {
		for _, lElem := range lineElems {
			if lElem == l {
				continue nextLine
			}
		}
		return false
	}
	return true
}

// Not useful for the purpose of the exercise as the sides are only horizontal or vertical
// isHorizontalRect is a much quicker check under this condition
func isRectangle(a, b, c, d struct{ x, y int }) bool {
	center := struct{ x, y float64 }{x: float64(a.x+b.x+c.x+d.x) / 4, y: float64(a.y+b.y+c.y+d.y) / 4}

	sqr := func(x float64) float64 { return x * x }
	sqrDist := func(a struct{ x, y int }, b struct{ x, y float64 }) float64 {
		return sqr(b.x-float64(a.x)) + sqr(b.y-float64(a.y))
	}

	dist := sqrDist(a, center)
	return dist == sqrDist(b, center) && dist == sqrDist(c, center) && dist == sqrDist(d, center)
}
