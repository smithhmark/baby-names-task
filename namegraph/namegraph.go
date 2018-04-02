package namegraph

type NameGraph struct {
	nodes map[string]int
	adjacencies map[string][]string
}

func NewNameGraph() (*NameGraph) {
	g := &NameGraph{map[string]int{}, map[string][]string{}}
	return g
}

func (g *NameGraph) Size() int {
	return len(g.nodes)
}

func (g *NameGraph) AddName(nm string) {
	g.nodes[nm] = 0
}
