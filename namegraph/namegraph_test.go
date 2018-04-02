package namegraph

import (
	"testing"
)

func TestConstruction(t *testing.T) {
	g1 := new(NameGraph)
	if g1.Size() != 0 {
		t.Fatalf("initial size should be Zero")
	}
	if len(g1.adjacencies) != 0 {
		t.Fatalf("initial adjacency list should be empty")
	}
	g2 := NewNameGraph()
	if g2.Size() != 0 {
		t.Fatalf("initial size should be Zero")
	}
	if len(g2.adjacencies) != 0 {
		t.Fatalf("initial adjacency list should be empty")
	}

}

func TestAddName(t *testing.T) {
	g := NewNameGraph()

	if g.Size() != 0 {
		t.Fatalf("initial size should be Zero")
	}
	g.AddName("Sue")
	if g.Size() != 1 {
		t.Fatalf("size should be One after an add")
	}
	g.AddName("Sue")
	if g.Size() != 1 {
		t.Fatalf("size should STILL be One after a repeated add")
	}
}

func TestAddAdjacency(t *testing.T) {
	g := NewNameGraph()

	g.AddAdjacency("Sue", "Susan")
	if len(g.adjacencies) != 2 {
		t.Fatalf("Failed to add both names to the adjacency list")
	}
	if len(g.nodes) != 2 {
		t.Fatalf("Failed to add both names to the node list")
	}
	g.AddAdjacency("Chris", "Kris")
	g.AddAdjacency("Chris", "Kristopher")
	g.AddAdjacency("Chris", "Cristopher")
	if len(g.adjacencies["Chris"]) != 3 {
		t.Fatalf("Failed to grow adjacency list")
	}
}

func TestSetCount(t *testing.T) {
	g := NewNameGraph()
	g.SetCount("Sue", 42)
	if g.nodes["Sue"] != 42 {
		t.Fatalf("Failed to add name and set count")
	}

	g.SetCount("Sue", 420)
	if g.nodes["Sue"] != 420 {
		t.Fatalf("Failed to change and existing count")
	}
}

