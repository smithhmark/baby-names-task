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
	_, ok := g.nodes[nm]
	if !ok {
		g.nodes[nm] = 0
	}
}

func (g *NameGraph) SetCount(name string, cnt int) {
	g.nodes[name] = cnt
}

func (g *NameGraph) addAdjacency(s1, s2 string) {
	nodes, ok := g.adjacencies[s1]
	if !ok {
		nodes = make([]string, 0, 2)
	} else {
		if len(nodes) == cap(nodes) {
			newNodes := make([]string, len(nodes), 2*cap(nodes))
			copy(newNodes, nodes)
			nodes = newNodes
		}
	}
	g.adjacencies[s1] = append(nodes, s2)
}

func (g *NameGraph) AddAdjacency(s1, s2 string) {
	g.AddName(s1)
	g.AddName(s2)
	g.addAdjacency(s1, s2)
	g.addAdjacency(s2, s1)
}

func (g *NameGraph) sever(s1, s2 string) {
	nodes, ok := g.adjacencies[s1]
	if ok {
		idxToRm := -1
		for idx, val := range nodes {
			if val == s2 {
				idxToRm = idx
				break
			}
		}
		if idxToRm > -1 {
			g.adjacencies[s1] = append(nodes[:idxToRm], nodes[idxToRm+1:]...)
		}
	}
}

func (g *NameGraph) RemoveAdjacency(s1, s2 string) {
	g.sever(s1, s2)
	g.sever(s2, s1)
}
