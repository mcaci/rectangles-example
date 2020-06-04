package rectangles

import "testing"

var worstCase = struct {
	description string
	input       []string
	expected    int
}{
	description: "as many rectangles as possible",
	input: []string{
		"+++++++++++++",
		"+++++++++++++",
		"+++++++++++++",
		"+++++++++++++",
		"+++++++++++++",
		"+++++++++++++",
		"+++++++++++++",
		"+++++++++++++",
	},
	expected: 2184,
}

func TestCountWorst(t *testing.T) {
	tc := worstCase
	if actual := CountAll(tc.input); actual != tc.expected {
		t.Fatalf("FAIL: %s\nExpected: %#v\nActual: %#v", tc.description, tc.expected, actual)
	}
	t.Logf("PASS: %s", tc.description)
}

func BenchmarkRectanglesWorst(b *testing.B) {
	tc := worstCase
	for i := 0; i < b.N; i++ {
		CountBase(tc.input)
	}
}

func BenchmarkRectanglesAllEdgesTakenWorst(b *testing.B) {
	tc := worstCase
	for i := 0; i < b.N; i++ {
		CountAll(tc.input)
	}
}

func BenchmarkRectanglesEdgeAndSideTogetherWorst(b *testing.B) {
	tc := worstCase
	for i := 0; i < b.N; i++ {
		CountEdgeAndSideTogether(tc.input)
	}
}

func BenchmarkRectanglesEdgesFirstWorst(b *testing.B) {
	tc := worstCase
	for i := 0; i < b.N; i++ {
		CountEdgesFirst(tc.input)
	}
}

func BenchmarkRectanglesSidesConcWorst(b *testing.B) {
	tc := worstCase
	for i := 0; i < b.N; i++ {
		CountSidesConc(tc.input)
	}
}
