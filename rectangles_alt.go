package rectangles

// CountBase counts the number of quadrilaterals drawn from the input
func CountBase(in []string) int {
	vertices := parseVertices(in)

	rectangles := make([]struct{ a, b, c, d struct{ x, y int } }, 0)
	for a := 0; a < len(vertices); a++ {
		for b := a + 1; b < len(vertices); b++ {
			for c := b + 1; c < len(vertices); c++ {
				for d := c + 1; d < len(vertices); d++ {
					if !isHorizontalRect(vertices[a], vertices[b], vertices[c], vertices[d]) {
						continue
					}
					rectangles = append(rectangles, struct{ a, b, c, d struct{ x, y int } }{a: vertices[a], b: vertices[b], c: vertices[c], d: vertices[d]})
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
	vertices := parseVertices(in)

	var count int
	x := drawHorizontalLine(in)
	y := drawVerticalLine(in)
	for a := 0; a < len(vertices); a++ {
		for b := a + 1; b < len(vertices); b++ {
			if !sameX(vertices[a], vertices[b]) {
				continue
			}
			if !checkLineFilling(x(vertices[a], vertices[b]), '+', '-') {
				continue
			}
			for c := b + 1; c < len(vertices); c++ {
				if !sameY(vertices[a], vertices[c]) {
					continue
				}
				if !checkLineFilling(y(vertices[a], vertices[c]), '+', '|') {
					continue
				}
				for d := c + 1; d < len(vertices); d++ {
					switch {
					case !sameX(vertices[c], vertices[d]):
					case !checkLineFilling(x(vertices[c], vertices[d]), '+', '-'):
					case !sameY(vertices[b], vertices[d]):
					case !checkLineFilling(y(vertices[b], vertices[d]), '+', '|'):
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
	vertices := parseVertices(in)

	var count int
	x := drawHorizontalLine(in)
	y := drawVerticalLine(in)
	for a := 0; a < len(vertices); a++ {
		for b := a + 1; b < len(vertices); b++ {
			if !sameX(vertices[a], vertices[b]) {
				continue
			}
			for c := b + 1; c < len(vertices); c++ {
				if !sameY(vertices[a], vertices[c]) {
					continue
				}
				for d := c + 1; d < len(vertices); d++ {
					switch {
					case !sameX(vertices[c], vertices[d]):
					case !sameY(vertices[b], vertices[d]):
					case !checkLineFilling(x(vertices[a], vertices[b]), '+', '-'):
					case !checkLineFilling(x(vertices[c], vertices[d]), '+', '-'):
					case !checkLineFilling(y(vertices[a], vertices[c]), '+', '|'):
					case !checkLineFilling(y(vertices[b], vertices[d]), '+', '|'):
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
	vertices := parseVertices(in)

	var count int
	x := drawHorizontalLine(in)
	y := drawVerticalLine(in)
	for a := 0; a < len(vertices); a++ {
		for b := a + 1; b < len(vertices); b++ {
			if !sameX(vertices[a], vertices[b]) {
				continue
			}
			for c := b + 1; c < len(vertices); c++ {
				if !sameY(vertices[a], vertices[c]) {
					continue
				}
			nextVert:
				for d := c + 1; d < len(vertices); d++ {
					if !sameY(vertices[b], vertices[d]) || !sameX(vertices[c], vertices[d]) {
						continue
					}

					checkChan := make(chan bool)
					eA, eB, eC, eD := vertices[a], vertices[b], vertices[c], vertices[d]
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
