package domains

import (
	"bytes"
	"testing"

	"github.com/codecrafters-io/shell-starter-go/app/infra"
)

func TestCompleteCommandPrintsRegisteredSpecification(t *testing.T) {
	store := infra.NewInMemoryCompletionsStore()
	command := NewCompleteCommand(store)

	var out bytes.Buffer
	var errOut bytes.Buffer

	_, err := command.Execute(&Command{
		Args:      []string{"-C", "/tmp/pear.py", "git"},
		Writer:    &out,
		ErrWriter: &errOut,
	})
	if err != nil {
		t.Fatalf("unexpected error registering completion: %v", err)
	}

	_, err = command.Execute(&Command{
		Args:      []string{"-p", "git"},
		Writer:    &out,
		ErrWriter: &errOut,
	})
	if err != nil {
		t.Fatalf("unexpected error printing completion: %v", err)
	}

	want := "complete -C '/tmp/pear.py' git\n"
	if got := out.String(); got != want {
		t.Fatalf("got %q, want %q", got, want)
	}
}
