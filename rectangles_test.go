package rectangles

import (
	"testing"
)

func TestRectangles(t *testing.T) {
	for _, tc := range testCases {
		if actual := Count(tc.input); actual != tc.expected {
			t.Fatalf("FAIL: %s\nExpected: %#v\nActual: %#v", tc.description, tc.expected, actual)
		}
		t.Logf("PASS: %s", tc.description)
	}
}

func BenchmarkRectangles(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testCases {
			Count(tc.input)
		}
	}
}

func BenchmarkRectanglesBaseImprov(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testCases {
			CountBaseImprov(tc.input)
		}
	}
}

func BenchmarkRectanglesFast(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testCases {
			CountFast(tc.input)
		}
	}
}

func BenchmarkRectanglesFastDoubleAlloc(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testCases {
			CountFastDoubleAlloc(tc.input)
		}
	}
}

func BenchmarkRectanglesConc(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testCases {
			CountConc(tc.input)
		}
	}
}

func BenchmarkRectanglesConcPrefQuad(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testCases {
			CountConcPrefQuad(tc.input)
		}
	}
}