package main

import "testing"

func TestLevelIndex(t *testing.T) {
	var tests = []struct {
		current   int
		levelSize int
		max       int
		want      int
	}{
		{999999, 15000, 120000, 7},
		{120000, 15000, 120000, 7},

		{119999, 15000, 120000, 6},
		{105000, 15000, 120000, 6},

		{104999, 15000, 120000, 5},
		{90000, 15000, 120000, 5},

		{89999, 15000, 120000, 4},
		{75000, 15000, 120000, 4},

		{74999, 15000, 120000, 3},
		{60000, 15000, 120000, 3},

		{59999, 15000, 120000, 2},
		{45000, 15000, 120000, 2},

		{44999, 15000, 120000, 1},
		{30000, 15000, 120000, 1},

		{29999, 15000, 120000, 0},
		{15000, 15000, 120000, 0},
		{0, 15000, 120000, 0},
		{-1, 15000, 120000, 0},
	}
	for _, test := range tests {
		if got := levelIndex(test.current, test.levelSize, test.max); got != test.want {
			t.Errorf("LevelIndex(%d, %d, %d) = %v (%v)", test.current, test.levelSize, test.max, got, test.want)
		}
	}
}

func TestLevelGlyph(t *testing.T) {
	// NB for some small ɛ
	var tests = map[string]struct {
		current   int
		levelSize int
		max       int
		want      string
	}{
		">1": {current: 999999, levelSize: 15000, max: 120000, want: "█"},
		"1":  {current: 120000, levelSize: 15000, max: 120000, want: "█"},

		"1-ɛ": {current: 119999, levelSize: 15000, max: 120000, want: "▇"},
		"⅞":   {current: 105000, levelSize: 15000, max: 120000, want: "▇"},

		"⅞-ɛ": {current: 104999, levelSize: 15000, max: 120000, want: "▆"},
		"¾":   {current: 90000, levelSize: 15000, max: 120000, want: "▆"},

		"¾-ɛ": {current: 89999, levelSize: 15000, max: 120000, want: "▅"},
		"⅝":   {current: 75000, levelSize: 15000, max: 120000, want: "▅"},

		"⅝-ɛ": {current: 74999, levelSize: 15000, max: 120000, want: "▄"},
		"½":   {current: 60000, levelSize: 15000, max: 120000, want: "▄"},

		"½-ɛ": {current: 59999, levelSize: 15000, max: 120000, want: "▃"},
		"⅜":   {current: 45000, levelSize: 15000, max: 120000, want: "▃"},

		"⅜-ɛ": {current: 44999, levelSize: 15000, max: 120000, want: "▂"},
		"¼":   {current: 30000, levelSize: 15000, max: 120000, want: "▂"},

		"¼-ɛ": {current: 29999, levelSize: 15000, max: 120000, want: "▁"},
		"⅛":   {current: 15000, levelSize: 15000, max: 120000, want: "▁"},

		"⅛-ɛ": {current: 14999, levelSize: 15000, max: 120000, want: "‼"},
		"0":   {current: 0, levelSize: 15000, max: 120000, want: "‼"},
		"<0":  {current: -1, levelSize: 15000, max: 120000, want: "‼"},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			if got := LevelGlyph(tc.current, tc.levelSize, tc.max); got != tc.want {
				t.Errorf("wanted: %s, got %s", tc.want, got)
			}
		})
	}
}

func TestInc(t *testing.T) {
	tests := []struct {
		current   int
		levelSize int
		max       int
		want      int
	}{
		{-1, 15000, 120000, 15000},
		{0, 15000, 120000, 15000},

		{15000, 15000, 120000, 30000},

		{120000, 15000, 120000, 120000},
		{999999, 15000, 120000, 120000},
	}
	for _, test := range tests {
		if got := inc(test.current, test.levelSize, test.max); got != test.want {
			t.Errorf("increment(%d, %d, %d) = %v", test.current, test.levelSize, test.max, got)
		}
	}
}

func TestDec(t *testing.T) {
	tests := []struct {
		current   int
		levelSize int
		max       int
		want      int
	}{
		{-1, 15000, 120000, 15000},
		{0, 15000, 120000, 15000},
		{15000, 15000, 120000, 15000},

		{30000, 15000, 120000, 15000},

		{120000, 15000, 120000, 105000},
		{999999, 15000, 120000, 105000},
	}
	for _, test := range tests {
		if got := dec(test.current, test.levelSize, test.max); got != test.want {
			t.Errorf("dec(%d, %d, %d) = %v", test.current, test.levelSize, test.max, got)
		}
	}
}

func TestLevelSize(t *testing.T) {
	tests := map[string]struct {
		max  int
		want int
	}{
		"zero": {max: 0, want: 0},
		"max":  {max: 120000, want: 120000 / (len([]rune(levels)))},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			if got := levelSize(tc.max); got != tc.want {
				t.Errorf("wanted: %v, got: %v", tc.want, got)
			}
		})
	}
}

func TestNoop(t *testing.T) {
	current := 30000
	levelSize := 15000
	max := 120000
	if got := noop(current, levelSize, max); got != current {
		t.Errorf("noop(%d, %d, %d) = %v (%v)", current, levelSize, max, current, current)
	}
}
