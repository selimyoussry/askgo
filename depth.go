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

// Deepen freezes the current result
// and allows exploration of each of its nodes independently
func (t *Trv) Deepen() *Trv {

	// Deep Calls
	if t.isDeep {
		for _, nestedTrv := range t.trvs {
			nestedTrv.Deepen()
		}
		return t
	}

	trvs := map[string]*Trv{}
	for _, node := range t.result {
		// Use the node key as a query key
		nodeKey := node.GetKey()
		trvs[nodeKey] = NewTrvWithPath(t.graph, t.path, nodeKey)
	}

	t.trvs = trvs
	t.isDeep = true
	return t
}

// Flatten flattens a traversal to the lower level
func (t *Trv) Flatten() *Trv {

	// Nothing to flatten
	if !t.isDeep {
		return t
	}

	// If it's actually too deep, we keep going
	if t.IsVeryDeep() {
		for _, nestedTrv := range t.trvs {
			nestedTrv.Flatten()
		}
		return t
	}

	t.trvs = map[string]*Trv{}
	t.isDeep = false
	return t

}
