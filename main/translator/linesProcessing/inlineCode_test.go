package linesProcessing

import "testing"

func Test_addInlineCode(t *testing.T) {

	type LineTests struct {
		input  []string
		wanted []string
	}

	lineTests := LineTests{
		input: []string{
			"hello, `here is code`, here is not",
			"there isn`t any code here",
		},
		wanted: []string{
			`hello, <code>here is code</code>, here is not`,
			"there isn`t any code here",
		},
	}

	got := AddInlineCode(lineTests.input)

	if len(got) != len(lineTests.wanted) {
		t.Errorf("different count of lines, received = %v, expected =  %v", len(got), len(lineTests.wanted))
	}

	for i := 0; i < len(got); i++ {
		if got[i] != lineTests.wanted[i] {
			t.Errorf("error with adding inline code, received = %v, expected =  %v", got[i], lineTests.wanted[i])
		}
	}

}
