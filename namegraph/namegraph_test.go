package namegraph

import (
	"testing"
)

func TestConstruction(t *testing.T) {
	g1 := new(NameGraph)
	if g1.Size() != 0 {
		t.Fatalf("initial size should be Zero")
	}
	g2 := NewNameGraph()
	if g2.Size() != 0 {
		t.Fatalf("initial size should be Zero")
	}

}

func TestAdd(t *testing.T) {
	g := NewNameGraph()

	if g.Size() != 0 {
		t.Fatalf("initial size should be Zero")
	}
	g.AddName("Sue")
	if g.Size() != 1 {
		t.Fatalf("size should be One after an add")
	}

}
