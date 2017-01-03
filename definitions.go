// Package askgo provides traversal methods for directed property graphs
package askgo

import "github.com/hippoai/graphgo"

// Trv implements the graph traversal
type Trv struct {
	// Underlying graph
	graph graphgo.IGraph

	// Current map of nodes (indexed by their keys) returned by the traversal
	result map[string]graphgo.INode

	// Cache results as the traversal grows
	cache map[string](map[string]interface{})

	// Remember the traversal path
	path map[string][]*Step

	// Trvs are deeper traversals done to explore a specific aspect / filter
	trvs   map[string]*Trv
	isDeep bool

	Errors []error
}

// NewTrvWithPath instanciates
func NewTrvWithPath(graph graphgo.IGraph, path map[string][]*Step, starts ...string) *Trv {
	result := map[string]graphgo.INode{}
	errors := []error{}

	// Get the starting points in the initial result
	for _, start := range starts {
		node, err := graph.GetNode(start)
		if err != nil {
			errors = append(errors, err)
			continue
		}
		result[node.GetKey()] = node
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

// NewTrv instanciates with an empty path
func NewTrv(g graphgo.IGraph, starts ...string) *Trv {

	// Create the "empty" path
	path := map[string][]*Step{}
	for _, start := range starts {
		path[start] = []*Step{}
	}

	return NewTrvWithPath(g, path, starts...)
}

// Step remembers the node + edge needed to get from A to B
type Step struct {
	Node graphgo.INode
	Edge graphgo.IEdge
}

// NewStep instanciates
func NewStep(node graphgo.INode, edge graphgo.IEdge) *Step {
	return &Step{
		Node: node,
		Edge: edge,
	}
}
