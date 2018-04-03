package namegraph

import "github.com/smithhmark/dumbqueue"
import "fmt"
import "strings"

func printGraph(nm string, cnt, depth int) {
	fmt.Printf("%s%s:%v\n", strings.Repeat(" ", depth), nm, cnt)
}

type GraphVisitor interface {
	Visit(nm string, cnt, depth int)
}

type Stringer struct {
	bits []string
}

func (s *Stringer) Visit(nm string, cnt, depth int) {
	//fmt.Println("len(accum)", len(s.bits))
	//fmt.Println("cap(accum)", cap(s.bits))
	if len(s.bits) == cap(s.bits) {
		newBits := make([]string, len(s.bits), 2*cap(s.bits))
		copy(newBits, s.bits)
		s.bits = newBits
	}
	str := fmt.Sprintf("%s%s:%v", strings.Repeat(" ", depth), nm, cnt)
	//fmt.Println("adding to accum:", str)
	s.bits = append(s.bits, str)
}

type visitRecord struct{
	node string
	depth int
}

func (g *NameGraph) bfs(root string, v GraphVisitor) {
	q := dumbqueue.NewModeQueue()
	var curNode string
	var curCnt int
	var rootVR = visitRecord{root, 0}
	visits := make(map[string]visitRecord)
	q.Put(rootVR)

	for q.Size() > 0 {
		val, _ := q.Get()
		curr := val.(visitRecord)
		curNode = curr.node
		curCnt = g.nodes[curNode]
		_, visited := visits[curNode]
		if !visited {
			visits[curNode] = curr
			if adjs, ok := g.adjacencies[curNode]; ok{ 
				for _, nextNode := range adjs {
					q.Put(visitRecord{nextNode, curr.depth+1})
				}
			}
			v.Visit(curNode, curCnt, curr.depth)
		}
	}
}

type NamesSummary struct {
	current string
	Counts map[string]int
	Synonyms map[string][]string
	KeyNames map[string]string
}

func NewNamesSummary() (*NamesSummary) {
	ns := &NamesSummary{
		"",
		map[string]int{},
		map[string][]string{},
		map[string]string{},
	}
	return ns
}

func (ns *NamesSummary) Visit(nm string, cnt, depth int) {
	ns.KeyNames[nm] = ns.current
	s := ns.Synonyms[ns.current]
	if len(s) == cap(s) {
		newlist := make([]string, len(s), 2*cap(s))
		copy(newlist, s)
		ns.Synonyms[ns.current] = newlist
	}
	ns.Synonyms[ns.current] = append(ns.Synonyms[ns.current], nm)
}

func (g *NameGraph) Breakdown() (ns *NamesSummary) {
	ns = NewNamesSummary()

	for name, _ := range g.nodes {
		if _, alreadyVisited := ns.KeyNames[name]; !alreadyVisited {
			ns.current = name
			g.bfs(name, ns)
		}
	}

	for keyNode, nodes := range ns.Synonyms {
		for _, nd := range nodes {
			ns.Counts[keyNode] = g.nodes[nd] + ns.Counts[keyNode]
		}
	}
	return
}
