package askgo

import "github.com/hippoai/graphgo"

// Size returns the size of the result
func (t *Trv) Size() int {
	return len(t.result)
}

// Result returns the nodes we found
func (t *Trv) Result() map[string]graphgo.INode {
	return t.result
}

// Return returns the cache as a map
func (t *Trv) Return() map[string](map[string]interface{}) {
	return t.cache
}

// ReturnSlice returns the cache as a slice and not a map
func (t *Trv) ReturnSlice() []map[string]interface{} {
	r := []map[string]interface{}{}
	for _, value := range t.cache {
		r = append(r, value)
	}
	return r
}
