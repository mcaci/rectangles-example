package rectangles

// CountAll counts the number of rectangles drawn from the input
func CountAll(in []string) int {
	edges := parseEdges(in)

	var count int
	hLine := lineFilling(drawnLine{points: drawHorizontals(in), reqChars: []byte{'+', '-'}})
	vLine := lineFilling(drawnLine{points: drawVerticals(in), reqChars: []byte{'+', '|'}})
	for a := 0; a < len(edges); a++ {
		for b := a + 1; b < len(edges); b++ {
			for c := b + 1; c < len(edges); c++ {
				for d := c + 1; d < len(edges); d++ {
					switch {
					case !sameX(edges[a], edges[b]):
					case !sameY(edges[a], edges[c]):
					case !sameX(edges[c], edges[d]):
					case !sameY(edges[b], edges[d]):
					case !hLine(edges[a], edges[b]):
					case !vLine(edges[a], edges[c]):
					case !hLine(edges[c], edges[d]):
					case !vLine(edges[b], edges[d]):
						continue
					default:
						count++
					}
				}
			}
		}
	}
	return count
}

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

func drawHorizontals(in []string) func(a, b struct{ x, y int }) []byte {
	return func(a, b struct{ x, y int }) []byte {
		return []byte(in[a.x][a.y : b.y+1])
	}
}

func drawVerticals(in []string) func(a, b struct{ x, y int }) []byte {
	return func(a, b struct{ x, y int }) []byte {
		line := make([]byte, b.x-a.x+1)
		for i := range line {
			line[i] = in[i+a.x][a.y]
		}
		return line
	}
}

type drawnLine struct {
	points   func(a, b struct{ x, y int }) []byte
	reqChars []byte
}

func lineFilling(line drawnLine) func(a, b struct{ x, y int }) bool {
	return func(a, b struct{ x, y int }) bool {
	nextLine:
		for _, char := range line.points(a, b) {
			for _, reqChar := range line.reqChars {
				if char == reqChar {
					continue nextLine
				}
			}
			return false
		}
		return true
	}
}
