package rectangles

import (
	"testing"
)

var avgCase = struct {
	description string
	input       []string
	expected    int
}{
	description: "large input with many rectangles",
	input: []string{
		"+++-+--+--+-+",
		"|++ +--+--+-+",
		"+++-+--+++++|",
		"|++++--+----+",
		"+---+--+--+-+",
		"+---+--+--+-+",
		"+--+--++  | |",
		"++++--++  +-+",
	},
	expected: 102,
}

func TestCountAvg(t *testing.T) {
	tc := avgCase
	if actual := CountAll(tc.input); actual != tc.expected {
		t.Fatalf("FAIL: %s\nExpected: %#v\nActual: %#v", tc.description, tc.expected, actual)
	}
	t.Logf("PASS: %s", tc.description)
}

func BenchmarkRectanglesAvg(b *testing.B) {
	tc := avgCase
	for i := 0; i < b.N; i++ {
		CountBase(tc.input)
	}
}

func BenchmarkRectanglesAllEdgesTakenAvg(b *testing.B) {
	tc := avgCase
	for i := 0; i < b.N; i++ {
		CountAll(tc.input)
	}
}

func BenchmarkRectanglesEdgeAndSideTogetherAvg(b *testing.B) {
	tc := avgCase
	for i := 0; i < b.N; i++ {
		CountEdgeAndSideTogether(tc.input)
	}
}

func BenchmarkRectanglesEdgesFirstAvg(b *testing.B) {
	tc := avgCase
	for i := 0; i < b.N; i++ {
		CountEdgesFirst(tc.input)
	}
}

func BenchmarkRectanglesSidesConcAvg(b *testing.B) {
	tc := avgCase
	for i := 0; i < b.N; i++ {
		CountSidesConc(tc.input)
	}
}
