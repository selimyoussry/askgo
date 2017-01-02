package askgo

// Graph needs to be able to find a node and an edge given their key
type Graph interface {
	GetNode(key string) (Node, error)
	GetEdge(key string) (Edge, error)
}

// Edge needs to be able to access its properties, start and end node, and label
// it has a unique key
type Edge interface {
	Get(key string) (interface{}, error)
	Start(graph Graph) (Node, error)
	End(graph Graph) (Node, error)
	Label() string
	Key() string
}

// Node needs to be able to access its properties,
// ingoing and outgoing edges
// it has a unique key
type Node interface {
	Get(key string) (interface{}, error)
	In(graph Graph, label string) (map[string]Edge, error)
	End(graph Graph, label string) (map[string]Edge, error)
	Key() string
}
