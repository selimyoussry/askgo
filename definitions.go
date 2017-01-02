// Package askgo provides traversal methods for directed property graphs
package askgo

// Trv implements the graph traversal
type Trv struct {
	// Underlying graph
	graph Graph

	// Current map of nodes (indexed by their keys) returned by the traversal
	result map[string]Node

	// Cache results as the traversal grows
	cache map[string](map[string]interface{})

	// Remember the traversal path
	path map[string][]*Step

	// Trvs are deeper traversals done to explore a specific aspect / filter
	trvs   map[string]*Trv
	isDeep bool

	errors []error
}

// NewQueryWithPath instanciates
func NewQueryWithPath(graph Graph, path map[string][]*Step, starts ...string) *Trv {
	result := map[string]Node{}
	errors := []error{}

	// Get the starting points in the initial result
	for _, start := range starts {
		node, err := graph.GetNode(start)
		if err != nil {
			errors = append(errors, err)
			continue
		}
		result[node.Key()] = node
	}

	return &Trv{
		graph:  graph,
		result: result,
		cache:  map[string](map[string]interface{}){},
		path:   path,
		trvs:   map[string]*Trv{},
		isDeep: false,
	}
}

// NewQuery instanciates with an empty path
func NewQuery(g Graph, starts ...string) *Trv {

	// Create the "empty" path
	path := map[string][]*Step{}
	for _, start := range starts {
		path[start] = []*Step{}
	}

	return NewQueryWithPath(g, path, starts...)
}

// Step remembers the node + edge needed to get from A to B
type Step struct {
	Node Node
	Edge Edge
}

// NewStep instanciates
func NewStep(node Node, edge Edge) *Step {
	return &Step{
		Node: node,
		Edge: edge,
	}
}
