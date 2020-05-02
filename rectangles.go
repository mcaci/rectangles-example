package rectangles

// Count counts the number of quadrilaterals drawn from the input
func Count(in []string) int {
	edges := parseEdges(in)

	rectangles := make([]struct{ a, b, c, d struct{ x, y int } }, 0)
	for a := 0; a < len(edges); a++ {
		for b := a + 1; b < len(edges); b++ {
			for c := b + 1; c < len(edges); c++ {
				for d := c + 1; d < len(edges); d++ {
					switch {
					case !isRectangle(edges[a], edges[b], edges[c], edges[d]):
					case !isHorizontalRect(edges[a], edges[b], edges[c], edges[d]):
						continue
					default:
						rectangles = append(rectangles, struct{ a, b, c, d struct{ x, y int } }{a: edges[a], b: edges[b], c: edges[c], d: edges[d]})
					}
				}
			}
		}
	}
	var count int
	for _, r := range rectangles {
		switch {
		case !isXLinePresent(in)(r.a, r.b, r.c, r.d):
		case !isYLinePresent(in)(r.a, r.b, r.c, r.d):
			continue
		default:
			count++
		}
	}
	return count
}

// CountBaseImprov counts the number of rectangles drawn from the input
func CountBaseImprov(in []string) int {
	edges := parseEdges(in)

	var count int
	xLinePresent := isXLinePresent(in)
	yLinePresent := isYLinePresent(in)
	for a := 0; a < len(edges); a++ {
		for b := a + 1; b < len(edges); b++ {
			for c := b + 1; c < len(edges); c++ {
				for d := c + 1; d < len(edges); d++ {
					switch {
					// case !isRectangle(edges[a], edges[b], edges[c], edges[d]):
					case !isHorizontalRect(edges[a], edges[b], edges[c], edges[d]):
					case !xLinePresent(edges[a], edges[b], edges[c], edges[d]):
					case !yLinePresent(edges[a], edges[b], edges[c], edges[d]):
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

func isRectangle(a, b, c, d struct{ x, y int }) bool {
	center := struct{ x, y float64 }{x: float64(a.x+b.x+c.x+d.x) / 4, y: float64(a.y+b.y+c.y+d.y) / 4}

	sqr := func(x float64) float64 { return x * x }
	sqrDist := func(a struct{ x, y int }, b struct{ x, y float64 }) float64 {
		return sqr(b.x-float64(a.x)) + sqr(b.y-float64(a.y))
	}

	dist := sqrDist(a, center)
	return dist == sqrDist(b, center) && dist == sqrDist(c, center) && dist == sqrDist(d, center)
}
