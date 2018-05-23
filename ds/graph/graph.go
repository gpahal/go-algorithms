package graph

type Node struct {
	ID    int
	Value int
}

type Edge struct {
	SourceID int
	TargetID int
	Weight   int
}

type Graph struct {
	nodes             map[int]*Node
	edges             map[int]map[int]int
	edgesReverseIndex map[int]map[int]int

	currID int
}

func New() *Graph {
	return &Graph{
		nodes:             make(map[int]*Node),
		edges:             make(map[int]map[int]int),
		edgesReverseIndex: make(map[int]map[int]int),
		currID:            1,
	}
}

func (g *Graph) NodeCount() int {
	return len(g.nodes)
}

func (g *Graph) Node(id int) *Node {
	node, ok := g.nodes[id]
	if ok {
		return node
	}

	return nil
}

func (g *Graph) HasNode(id int) bool {
	return g.Node(id) != nil
}

func (g *Graph) AddNode(value int) {
	g.nodes[g.currID] = &Node{
		ID:    g.currID,
		Value: value,
	}

	g.currID++
}

func (g *Graph) UpdateNode(id, value int) bool {
	node, ok := g.nodes[id]
	if ok {
		node.Value = value
		return true
	}

	return false
}

func (g *Graph) DeleteNode(id int) bool {
	_, ok := g.nodes[id]
	if ok {
		delete(g.nodes, id)
		return true
	}

	return false
}

func (g *Graph) Edge(sourceID, targetID int) *Edge {
	ett, ok := g.edges[sourceID]
	if ok {
		w, ok := ett[targetID]
		if ok {
			return &Edge{
				SourceID: sourceID,
				TargetID: targetID,
				Weight:   w,
			}
		}
	}

	return nil
}

func (g *Graph) HasEdge(sourceID, targetID int) bool {
	return g.Edge(sourceID, targetID) != nil
}

func (g *Graph) AddEdge(sourceID, targetID, weight int) bool {
	_, ok := g.nodes[sourceID]
	if !ok {
		return false
	}

	_, ok = g.nodes[targetID]
	if !ok {
		return false
	}

	ett, ok := g.edges[sourceID]
	if ok {
		_, ok = ett[targetID]
		if ok {
			return false
		}

		ett[targetID] = weight
	} else {
		ett := make(map[int]int, 1)
		ett[targetID] = weight

		g.edges[sourceID] = ett
	}

	ett, ok = g.edgesReverseIndex[targetID]
	if ok {
		ett[sourceID] = weight
	} else {
		ett := make(map[int]int, 1)
		ett[sourceID] = weight

		g.edgesReverseIndex[targetID] = ett
	}

	return true
}

func (g *Graph) UpdateEdge(sourceID, targetID, weight int) bool {
	_, ok := g.nodes[sourceID]
	if !ok {
		return false
	}

	_, ok = g.nodes[targetID]
	if !ok {
		return false
	}

	ett, ok := g.edges[sourceID]
	if !ok {
		return false
	}
	w, ok := ett[targetID]
	if !ok {
		return false
	}
	if w == weight {
		// New weight same as the previous value. No update required.
		return true
	}
	ett[targetID] = weight

	g.edgesReverseIndex[targetID][sourceID] = weight

	return true
}

func (g *Graph) AddOrUpdateEdge(sourceID, targetID, weight int) bool {
	_, ok := g.nodes[sourceID]
	if !ok {
		return false
	}

	_, ok = g.nodes[targetID]
	if !ok {
		return false
	}

	ett, ok := g.edges[sourceID]
	if ok {
		var w int
		w, ok = ett[targetID]
		if ok {
			// Edge already present.

			if w != weight {
				// New weight not the same as the previous value. Update required.
				g.edgesReverseIndex[targetID][sourceID] = weight
			}

			return true
		}

		ett[targetID] = weight
	} else {
		ett := make(map[int]int, 1)
		ett[targetID] = weight

		g.edges[sourceID] = ett
	}

	ett, ok = g.edgesReverseIndex[targetID]
	if ok {
		ett[sourceID] = weight
	} else {
		ett := make(map[int]int, 1)
		ett[sourceID] = weight

		g.edgesReverseIndex[targetID] = ett
	}

	return true
}
