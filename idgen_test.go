package idgen_test

import (
	"testing"

	"github.com/alfzs/idgen"
)

func TestGenerate(t *testing.T) {
	id, err := idgen.Generate(9)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(id) < 12 {
		t.Errorf("expected at least 12 characters, got %d: %s", len(id), id)
	}
}
