package rectangles

// CountFast counts the number of quadrilaterals drawn from the input
func CountFast(in []string) int {
	edges := parseEdges(in)

	var count int
	xOk := xLinePresent(in)
	yOk := yLinePresent(in)
	for a := 0; a < len(edges); a++ {
		for b := a + 1; b < len(edges); b++ {
			if !sameX(edges[a], edges[b]) {
				continue
			}
			// if !xOk(edges[a], edges[b]) {
			// 	continue
			// }
			for c := b + 1; c < len(edges); c++ {
				if !sameY(edges[a], edges[c]) {
					continue
				}
				// if !yOk(edges[a], edges[c]) {
				// 	continue
				// }
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
