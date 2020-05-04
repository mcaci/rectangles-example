package rectangles

// Count counts the number of quadrilaterals drawn from the input
func Count(in []string) int {
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
	xOk := xLinePresent(in)
	yOk := yLinePresent(in)
	for _, r := range rectangles {
		switch {
		case !xOk(r.a, r.b):
		case !xOk(r.c, r.d):
		case !yOk(r.a, r.c):
		case !yOk(r.b, r.d):
			continue
		default:
			count++
		}
	}
	return count
}

// CountAllEdgesTaken counts the number of rectangles drawn from the input
func CountAllEdgesTaken(in []string) int {
	edges := parseEdges(in)

	var count int
	xOk := xLinePresent(in)
	yOk := yLinePresent(in)
	for a := 0; a < len(edges); a++ {
		for b := a + 1; b < len(edges); b++ {
			for c := b + 1; c < len(edges); c++ {
				for d := c + 1; d < len(edges); d++ {
					switch {
					case !sameX(edges[a], edges[b]):
					case !sameY(edges[a], edges[c]):
					case !sameX(edges[c], edges[d]):
					case !sameY(edges[b], edges[d]):
					case !xOk(edges[a], edges[b]):
					case !xOk(edges[c], edges[d]):
					case !yOk(edges[a], edges[c]):
					case !yOk(edges[b], edges[d]):
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

// CountEdgeAndSideTogether counts the number of quadrilaterals drawn from the input
func CountEdgeAndSideTogether(in []string) int {
	edges := parseEdges(in)

	var count int
	xOk := xLinePresent(in)
	yOk := yLinePresent(in)
	for a := 0; a < len(edges); a++ {
		for b := a + 1; b < len(edges); b++ {
			if !sameX(edges[a], edges[b]) {
				continue
			}
			if !xOk(edges[a], edges[b]) {
				continue
			}
			for c := b + 1; c < len(edges); c++ {
				if !sameY(edges[a], edges[c]) {
					continue
				}
				if !yOk(edges[a], edges[c]) {
					continue
				}
				for d := c + 1; d < len(edges); d++ {
					switch {
					case !sameX(edges[c], edges[d]):
					case !xOk(edges[c], edges[d]):
					case !sameY(edges[b], edges[d]):
					case !yOk(edges[b], edges[d]):
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
	xOk := xLinePresent(in)
	yOk := yLinePresent(in)
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
					case !xOk(edges[a], edges[b]):
					case !xOk(edges[c], edges[d]):
					case !yOk(edges[a], edges[c]):
					case !yOk(edges[b], edges[d]):
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
