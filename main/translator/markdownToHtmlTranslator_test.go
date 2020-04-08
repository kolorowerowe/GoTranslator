package translator

import "testing"

func Test_wrapInHtml(t *testing.T) {
	type args struct {
		content  string
		fileName string
	}
	tests := []struct {
		args args
		want string
	}{
		{args{"Hello there!", "naaaame.md"}, `<html lang="en"><head><meta charset="utf-8"><title>naaaame.md</title></head><body>Hello there!</body></html>`},
	}
	for _, tt := range tests {
		if got := wrapInHtml(tt.args.content, tt.args.fileName); got != tt.want {
			t.Errorf("wrapInHtml() = %v, want %v", got, tt.want)
		}
	}
}
