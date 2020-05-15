package rectangles

// CountSidesConc counts the number of quadrilaterals drawn from the input
func CountSidesConc(in []string) int {
	edges := parseEdges(in)

	var count int
	xOk := xlineFilled(in)
	yOk := ylineFilled(in)
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
					go func() { checkChan <- !xOk(eA, eB) }()
					go func() { checkChan <- !xOk(eC, eD) }()
					go func() { checkChan <- !yOk(eA, eC) }()
					go func() { checkChan <- !yOk(eB, eD) }()

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

// CountEdgeAndSideConc counts the number of quadrilaterals drawn from the input
func CountEdgeAndSideConc(in []string) int {
	edges := parseEdges(in)

	var count int
	xOk := xlineFilled(in)
	yOk := ylineFilled(in)
	for a := 0; a < len(edges); a++ {
		for b := a + 1; b < len(edges); b++ {
			if !sameX(edges[a], edges[b]) {
				continue
			}
			if !xOk(edges[a], edges[b]) {
				continue
			}
		}
	}

	for a := 0; a < len(edges); a++ {
		for b := a + 1; b < len(edges); b++ {
			if !sameY(edges[a], edges[b]) {
				continue
			}
			if !yOk(edges[a], edges[b]) {
				continue
			}
		}
	}

	for a := 0; a < len(edges); a++ {
		for b := a + 1; b < len(edges); b++ {
			if !sameX(edges[a], edges[b]) {
				continue
			}
			if !xOk(edges[a], edges[b]) {
				continue
			}
		}
	}

	for a := 0; a < len(edges); a++ {
		for b := a + 1; b < len(edges); b++ {
			if !sameX(edges[a], edges[b]) {
				continue
			}
			if !xOk(edges[a], edges[b]) {
				continue
			}
		}
	}

	return count
}
