package main

import "testing"

type testPair struct {
	value    string
	unpacked string
}

var tests = []testPair{
	{"", ""},
	{"///230/a10", "/222222222222222222222222222222aaaaaaaaaa"},
	{"abcd", "abcd"},
	{"100", ""},
	{"a5", "aaaaa"},
	{"a10", "aaaaaaaaaa"},
	{"10a", ""},
	{"/", ""},
	{"////", "//"},
	{"///", ""},
	{"/a/b/c", "abc"},
	{"/a3", "aaa"},
	{"a/3", "a3"},
}

func Test_unpack(t *testing.T) {
	for _, pair := range tests {
		v := unpack(pair.value)
		if v != pair.unpacked {
			t.Error(
				"For", pair.value,
				"expected", pair.unpacked,
				"got", v,
			)
		}
	}
}
