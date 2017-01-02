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
		nodeKey := node.Key()
		trvs[nodeKey] = NewTrvWithPath(t.graph, t.path, nodeKey)
	}

	t.trvs = trvs
	t.isDeep = true
	return t
}
