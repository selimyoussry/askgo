package askgo

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"testing"

	"github.com/hippoai/goerr"
	"github.com/hippoai/goutil"
	"github.com/hippoai/graphgo"
)

type Level0 map[string]Level1
type Level1 struct {
	Father Level2 `json:"father"`
	Name   string `json:"name"`
}
type Level2 struct {
	FatherName string `json:"fatherName"`
}

func TestTrv(t *testing.T) {

	g, err := build()
	if err != nil {
		t.Errorf(goutil.Pretty(err))
	}

	traversal := NewTrv(g, "company.OnePlus")

	worksForOnePlus := func(t *Trv, path []*Step) bool {
		for _, node := range t.Result() {
			companyName, err := node.Get("name")
			if (err != nil) || (companyName != "OnePlus") {
				continue
			}
			return true
		}

		return false
	}

	result := traversal.
		In("WORKS_IN", false).           // We find the employee nodes, no need to save how we got there, hence the "false" argument
		Deepen().                        // Now that we are at the employees level, we need to explore their IS_FATHER_OF relationship and discard the ones with no father working here. This is called "Deepen" because we freeze the first query at level 1, and go to level 2 just for father exploration
		In("IS_FATHER_OF", false).       // In the deep query now, move to the fathers, for each employee
		DeepFilter(HasResult).           // This function, described below, returns true iff the current employee who have a father. It will filter out the employees not fulfilling this requirement at the lower level, i.e. 1.
		Deepen().                        // Now for every father, we need to make sure they also work at OnePlus. We need to explore their own relationship
		Out("WORKS_IN", false).          // Get the companies they work for (Nota Bene: we are at level 2 here)
		DeepFilter(worksForOnePlus).     // Filter at level 1, the fathers working for OnePlus only
		Flatten().                       // Go back to level 1
		DeepFilter(HasResult).           // Because of the level 2 filtering above, there might be empty traversals at level 1, corresponding to fathers not working at OnePlus. We need to filter them out. This is very common, so askgo provides a built-in function for this.
		ShallowSave("name::fatherName"). // Save in the top level the father name, call it fatherName
		ShallowSave("inexistantField").  // This field does not exist, this will be reported in the errors
		DeepSave("father", true).        // Saves the fathers' (level 2) cache in the lower level cache, under the name "father", for each employee
		Flatten().           // We are done at this level, so go back to level 1
		ShallowSave("name"). // Now we save the employee's name (no alias this time, so no need for the a::b pattern)

	// Make sure we can reformat the response
	var trvResponse Level0
	err = goutil.JsonRestruct(result.Return(), &trvResponse)
	if err != nil {
		t.Errorf(goutil.Pretty(err))
	}

	// Now make sure the response has the right answer
	clara, ok := trvResponse["person.clara"]
	if !ok {
		t.Fatalf(goutil.Pretty(goerr.NewS("ERR_RESPONSE")))
	}

	if clara.Father.FatherName != "Tim" {
		t.Fatalf(goutil.Pretty(goerr.NewS("ERR_RESPONSE")))
	}

	// Length of the response
	if len(trvResponse) != 3 {
		t.Fatalf(goutil.Pretty(goerr.NewS("ERR_RESPONSE")))
	}

	// Contains a name
	if clara.Name != "Clara" {
		t.Fatalf(goutil.Pretty(goerr.NewS("ERR_RESPONSE")))
	}

}

func loadJson(fileName string) (*graphgo.Output, error) {
	f, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}

	b, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}

	var o graphgo.Output
	err = json.Unmarshal(b, &o)
	if err != nil {
		return nil, err
	}

	return &o, nil

}

// build a graph
func build() (*graphgo.Graph, error) {

	g := graphgo.NewEmptyGraph()
	o, err := loadJson("./data/data1.json")
	if err != nil {
		return nil, err
	}
	g.Merge(o)

	return g, nil
}
