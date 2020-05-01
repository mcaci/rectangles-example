package rectangles

// CountFast1 counts the number of quadrilaterals drawn from the input
func CountFast1(in []string) int {
	edges := parseEdges(in)

	var count int
	horizComplete := isXLinePresent(in)
	vertComplete := isYLinePresent(in)
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
					case !horizComplete(edges[a], edges[b], edges[c], edges[d]):
					case !vertComplete(edges[a], edges[b], edges[c], edges[d]):
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

// CountFastDoubleAlloc counts the number of quadrilaterals drawn from the input
func CountFastDoubleAlloc(in []string) int {
	edges := parseEdges(in)

	var count int
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
					case !isYLinePresent(in)(edges[a], edges[b], edges[c], edges[d]):
					case !isXLinePresent(in)(edges[a], edges[b], edges[c], edges[d]):
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
