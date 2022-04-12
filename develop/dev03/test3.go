package main

import (
	"bufio"
	"os"
	"testing"
	"text/scanner"
)

func TestTask(t *testing.T) {
	testTable := []struct {
		input    Config
		expected []string
	}{
		{
			input: Config{
				K:        0,
				Filename: "",
				N:        false,
				R:        false,
				U:        false,
			},
			expected: []string{"44477ddeeeeffffhhjjjkkkksuwxyyyz", "44477ddeeeeffffhhjjjkkkksuwxyyyz",
				"444777eefiijklrttttuxyyyyz", "ads", "ads", "ads", "aaaaddddssss", "3addddffffgghss"},
		},
		{
			input: Config{
				K:        1,
				Filename: "",
				N:        false,
				R:        false,
				U:        false,
			},
			expected: []string{"44477ddeeeeffffhhjjjkkkksuwxyyyz", "kufdehjwyskf4ehde7fyejzk474fxjyk",
				"ui4jxytrfklezi7yt47tye47ty", "asd", "asd", "dsa", "asdasdasdasd", "asfdgdfgfdhsdf3"},
		},
		{
			input: Config{
				K:        0,
				Filename: "",
				N:        true,
				R:        false,
				U:        false,
			},
			expected: []string{"44477", "44477", "444777", "3"},
		},
		{
			input: Config{
				K:        0,
				Filename: "",
				N:        false,
				R:        true,
				U:        false,
			},
			expected: []string{"zyyyxwuskkkkjjjhhffffeeeedd77444", "zyyyxwuskkkkjjjhhffffeeeedd77444",
				"zyyyyxuttttrlkjiifee777444", "sda", "sda", "sda", "ssssddddaaaa", "sshggffffdddda3"},
		},
		{
			input: Config{
				K:        0,
				Filename: "",
				N:        false,
				R:        false,
				U:        true,
			},
			expected: []string{"44477ddeeeeffffhhjjjkkkksuwxyyyz", "444777eefiijklrttttuxyyyyz", "ads", "aaaaddddssss",
				"3addddffffgghss"},
		},
	}

	f, err := os.Open("develop\\dev03\\file.txt")
	if err != nil {
		t.Errorf("error with open file")
	}

	var s scanner.Scanner
	s.Init(f)
	s.Whitespace ^= 1<<' ' | 1<<'\n' // don't skip tabs and new lines

	text := make([][]string, 0)
	var helpstring []string
	for tok := s.Scan(); tok != scanner.EOF; tok = s.Scan() {
		switch tok {
		case '\n':
			str := make([]string, len(helpstring))
			copy(str, helpstring)
			text = append(text, str)
			helpstring = helpstring[:0]
		case ' ':
			continue
		default:
			helpstring = append(helpstring, s.TokenText())
		}
	}
	str := make([]string, len(helpstring))
	copy(str, helpstring)
	text = append(text, str)

	for _, tt := range testTable {
		got := SortFile(text, tt.input)
		if Equal(got, tt.expected) {
			t.Errorf("Expected %s, got %s", tt.expected, got)
		}
	}
}

func Equal(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
