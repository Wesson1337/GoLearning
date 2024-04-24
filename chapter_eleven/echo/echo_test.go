package main

import (
	"bytes"
	"fmt"
	"testing"
)

func TestEcho(t *testing.T) {
	var tests = []struct {
		newline bool
		sep     string
		args    []string
		want    string
	}{
		{true, " ", []string{"a", "b", "c"}, "a b c\n"},
		{false, " ", []string{"a", "b", "c"}, "a b c"},
		{true, "\t", []string{"a", "b", "c"}, "a\tb\tc\n"},
		{false, "\t", []string{"a", "b", "c"}, "a\tb\tc"},
		{true, ",", []string{"a", "b", "c"}, "a,b,c\n"},
		{false, ",", []string{"a", "b", "c"}, "a,b,c"},
	}
	for _, test := range tests {
		descr := fmt.Sprintf("echo(%v, %q, %q)", test.newline, test.sep, test.args)
		out = new(bytes.Buffer)
		if err := echo(test.newline, test.sep, test.args); err != nil {
			t.Errorf("%s failed: %v", descr, err)
			continue
		}
		got := out.(*bytes.Buffer).String()
		if got != test.want {
			t.Errorf("%s = %q, want %q", descr, got, test.want)
		}
	}
}
