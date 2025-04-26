package hamming

var myTestCases = []struct {
	description string
	s1          string
	s2          string
	want        int
	expectError bool
}{
	{
		description: "",
		s1:          "",
		s2:          "",
		want:        0,
		expectError: false,
	},
}
