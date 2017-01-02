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

// Save a bunch of keys in the traversal cache
func (t *Trv) Save(keys ...string) *Trv {

	// Deep Calls
	if t.isDeep {
		for _, nestedTrv := range t.trvs {
			nestedTrv.Save(keys...)
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
