package askgo

// IsVeryDeep returns true iff depth > 1
func (t *Trv) IsVeryDeep() bool {
	if !t.isDeep {
		return false
	}

	for _, nestedTrv := range t.trvs {
		return nestedTrv.isDeep
	}

	return false
}

// hop from a current result to its neighbors, given a label
// optionally remember the path
func (t *Trv) hop(getIncomingNodes bool, label string, rememberPath bool) *Trv {

	// Deep calls
	if t.isDeep {
		for _, nestedTrv := range t.trvs {
			nestedTrv.hop(getIncomingNodes, label, rememberPath)
		}
		return t
	}

	newResult := map[string]Node{}
	newPath := map[string][]*Step{}

	var edges map[string]Edge
	var err error

	// Loop over all the nodes in current result
	for aNodeKey, aNode := range t.result {

		// Loop over all its relationships
		if getIncomingNodes {
			edges, err = aNode.In(t.graph, label)
		} else {
			edges, err = aNode.Out(t.graph, label)
		}
		if err != nil {
			t.AddError(err)
			continue
		}

		for _, edge := range edges {

			bNode, err := edge.Hop(t.graph, aNodeKey)
			if err != nil {
				t.AddError(err)
				continue
			}

			bNodeKey := bNode.Key()
			newResult[bNodeKey] = bNode
			if rememberPath {
				newPath[bNodeKey] = append(t.path[bNodeKey], NewStep(aNode, edge))
			} else {
				newPath[bNodeKey] = t.path[bNodeKey]
			}

		}

	}

	t.result = newResult
	t.path = newPath

	return t

}

// In moves the traversal to the incoming neighbors
func (t *Trv) In(label string, rememberPath bool) *Trv {
	return t.hop(true, label, rememberPath)
}

// Out moves the traversal to the outgoing neighbors
func (t *Trv) Out(label string, rememberPath bool) *Trv {
	return t.hop(false, label, rememberPath)
}
