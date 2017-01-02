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
