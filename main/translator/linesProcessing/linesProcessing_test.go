package linesProcessing

import "testing"

func Test_addLinks(t *testing.T) {

	type LineTests struct {
		input  []string
		wanted []string
	}

	lineTests := LineTests{
		input: []string{
			"[I'm an inline-style link](https://www.google.com)",
			"[I'm a relative reference to a repository file](../blob/master/LICENSE)",
			"[I'm a reference link][1]",
		},
		wanted: []string{
			`<a href="https://www.google.com">I'm an inline-style link</a>`,
			`<a href="../blob/master/LICENSE">I'm a relative reference to a repository file</a>`,
			`<a href="link1" title="title1">I'm a reference link</a>`,
		},
	}

	referencesMap := make(map[string]ReferencesInfo)
	referencesMap["1"] = ReferencesInfo{link: "link1", title: "title1"}

	got := AddLinks(lineTests.input, referencesMap)

	if len(got) != len(lineTests.wanted) {
		t.Errorf("different count of lines, received = %v, expected =  %v", len(got), len(lineTests.wanted))
	}

	for i := 0; i < len(got); i++ {
		if got[i] != lineTests.wanted[i] {
			t.Errorf("error with adding links, received = %v, expected =  %v", got[i], lineTests.wanted[i])
		}
	}

}
