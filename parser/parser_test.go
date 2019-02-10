package parser_test

import (
	"strings"
	"testing"

	. "github.com/cirocosta/psi_exporter/parser"
)

func rowsEqual(a, b []*Row) bool {
	if len(a) != len(b) {
		return false
	}

	for idx, elem := range a {
		if *b[idx] != *elem {
			return false
		}
	}

	return true
}

func TestParser(t *testing.T) {
	var testCases = []struct {
		desc        string
		input       string
		expected    []*Row
		shouldError bool
	}{
		{
			desc:     "empty",
			input:    ``,
			expected: nil,
		},
		{
			desc:        "lacking fields",
			input:       `some avg10=0.00 avg60=0.00 total=1454751`,
			shouldError: true,
		},
		{
			desc:        "wrong metric type",
			input:       `some avg10=0.00 avg60=0.00 avg300=this-is-wrong total=1454751`,
			shouldError: true,
		},
		{
			desc: "well-formed some and full",
			input: `some avg10=0.00 avg60=0.00 avg300=0.13 total=1454751
	full avg10=0.00 avg60=12.20 avg300=0.00 total=454662`,
			expected: []*Row{
				{
					Type:   "some",
					Avg10:  0,
					Avg60:  0,
					Avg300: 0.13,
					Total:  1454751,
				},
				{

					Type:   "full",
					Avg10:  0,
					Avg60:  12.2,
					Avg300: 0,
					Total:  454662,
				},
			},
		},
	}

	var (
		actual []*Row
		err    error
	)

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			actual, err = Parse(strings.NewReader(tc.input))
			if tc.shouldError {
				if err == nil {
					t.Errorf("should've errored")
				}

				return
			}

			if err != nil {
				t.Errorf("should not have errored")
			}

			if !rowsEqual(actual, tc.expected) {
				t.Errorf("%+v != %+v\n", actual, tc.expected)
			}
		})
	}
}
