package namegraph


type NameGraph struct {
	nodes map[string]int
	adjacencies map[string][]string
}
