package askgo

import "github.com/hippoai/graphgo"

func (trv *Trv) Copy() *Trv {

	newTrv := &Trv{
		graph:  trv.graph,
		result: map[string]graphgo.INode{},
		cache:  map[string](map[string]interface{}){},
		path:   map[string][]*Step{},
		trvs:   map[string]*Trv{},
		isDeep: trv.isDeep,
		Errors: []error{},
	}

	// Copy result
	for key, value := range trv.result {
		newTrv.result[key] = value
	}

	// Copy cache
	for key, value := range trv.cache {
		newTrv.cache[key] = value
	}

	// Copy path

	return newTrv

}
