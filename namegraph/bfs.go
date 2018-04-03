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

