package linesProcessing

import "testing"

func Test_addImages(t *testing.T) {

	type LineTests struct {
		input  []string
		wanted []string
	}

	lineTests := LineTests{
		input: []string{
			"![Minion](https://octodex.github.com/images/minion.png)",
			"![Stormtroopocat](https://octodex.github.com/images/stormtroopocat.jpg \"The Stormtroopocat\")",
		},
		wanted: []string{
			`<img alt="Minion" src="https://octodex.github.com/images/minion.png" />`,
			`<img alt="Stormtroopocat" src="https://octodex.github.com/images/stormtroopocat.jpg" title="The Stormtroopocat" />`,
		},
	}

	referencesMap := make(map[string]ReferencesInfo)

	got := AddImages(lineTests.input, referencesMap)

	if len(got) != len(lineTests.wanted) {
		t.Errorf("different count of lines, received = %v, expected =  %v", len(got), len(lineTests.wanted))
	}

	for i := 0; i < len(got); i++ {
		if got[i] != lineTests.wanted[i] {
			t.Errorf("error with adding links, received = %v, expected =  %v", got[i], lineTests.wanted[i])
		}
	}

}
