package hamming

import "testing"

func MyTestHamming(t *testing.T) {
	for _, tc := range myTestCases {
		t.Run(tc.description, func(t *testing.T) {
			got, err := Distance(tc.s1, tc.s2)
			switch {
			case tc.expectError:
				if err == nil {
					t.Fatalf("Distance(%q, %q) expected error, got: %d", tc.s1, tc.s2, got)
				}
			case err != nil:
				t.Fatalf("Distance(%q, %q) returned error: %v, want: %d", tc.s1, tc.s2, err, tc.want)
			case got != tc.want:
				t.Fatalf("Distance(%q, %q) = %d, want: %d", tc.s1, tc.s2, got, tc.want)
			}
		})
	}
}

func MyBenchmarkHamming(b *testing.B) {
	if testing.Short() {
		b.Skip("skipping benchmark in short mode.")
	}
	for i := 0; i < b.N; i++ {
		for _, tc := range myTestCases {
			_, _ = Distance(tc.s1, tc.s2)
		}
	}
}
