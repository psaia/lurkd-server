package github

import (
	"testing"
)

func TestInitialize(t *testing.T) {
	p := Github{
		ModuleName: "facebook/React",
	}

	v, err := p.Initialize()

	if err != nil {
		t.Fatalf("expected err to be nil. got: %v", err)
	}

	if v == "" {
		t.Fatal("did not expect version to be empty.")
	}
}
