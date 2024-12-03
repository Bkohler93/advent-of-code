package main

import "testing"

func TestSplitLists(t *testing.T) {
	testCases := []struct {
		input         string
		expectedLeft  []int
		expectedRight []int
	}{
		{
			input:         "1   1",
			expectedLeft:  []int{1},
			expectedRight: []int{1},
		},
		{
			input:         "12   33\n10   55\n2   44",
			expectedLeft:  []int{2, 10, 12},
			expectedRight: []int{33, 44, 55},
		},
	}

	for _, tc := range testCases {
		l, r := splitLists(tc.input)

		if len(l) != len(tc.expectedLeft) {
			t.Errorf("left lists are not same length: %d and %d", len(l), len(tc.expectedLeft))
		}

		if len(r) != len(tc.expectedRight) {
			t.Errorf("right lists are not same length: %d and %d", len(r), len(tc.expectedRight))
		}

		for i := range l {
			if l[i] != tc.expectedLeft[i] {
				t.Errorf("values are not the same: %d %d", l[i], tc.expectedLeft[i])
			}
		}

		for i := range r {
			if r[i] != tc.expectedRight[i] {
				t.Errorf("values are not the same: %d %d", r[i], tc.expectedRight[i])
			}
		}
	}
}
