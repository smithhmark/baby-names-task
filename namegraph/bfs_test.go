package namegraph

import "testing"
//import "fmt"

func TestStringer(t *testing.T) {
	g := NewNameGraph()

	g.AddAdjacency("Sally", "Sue")
	g.AddAdjacency("Sally", "Steve")
	g.AddAdjacency("Sally", "Sven")
	g.AddAdjacency("Wally", "Sven")

	s := Stringer{make([]string, 0, 2)}

	g.bfs("Sally", &s)
	
	expected := []string{
		"Sally:0",
		" Sue:0",
		" Steve:0",
		" Sven:0",
		"  Wally:0",
	}
	/*
	for i := 0 ; i < len(expected); i++ {
		fmt.Println(expected[i])
	}
	*/
	if g.Size() != len(s.bits) {
		t.Fatalf("expected one line per node, got %v", len(s.bits))
	}
	for i := 0 ; i < len(expected); i++ {
		if s.bits[i] != expected[i] {
			t.Fatalf("%s != %s\n", s.bits[i], expected[i])
		}
	}
}

func TestBreakdown(t *testing.T) {
	g := NewNameGraph()
	g.AddAdjacency("Susan", "Sue")
	g.AddAdjacency("Sue", "Susanne")
	g.AddName("Wilma")
	g.AddAdjacency("Mike", "Michael")
	g.AddAdjacency("Wally", "Walter")
	g.AddName("Milly")

	g.SetCount("Milly", 10)
	g.SetCount("Susan", 10)
	g.SetCount("Sue", 10)
	g.SetCount("Susanne", 10)
	g.SetCount("Wilma", 10)
	g.SetCount("Mike", 10)
	g.SetCount("Michael", 10)
	g.SetCount("Wally", 10)
	g.SetCount("Walter", 10)
	ns := g.Breakdown()
	if val, ok := ns.KeyNames["Sue"]; ok {
		if ns.Counts[val] != 30 {
			t.Fatalf("Failed to capture the counts across the entire subgraph")
		}
	} else {
		t.Fatalf("Failed to map name to Keyname")
	}

	if ns.KeyNames["Wally"] != ns.KeyNames["Walter"] {
		t.Fatalf("Adjacent names should share a keynode")
	}
}
