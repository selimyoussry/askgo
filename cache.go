package askgo

import "strings"

// renameKey allows for key renaming, using the magic KEY_SEPARATOR "::"
// Writing a::b will extract the property "a" and save it under name "b" in the cache
func renameKey(key string) (string, string) {
	splitted := strings.Split(key, KEY_SEPARATOR)
	if len(splitted) == 1 {
		return splitted[0], splitted[0]
	}

	return splitted[0], splitted[1]
}

// ShallowSave saves a bunch of keys in the traversal cache
func (t *Trv) ShallowSave(keys ...string) *Trv {

	// Deep Calls
	if t.isDeep {
		for _, nestedTrv := range t.trvs {
			nestedTrv.ShallowSave(keys...)
		}
		return t
	}

	// Loop over every node in the result
	for nodeKey, node := range t.result {
		_, exists := t.cache[nodeKey]
		if !exists {
			t.cache[nodeKey] = map[string]interface{}{}
		}

		// Loop over every key we care about
		for _, key := range keys {
			oldKey, newKey := renameKey(key)
			value, err := node.Get(oldKey)
			if err != nil {
				t.AddError(err)
				continue
			}
			t.cache[nodeKey][newKey] = value
		}
	}
	return t
}

func (t *Trv) ShallowSaveF(f func(key string) (bool, string)) *Trv {

	// Deep Calls
	if t.isDeep {
		for _, nestedTrv := range t.trvs {
			nestedTrv.ShallowSaveF(f)
		}
		return t
	}

	// Loop over every node in the result
	for nodeKey, node := range t.result {
		_, exists := t.cache[nodeKey]
		if !exists {
			t.cache[nodeKey] = map[string]interface{}{}
		}

		// Loop over every key we care about
		for key, value := range node.GetProps() {
			keep, newKey := f(key)
			if keep {
				t.cache[nodeKey][newKey] = value
			}
		}
	}
	return t

}

// DeepSave the outermost cache the the level just below it
func (t *Trv) DeepSave(name string) *Trv {

	// Nothing to save, it is not deep
	if !t.isDeep {
		return t
	}

	// If it's actually too deep, we keep going
	if t.IsVeryDeep() {
		for _, nestedTrv := range t.trvs {
			nestedTrv.DeepSave(name)
		}
		return t
	}

	// Otherwise, this is the level before the lowest
	// We can flatten the cache
	for nodeKey, nestedTrv := range t.trvs {
		_, exists := t.cache[nodeKey]
		if !exists {
			t.cache[nodeKey] = map[string]interface{}{}
		}

		t.cache[nodeKey][name] = nestedTrv.cache
	}

	return t

}
