package linesProcessing

import "testing"

func Test_addEmphasis(t *testing.T) {

	type LineTests struct {
		input  []string
		wanted []string
	}

	lineTests := LineTests{
		input: []string{
			"hello, **that's bolded**, that's not",
			"hello, __that's bolded__, that's not",
			"hello, *that's italic*, that's not",
			"hello, _that's italic_, that's not",
			"hello, ~~that's deleted~~, that's not",
		},
		wanted: []string{
			"hello, <strong>that's bolded</strong>, that's not",
			"hello, <strong>that's bolded</strong>, that's not",
			"hello, <em>that's italic</em>, that's not",
			"hello, <em>that's italic</em>, that's not",
			"hello, <del>that's deleted</del>, that's not",
		},
	}

	got := AddEmphasis(lineTests.input)

	if len(got) != len(lineTests.wanted) {
		t.Errorf("different count of lines, received = %v, expected =  %v", len(got), len(lineTests.wanted))
	}

	for i := 0; i < len(got); i++ {
		if got[i] != lineTests.wanted[i] {
			t.Errorf("error with adding emphasis, received = %v, expected =  %v", got[i], lineTests.wanted[i])
		}
	}

}
