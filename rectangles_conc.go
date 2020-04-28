package rectangles

import "sync/atomic"

// CountConc counts the number of quadrilaterals drawn from the input
func CountConc(in []string) int {
	edges := parseEdges(in)

	var count int
	type inParseFunc func(in []string) func(a, b struct{ x, y int }) bool
	lineCheck := func(check inParseFunc, a, b struct{ x, y int }, c chan<- bool) { c <- !check(in)(a, b) }
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
					go lineCheck(xLinePresent, edges[a], edges[b], checkChan)
					go lineCheck(xLinePresent, edges[c], edges[d], checkChan)
					go lineCheck(yLinePresent, edges[a], edges[c], checkChan)
					go lineCheck(yLinePresent, edges[b], edges[d], checkChan)

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

// CountConcPrefQuad counts the number of quadrilaterals drawn from the input
func CountConcPrefQuad(in []string) int {
	edges := parseEdges(in)

	quadrilaterals := make([]struct{ a, b, c, d struct{ x, y int } }, 0)
	for a := 0; a < len(edges); a++ {
		for b := a + 1; b < len(edges); b++ {
			for c := b + 1; c < len(edges); c++ {
				for d := c + 1; d < len(edges); d++ {
					if !isPlainRect(edges[a], edges[b], edges[c], edges[d]) {
						continue
					}
					quadrilaterals = append(quadrilaterals, struct{ a, b, c, d struct{ x, y int } }{a: edges[a], b: edges[b], c: edges[c], d: edges[d]})
				}
			}
		}
	}
	var count uint32
	for _, r := range quadrilaterals {
		go func(r struct{ a, b, c, d struct{ x, y int } }) {
			switch {
			case !isYLinePresent(in)(r.a, r.b, r.c, r.d):
			case !isXLinePresent(in)(r.a, r.b, r.c, r.d):
				return
			default:
				atomic.AddUint32(&count, 1)
			}
		}(r)
	}
	return int(count)
}
