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
func isPlainRect(a, b, c, d struct{ x, y int }) bool {
	return sameX(a, b) && sameX(c, d) && sameY(a, c) && sameY(b, d)
}

func isXLinePresent(in []string) func(a, b, c, d struct{ x, y int }) bool {
	f := xLinePresent(in)
	return func(a, b, c, d struct{ x, y int }) bool {
		return f(a, b) && f(c, d)
	}
}

func isYLinePresent(in []string) func(a, b, c, d struct{ x, y int }) bool {
	f := yLinePresent(in)
	return func(a, b, c, d struct{ x, y int }) bool {
		return f(a, c) && f(b, d)
	}
}

func xLinePresent(in []string) func(a, b struct{ x, y int }) bool {
	return func(a, b struct{ x, y int }) bool {
		return linePresent([]byte(in[a.x][a.y:b.y+1]), '-', '+')
	}
}

func yLinePresent(in []string) func(a, b struct{ x, y int }) bool {
	return func(a, b struct{ x, y int }) bool {
		side := make([]byte, 0)
		for i := a.x; i <= b.x; i++ {
			side = append(side, in[i][a.y])
		}
		return linePresent(side, '|', '+')
	}
}

func linePresent(line []byte, lineElems ...byte) bool {
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
