package stringutil

import "testing"

func TestReverse(t *testing.T) {
	cases := []struct {
		input, want string
	}{
		{"Hello, world", "dlrow ,olleH"},
		{"Hello, 世界", "界世 ,olleH"},
		{"", ""},
	}
	for _, testCase := range cases {
		got := Reverse(testCase.input)
		if got != testCase.want {
			t.Errorf("Reverse(%q) == %q, want %q", testCase.input, got, testCase.want)
		}
	}
}