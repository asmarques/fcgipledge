package fcgipledge

import "testing"

func TestCreatePromiseString(t *testing.T) {
	type test struct {
		input    []string
		expected string
	}
	tests := []test{
		{nil, "stdio unix"},
		{[]string{}, "stdio unix"},
		{[]string{Inet}, "stdio unix inet"},
		{[]string{Inet, DNS}, "stdio unix inet dns"},
	}

	for _, test := range tests {
		result := createPromiseString(test.input)
		if result != test.expected {
			t.Fatalf("for input %v, expected '%s', got '%s'\n", test.input, test.expected, result)
		}
	}
}
