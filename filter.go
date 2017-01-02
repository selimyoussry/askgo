package askgo

import "github.com/hippoai/graphgo"

// ShallowFilter the current nodes based on their properties and the path
func (t *Trv) ShallowFilter(predicate func(graphgo.INode, []*Step) bool) *Trv {

	// Deep Calls
	if t.isDeep {
		for _, nestedTrv := range t.trvs {
			nestedTrv.ShallowFilter(predicate)
		}
		return t
	}

	newResult := map[string]graphgo.INode{}

	// Loop over all the nodes in the current result
	for nodeKey, node := range t.result {

		if predicate(node, t.path[nodeKey]) {
			newResult[nodeKey] = node
		}

	}

	t.result = newResult
	return t

}

// DeepFilter filters the N-1 result based on the outermost (N) level
func (t *Trv) DeepFilter(keepQuery func(*Trv, []*Step) bool) *Trv {

	// Nothing to filter
	if !t.isDeep {
		return t
	}

	// If it's actually too deep, we keep going
	if t.IsVeryDeep() {
		for _, nestedTrv := range t.trvs {
			nestedTrv.DeepFilter(keepQuery)
		}
		return t
	}

	// Otherwise, this is the level before the lowest
	nodesToDiscard := []string{}
	for nodeKey, nestedTrv := range t.trvs {

		// if we need to filter this
		if !keepQuery(nestedTrv, t.path[nodeKey]) {
			nodesToDiscard = append(nodesToDiscard, nodeKey)
		}

	}

	// Delete the nodes that have been filtered
	for _, nodeKey := range nodesToDiscard {
		delete(t.result, nodeKey)
		delete(t.trvs, nodeKey)
	}

	return t
}
