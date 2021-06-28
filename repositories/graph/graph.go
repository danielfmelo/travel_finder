package graph

type edge struct {
	node  string
	value int
}

type Graph struct {
	nodes map[string][]edge
}

func NewGraph() *Graph {
	return &Graph{nodes: make(map[string][]edge)}
}

func (g *Graph) AddEdge(origin, destiny string, value int) {
	g.nodes[origin] = append(g.nodes[origin], edge{node: destiny, value: value})
	g.nodes[destiny] = append(g.nodes[destiny], edge{node: origin, value: value})
}

func (g *Graph) GetEdges(node string) []edge {
	return g.nodes[node]
}

func (g *Graph) GetPath(origin, destiny string) ([]string, int) {
	h := newHeap()
	h.push(path{value: 0, nodes: []string{origin}})
	visited := make(map[string]bool)

	for len(*h.values) > 0 {
		p := h.pop()
		node := p.nodes[len(p.nodes)-1]

		if visited[node] {
			continue
		}

		if node == destiny {
			return p.nodes, p.value
		}

		for _, e := range g.GetEdges(node) {
			if !visited[e.node] {
				ns := append([]string{}, append(p.nodes, e.node)...)
				pp := path{
					value: p.value + e.value,
					nodes: ns,
				}
				h.push(pp)
			}
		}

		visited[node] = true
	}

	return nil, 0
}
