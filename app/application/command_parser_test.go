package application

import (
	"reflect"
	"testing"
)

func TestTokenize_BackslashInDoubleQuotes(t *testing.T) {
	got, err := tokenize(`"exe with \\ backslash" /tmp/ant/f4`)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	want := []string{
		`exe with \ backslash`,
		`/tmp/ant/f4`,
	}

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("got=%#v, want=%#v", got, want)
	}
}

func TestTokenizeWholeCommandLine_BackslashInDoubleQuotedExecutable(t *testing.T) {
	got, err := tokenize(`"exe with \\ backslash" /tmp/ant/f4`)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	want := []string{
		`exe with \ backslash`,
		`/tmp/ant/f4`,
	}

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("got=%#v, want=%#v", got, want)
	}
}
