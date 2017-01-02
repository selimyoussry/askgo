package askgo

// ShallowFilter the current nodes based on their properties and the path
func (t *Trv) ShallowFilter(predicate func(Node, []*Step) bool) *Trv {

	// Deep Calls
	if t.isDeep {
		for _, nestedTrv := range t.trvs {
			nestedTrv.ShallowFilter(predicate)
		}
		return t
	}

	newResult := map[string]Node{}

	// Loop over all the nodes in the current result
	for nodeKey, node := range t.result {

		if predicate(node, t.path[nodeKey]) {
			newResult[nodeKey] = node
		}

	}

	t.result = newResult
	return t

}
