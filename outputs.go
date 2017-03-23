package askgo

import "github.com/hippoai/graphgo"

// Trvs are traversals for ar array of graphgo.Output
// Useful when we return a sorted response, and want to keep this order
// after the traversal
type OrderedTrvs struct {
	Trvs []*Trv `json:"trvs"`
}

// NewTrvs instanciates
func NewTrvs(outputs []*graphgo.Output, starts ...string) *OrderedTrvs {
	trvs := []*Trv{}
	for _, output := range outputs {
		trvs = append(trvs, NewTrv(output.Merge, starts...))
	}
	return &OrderedTrvs{
		Trvs: trvs,
	}
}

// Now copy the API for this Trvs
func (ts *OrderedTrvs) Size() int {
	ret := 0
	for _, trv := range ts.Trvs {
		ret = ret + trv.Size()
	}
	return ret
}

func (ts *OrderedTrvs) Result() []map[string]graphgo.INode {
	ret := []map[string]graphgo.INode{}
	for _, trv := range ts.Trvs {
		ret = append(ret, trv.Result())
	}
	return ret
}

// ReturnSlice
func (ts *OrderedTrvs) ReturnSlice() []map[string]interface{} {
	ret := []map[string]interface{}{}
	for _, trv := range ts.Trvs {
		for _, rs := range trv.ReturnSlice() {
			ret = append(ret, rs)
		}
	}
	return ret
}

func (ts *OrderedTrvs) ShallowSave(keys ...string) *OrderedTrvs {
	for _, trv := range ts.Trvs {
		trv.ShallowSave(keys...)
	}
	return ts
}

func (ts *OrderedTrvs) ShallowSaveF(f func(key string) (bool, string)) *OrderedTrvs {
	for _, trv := range ts.Trvs {
		trv.ShallowSaveF(f)
	}
	return ts
}

func (ts *OrderedTrvs) DeepSave(name string, simplify bool) *OrderedTrvs {
	for _, trv := range ts.Trvs {
		trv.DeepSave(name, simplify)
	}
	return ts
}

func (ts *OrderedTrvs) Copy() *OrderedTrvs {
	newTrvs := []*Trv{}
	for _, trv := range ts.Trvs {
		newTrvs = append(newTrvs, trv.Copy())
	}
	return &OrderedTrvs{
		Trvs: newTrvs,
	}
}

func (ts *OrderedTrvs) Deepen() *OrderedTrvs {
	for _, trv := range ts.Trvs {
		trv.Deepen()
	}
	return ts
}

func (ts *OrderedTrvs) Flatten() *OrderedTrvs {
	for _, trv := range ts.Trvs {
		trv.Flatten()
	}
	return ts
}

func (ts *OrderedTrvs) ShallowFilter(predicate func(graphgo.INode, []*Step) bool) *OrderedTrvs {
	for _, trv := range ts.Trvs {
		trv.ShallowFilter(predicate)
	}
	return ts
}

func (ts *OrderedTrvs) DeepFilter(keepQuery func(*Trv, []*Step) bool) *OrderedTrvs {
	for _, trv := range ts.Trvs {
		trv.DeepFilter(keepQuery)
	}
	return ts
}

func (ts *OrderedTrvs) In(label string, rememberPath bool) *OrderedTrvs {
	for _, trv := range ts.Trvs {
		trv.In(label, rememberPath)
	}
	return ts
}

func (ts *OrderedTrvs) Out(label string, rememberPath bool) *OrderedTrvs {
	for _, trv := range ts.Trvs {
		trv.Out(label, rememberPath)
	}
	return ts
}

func (ts *OrderedTrvs) InOut(inLabel string, inRememberPath bool, outLabel string, outRememberPath bool) *OrderedTrvs {
	for _, trv := range ts.Trvs {
		trv.InOut(inLabel, inRememberPath, outLabel, outRememberPath)
	}
	return ts
}

func (ts *OrderedTrvs) LogCache() *OrderedTrvs {
	for _, trv := range ts.Trvs {
		trv.LogCache()
	}
	return ts
}

func (ts *OrderedTrvs) LogResult() *OrderedTrvs {
	for _, trv := range ts.Trvs {
		trv.LogResult()
	}
	return ts
}

func (ts *OrderedTrvs) LogPath() *OrderedTrvs {
	for _, trv := range ts.Trvs {
		trv.LogPath()
	}
	return ts
}

func (ts *OrderedTrvs) Log(msgs ...string) *OrderedTrvs {
	for _, trv := range ts.Trvs {
		trv.Log(msgs...)
	}
	return ts
}
