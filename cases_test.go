package rectangles

var testCases = []struct {
	description string
	input       []string
	expected    int
}{
	{
		description: "large input with many many rectangles",
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
	},
}
