package main

import (
	"encoding/json"
	"fmt"

	"github.com/hippoai/askgo"
)

func main() {

	g := build()
	prettyPrint(g)

	traversal := askgo.NewTrv(g, "company.OnePlus")

	prettyPrint(traversal.Result())

	worksForOnePlus := func(t *askgo.Trv, path []*askgo.Step) bool {
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
		DeepFilter(askgo.HasResult).     // This function, described below, returns true iff the current employee who have a father. It will filter out the employees not fulfilling this requirement at the lower level, i.e. 1.
		Deepen().                        // Now for every father, we need to make sure they also work at OnePlus. We need to explore their own relationship
		Out("WORKS_IN", false).          // Get the companies they work for (Nota Bene: we are at level 2 here)
		DeepFilter(worksForOnePlus).     // Filter at level 1, the fathers working for OnePlus only
		Flatten().                       // Go back to level 1
		DeepFilter(askgo.HasResult).     // Because of the level 2 filtering above, there might be empty traversals at level 1, corresponding to fathers not working at OnePlus. We need to filter them out. This is very common, so askgo provides a built-in function for this.
		ShallowSave("name::fatherName"). // Save in the top level the father name, call it fatherName
		ShallowSave("inexistantField").  // This field does not exist, this will be reported in the errors
		DeepSave("father").              // Saves the fathers' (level 2) cache in the lower level cache, under the name "father", for each employee
		Flatten().                       // We are done at this level, so go back to level 1
		ShallowSave("name").             // Now we save the employee's name (no alias this time, so no need for the a::b pattern)
		Return()                         // Finally return the cache

	prettyPrint(result)
	prettyPrint(traversal.Errors)

}

func prettyPrint(x interface{}) {
	b, _ := json.MarshalIndent(x, "", "  ")
	fmt.Println(string(b))
}
