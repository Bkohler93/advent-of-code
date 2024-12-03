package main

import "testing"

func TestParseReports(t *testing.T) {
	testCases := []struct {
		input    string
		expected []report
	}{
		{
			input: "1",
			expected: []report{{
				levels: []int{1},
				isSafe: false,
			}},
		},
		{
			input: "1 2 3\n2 3 4",
			expected: []report{{
				levels: []int{1, 2, 3},
				isSafe: false,
			}, {
				levels: []int{2, 3, 4},
				isSafe: false,
			}},
		},
	}

	for _, tc := range testCases {
		reports := parseReports(tc.input)

		if len(reports) != len(tc.expected) {
			t.Errorf("report lengths incorrect - %d and %d", len(reports), len(tc.expected))
		}

		for i := range reports {
			ls := reports[i].levels
			tLs := tc.expected[i].levels
			if len(ls) != len(tLs) {
				t.Errorf("incorrect number of leves - %d and %d", len(ls), len(tLs))
			}

			for j := range ls {
				if ls[j] != tLs[j] {
					t.Errorf("incorrect level - %d and %d", ls[i], tLs[i])
				}
			}
		}
	}
}

func TestDetermineSafetyTolerate(t *testing.T) {
	testCases := []struct {
		report   report
		expected bool
		reason   string
	}{
		{
			report: report{
				levels: []int{62, 63, 66, 68, 70, 73},
			},
			expected: true,
		},
		{
			report: report{
				levels: []int{3, 2, 4, 1},
			},
			expected: true,
			reason:   "remove 4 to make true",
		},
		{
			report: report{
				levels: []int{1, 2, 7, 8, 9},
			},
			expected: false,
			reason:   "2 7 is an increase of 5",
		},
		{
			report: report{
				levels: []int{8, 6, 4, 4, 1},
			},
			expected: true,
			reason:   "remove 4 to make true",
		},
		{
			report: report{
				levels: []int{1, 3, 2, 4, 5},
			},
			expected: true,
			reason:   "remove 3 to make true",
		},
		{
			report: report{
				levels: []int{8, 6, 4, 4, 1},
			},
			expected: true,
			reason:   "remove 4 to make true",
		},
	}

	for _, tc := range testCases {
		r := tc.report
		r.determineSafetyTolerate()

		if r.isSafe != tc.expected {
			t.Errorf("%#v - expected %t got %t: %s", r, tc.expected, r.isSafe, tc.reason)
		}
	}
}

func TestRemoveIndex(t *testing.T) {
	testCases := []struct {
		slice    []int
		index    int
		original []int
		expected []int
	}{
		{
			slice:    []int{1, 2, 3},
			index:    1,
			original: []int{1, 2, 3},
			expected: []int{1, 3},
		},
	}

	for _, tc := range testCases {
		_ = removeIndex(tc.slice, tc.index)

		for i := 0; i < len(tc.original); i++ {
			if tc.original[i] != tc.slice[i] {
				t.Errorf("%v - %v", tc.original, tc.slice)
			}
		}
	}
}

func TestDetermineSafety(t *testing.T) {
	testCases := []struct {
		report   report
		expected bool
		reason   string
	}{
		{
			report: report{
				levels: []int{1, 2, 3},
			},
			expected: true,
		},
		{
			report: report{
				levels: []int{1, 6},
			},
			expected: false,
			reason:   "1 6 is an increase of 5",
		},
		{
			report: report{
				levels: []int{3, 2, 1},
			},
			expected: true,
		},
		{
			report: report{
				levels: []int{6, 1},
			},
			expected: false,
			reason:   "6 1 is a decrease of 5",
		},
		{
			report: report{
				levels: []int{62, 63, 66, 68, 70, 73},
			},
			expected: true,
		},
		{
			report: report{
				levels: []int{3, 2, 4, 1},
			},
			expected: false,
			reason:   "3 2 4 goes up and down",
		},
		{
			report: report{
				levels: []int{1, 2, 7, 8, 9},
			},
			expected: false,
			reason:   "2 7 is an increase of 5",
		},
		{
			report: report{
				levels: []int{8, 6, 4, 4, 1},
			},
			expected: false,
			reason:   "4 4 is neither up nor down",
		},
	}

	for _, tc := range testCases {
		r := tc.report
		r.determineSafety()

		if r.isSafe != tc.expected {
			t.Errorf("%#v - expected %t got %t", r, tc.expected, r.isSafe)
		}
	}
}

func TestSumSafeReports(t *testing.T) {
	testCases := []struct {
		reports  []report
		expected int
	}{
		{
			reports: []report{
				{
					levels: []int{1, 2, 3},
					isSafe: true,
				},
				{
					levels: []int{1, 6},
					isSafe: false,
				},
				{
					levels: []int{3, 2, 1},
					isSafe: true,
				},
				{
					levels: []int{6, 1},
					isSafe: false,
				},
			},
			expected: 2,
		},
		{
			reports: []report{
				{
					levels: []int{1, 5},
					isSafe: false,
				},
			},
			expected: 0,
		},
	}

	for _, tc := range testCases {
		actual := sumSafeReports(tc.reports)
		if actual != tc.expected {
			t.Errorf("sums do not match - %d and %d", actual, tc.expected)
		}
	}
}

func TestAnalyzeReports(t *testing.T) {
	testCases := []struct {
		reports                 []report
		expectedAnalyzedReports []report
	}{
		{
			reports: []report{
				{
					levels: []int{1, 2, 3},
					isSafe: false,
				},
				{
					levels: []int{3, 2, 1},
					isSafe: false,
				},
				{
					levels: []int{1, 2, 1},
					isSafe: false,
				},
			},
			expectedAnalyzedReports: []report{
				{
					levels: []int{1, 2, 3},
					isSafe: true,
				},
				{
					levels: []int{3, 2, 1},
					isSafe: true,
				},
				{
					levels: []int{1, 2, 1},
					isSafe: false,
				},
			},
		},
	}

	for _, tc := range testCases {
		rs := tc.reports

		rs = analyzeReports(rs)

		for i := range rs {
			if rs[i].isSafe != tc.expectedAnalyzedReports[i].isSafe {
				t.Errorf("got %t expected %t", rs[i].isSafe, tc.expectedAnalyzedReports[i].isSafe)
			}
		}
	}
}
