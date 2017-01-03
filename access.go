package askgo

import "github.com/hippoai/graphgo"

func (t *Trv) Size() int {
	return len(t.result)
}

func (t *Trv) Result() map[string]graphgo.INode {
	return t.result
}

func (t *Trv) Return() map[string](map[string]interface{}) {
	return t.cache
}
