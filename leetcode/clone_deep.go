package main

func main() {

}

func cloneGraph(graph *node) *node {
	nodeMap := make(map[string]*node)
	var cloneDeep func(graph *node) *node
	cloneDeep = func(graph *node) *node {
		if graph == nil {
			return graph
		}
		_, ok := nodeMap[graph.label]
		if !ok {
			nodeMap[graph.label] = &node{label: graph.label}
			for _, n := range graph.neighbors {
				nodeMap[graph.label].neighbors = append(nodeMap[graph.label].neighbors, cloneDeep(n))
			}
		}
		return nodeMap[graph.label]
	}

	return cloneDeep(graph)
}

type node struct {
	label     string
	neighbors []*node
}
