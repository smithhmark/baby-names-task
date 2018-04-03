package namegraph

import "testing"
import "fmt"

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
