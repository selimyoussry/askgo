package askgo

import (
	"github.com/hippoai/goutil"
	"github.com/hippoai/graphgo"
)

// HasResult is a built-in filter
// Returns true if and only if the current traversal is non-empty
func HasResult(trv *Trv, path []*Step) bool {
	return trv.Size() > 0
}

// IsInValuesString returns a shallow filter
// that will return true iff value is in given values
func IsInValuesString(key string, values ...string) func(graphgo.INode, []*Step) bool {
	return func(iNode graphgo.INode, path []*Step) bool {
		vItf, err := iNode.Get(key)
		if err != nil {
			return false
		}

		v, ok := vItf.(string)
		if !ok {
			return false
		}

		return goutil.IsIn(v, values...)
	}
}
