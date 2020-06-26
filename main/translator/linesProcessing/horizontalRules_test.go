package linesProcessing

import "testing"

func Test_addHorizontalRules(t *testing.T) {

	type LineTests struct {
		input  []string
		wanted []string
	}

	lineTests := LineTests{
		input: []string{
			"***werwerwerrs",
			"___ bla bla",
			"--- nice",
			"_ not a match",
			"-- not a match",
			"--**__ not a match",
		},
		wanted: []string{
			"<hr>",
			"<hr>",
			"<hr>",
			"_ not a match",
			"-- not a match",
			"--**__ not a match",
		},
	}

	got := AddHorizontalRules(lineTests.input)

	if len(got) != len(lineTests.wanted) {
		t.Errorf("different count of lines, received = %v, expected =  %v", len(got), len(lineTests.wanted))
	}

	for i := 0; i < len(got); i++ {
		if got[i] != lineTests.wanted[i] {
			t.Errorf("error with adding inline code, received = %v, expected =  %v", got[i], lineTests.wanted[i])
		}
	}

}
