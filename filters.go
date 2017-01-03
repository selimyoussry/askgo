package askgo

// HasResult is a built-in filter
// Returns true if and only if the current traversal is non-empty
func HasResult(trv *Trv, path []*Step) bool {
	return trv.Size() > 0
}
