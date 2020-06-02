package rectangles

// CountBase counts the number of quadrilaterals drawn from the input
func CountBase(in []string) int {
	edges := parseEdges(in)

	rectangles := make([]struct{ a, b, c, d struct{ x, y int } }, 0)
	for a := 0; a < len(edges); a++ {
		for b := a + 1; b < len(edges); b++ {
			for c := b + 1; c < len(edges); c++ {
				for d := c + 1; d < len(edges); d++ {
					if !isHorizontalRect(edges[a], edges[b], edges[c], edges[d]) {
						continue
					}
					rectangles = append(rectangles, struct{ a, b, c, d struct{ x, y int } }{a: edges[a], b: edges[b], c: edges[c], d: edges[d]})
				}
			}
		}
	}
	var count int
	x := drawHorizontalLine(in)
	y := drawVerticalLine(in)
	for _, r := range rectangles {
		switch {
		case !checkLineFilling(x(r.a, r.b), '+', '-'):
		case !checkLineFilling(x(r.c, r.d), '+', '-'):
		case !checkLineFilling(y(r.a, r.c), '+', '|'):
		case !checkLineFilling(y(r.b, r.d), '+', '|'):
			continue
		default:
			count++
		}
	}
	return count
}

// CountEdgeAndSideTogether counts the number of quadrilaterals drawn from the input
func CountEdgeAndSideTogether(in []string) int {
	edges := parseEdges(in)

	var count int
	x := drawHorizontalLine(in)
	y := drawVerticalLine(in)
	for a := 0; a < len(edges); a++ {
		for b := a + 1; b < len(edges); b++ {
			if !sameX(edges[a], edges[b]) {
				continue
			}
			if !checkLineFilling(x(edges[a], edges[b]), '+', '-') {
				continue
			}
			for c := b + 1; c < len(edges); c++ {
				if !sameY(edges[a], edges[c]) {
					continue
				}
				if !checkLineFilling(y(edges[a], edges[c]), '+', '|') {
					continue
				}
				for d := c + 1; d < len(edges); d++ {
					switch {
					case !sameX(edges[c], edges[d]):
					case !checkLineFilling(x(edges[c], edges[d]), '+', '-'):
					case !sameY(edges[b], edges[d]):
					case !checkLineFilling(y(edges[b], edges[d]), '+', '|'):
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

// CountEdgesFirst counts the number of quadrilaterals drawn from the input
func CountEdgesFirst(in []string) int {
	edges := parseEdges(in)

	var count int
	x := drawHorizontalLine(in)
	y := drawVerticalLine(in)
	for a := 0; a < len(edges); a++ {
		for b := a + 1; b < len(edges); b++ {
			if !sameX(edges[a], edges[b]) {
				continue
			}
			for c := b + 1; c < len(edges); c++ {
				if !sameY(edges[a], edges[c]) {
					continue
				}
				for d := c + 1; d < len(edges); d++ {
					switch {
					case !sameX(edges[c], edges[d]):
					case !sameY(edges[b], edges[d]):
					case !checkLineFilling(x(edges[a], edges[b]), '+', '-'):
					case !checkLineFilling(x(edges[c], edges[d]), '+', '-'):
					case !checkLineFilling(y(edges[a], edges[c]), '+', '|'):
					case !checkLineFilling(y(edges[b], edges[d]), '+', '|'):
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

// CountSidesConc counts the number of quadrilaterals drawn from the input
func CountSidesConc(in []string) int {
	edges := parseEdges(in)

	var count int
	x := drawHorizontalLine(in)
	y := drawVerticalLine(in)
	for a := 0; a < len(edges); a++ {
		for b := a + 1; b < len(edges); b++ {
			if !sameX(edges[a], edges[b]) {
				continue
			}
			for c := b + 1; c < len(edges); c++ {
				if !sameY(edges[a], edges[c]) {
					continue
				}
			nextVert:
				for d := c + 1; d < len(edges); d++ {
					if !sameY(edges[b], edges[d]) || !sameX(edges[c], edges[d]) {
						continue
					}

					checkChan := make(chan bool)
					eA, eB, eC, eD := edges[a], edges[b], edges[c], edges[d]
					go func() { checkChan <- !checkLineFilling(x(eA, eB)) }()
					go func() { checkChan <- !checkLineFilling(x(eC, eD)) }()
					go func() { checkChan <- !checkLineFilling(y(eA, eC)) }()
					go func() { checkChan <- !checkLineFilling(y(eB, eD)) }()

					for i := 0; i < 4; i++ {
						if <-checkChan {
							continue nextVert
						}
					}
					count++
				}
			}
		}
	}
	return count
}

func drawHorizontalLine(in []string) func(a, b struct{ x, y int }) []byte {
	return func(a, b struct{ x, y int }) []byte {
		return []byte(in[a.x][a.y : b.y+1])
	}
}

func drawVerticalLine(in []string) func(a, b struct{ x, y int }) []byte {
	return func(a, b struct{ x, y int }) []byte {
		line := make([]byte, b.x-a.x+1)
		for i := range line {
			line[i] = in[i+a.x][a.y]
		}
		return line
	}
}

func checkLineFilling(line []byte, requiredChars ...byte) bool {
nextLine:
	for _, char := range line {
		for _, reqChar := range requiredChars {
			if char == reqChar {
				continue nextLine
			}
		}
		return false
	}
	return true
}

func isHorizontalRect(a, b, c, d struct{ x, y int }) bool {
	return sameX(a, b) && sameX(c, d) && sameY(a, c) && sameY(b, d)
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
