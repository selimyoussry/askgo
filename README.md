# Ask [questions to your directed property graphs in] Go

If some data is stored in a directed property graph - [like this implementation](https://github.com/hippoai/graphgo.git) -, you can query it using this AskGo.


## Install

`go get github.com/hippoai/askgo.git`

## Requirements

We will run queries on top of a graph. Its purpose is to store data. In order for the query engine (the AskGo library) to answer your question, the graph needs to implement the `askgo.Graph` interface. I.e. it needs the following methods:

```go
// GetEdge returns a pointer to an Edge, given its unique key, and an error if it could not be found
GetEdge(key string) (askgo.Edge, error)

// GetNode returns a pointer to a Node, given its unique key, and an error if it could not be found
GetNode(key string) (askgo.Node, error)
```

where `askgo.Edge` implements

```go
// Get returns a property, given its key, and an error if it could not be found
Get(key string) (interface{}, error)

// Start returns the start node
Start(graph askgo.Graph) (askgo.Node, error)

// End returns the end node
End(graph askgo.Graph) (askgo.Node, error)

// Hop returns either the start or end node
Hop(graph askgo.Graph, key string) (askgo.Node, error)

// Label returns the edge label
Label() string

// Key returns the key
Key() string
```

and `askgo.Node` implements the following interface

```go
// Get returns a property, given its key, and an error if it could not be found
Get(key string) (interface{}, error)

// In returns a map of outgoing edges with the given label, indexed by their key
In(g askgo.Graph, label string) (map[string]askgo.Edge, error)

// Out returns a map of outgoing edges with the given label, indexed by their key
Out(g askgo.Graph, label string) (map[string]askgo.Edge, error)

// Key returns the key
Key() string
```
