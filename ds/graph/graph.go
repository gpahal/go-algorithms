package graph

type (
	ID     int
	Value  int
	Weight int
)

type Node struct {
	ID    ID
	Value Value
}

type Edge struct {
	SourceID ID
	TargetID ID
	Weight   Weight
}

type Graph struct {
	nodes             map[ID]*Node
	edges             map[ID]map[ID]Weight
	edgesReverseIndex map[ID]map[ID]Weight

	currID ID
}

func New() *Graph {
	return &Graph{
		nodes:             make(map[ID]*Node),
		edges:             make(map[ID]map[ID]Weight),
		edgesReverseIndex: make(map[ID]map[ID]Weight),
		currID:            1,
	}
}

func (g *Graph) NodeCount() int {
	return len(g.nodes)
}

func (g *Graph) Node(id ID) *Node {
	node, ok := g.nodes[id]
	if ok {
		return node
	}

	return nil
}

func (g *Graph) HasNode(id ID) bool {
	return g.Node(id) != nil
}

func (g *Graph) AddNode(value Value) {
	g.nodes[g.currID] = &Node{
		ID:    g.currID,
		Value: value,
	}

	g.currID++
}

func (g *Graph) UpdateNode(id ID, value Value) bool {
	node, ok := g.nodes[id]
	if ok {
		node.Value = value
		return true
	}

	return false
}

func (g *Graph) DeleteNode(id ID) bool {
	_, ok := g.nodes[id]
	if ok {
		delete(g.nodes, id)
		return true
	}

	return false
}

func (g *Graph) Edge(sourceID, targetID ID) *Edge {
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

func (g *Graph) HasEdge(sourceID, targetID ID) bool {
	return g.Edge(sourceID, targetID) != nil
}

func (g *Graph) AddEdge(sourceID, targetID ID, weight Weight) bool {
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
		ett := make(map[ID]Weight, 1)
		ett[targetID] = weight

		g.edges[sourceID] = ett
	}

	ett, ok = g.edgesReverseIndex[targetID]
	if ok {
		ett[sourceID] = weight
	} else {
		ett := make(map[ID]Weight, 1)
		ett[sourceID] = weight

		g.edgesReverseIndex[targetID] = ett
	}

	return true
}

func (g *Graph) UpdateEdge(sourceID, targetID ID, weight Weight) bool {
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

func (g *Graph) AddOrUpdateEdge(sourceID, targetID ID, weight Weight) bool {
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
		var w Weight
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
		ett := make(map[ID]Weight, 1)
		ett[targetID] = weight

		g.edges[sourceID] = ett
	}

	ett, ok = g.edgesReverseIndex[targetID]
	if ok {
		ett[sourceID] = weight
	} else {
		ett := make(map[ID]Weight, 1)
		ett[sourceID] = weight

		g.edgesReverseIndex[targetID] = ett
	}

	return true
}
