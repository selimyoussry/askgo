package askgo

import (
	"encoding/json"
	"fmt"
	"strings"
)

// LogCache
func (t *Trv) LogCache() *Trv {
	b, _ := json.MarshalIndent(t.cache, "", "  ")
	fmt.Println(string(b))
	return t
}

// LogResult
func (t *Trv) LogResult() *Trv {
	out := t.deepLog()
	b, _ := json.MarshalIndent(out, "", "  ")
	fmt.Println(string(b))
	return t
}

// deepLog is called on a deep the first time
func (t *Trv) deepLog() map[string]interface{} {

	if !t.isDeep {
		result := map[string]interface{}{}
		for nodeKey, _ := range t.result {
			result[nodeKey] = ""
		}
		return result
	}

	// Otherwise call deepLog depth first
	result := map[string]interface{}{}
	for nodeKey, nestedTrv := range t.trvs {
		oneResult := nestedTrv.deepLog()
		result[nodeKey] = oneResult
	}
	return result

}

// LogPath
func (t *Trv) LogPath() *Trv {

	for nodeKey, onePath := range t.path {

		steps := []string{}
		for _, step := range onePath {
			steps = append(steps, fmt.Sprintf("%s (%s)", step.Node.GetKey(), step.Edge.GetKey()))
		}

		fmt.Printf("%s: %s \n", nodeKey, strings.Join(steps, " > "))
	}

	return t
}

// Log logs everything
func (t *Trv) Log(msgs ...string) *Trv {
	for _, msg := range msgs {
		fmt.Println(msg)
	}

	fmt.Println("> Cache")
	t.LogCache()

	fmt.Println("> Result")
	t.LogResult()

	fmt.Println("> Path")
	t.LogPath()

	fmt.Println("--- --- ---")

	return t
}
