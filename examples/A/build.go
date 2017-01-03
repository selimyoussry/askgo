package main

import "github.com/hippoai/graphgo"

func build() *graphgo.Graph {

	g := graphgo.NewEmptyGraph()

	g.MergeNode("company.OnePlus", map[string]interface{}{
		"name":     "OnePlus",
		"location": "america",
	})
	g.MergeNode("person.patrick", map[string]interface{}{
		"name": "patrick",
		"age":  20,
	})
	g.MergeNode("person.tim", map[string]interface{}{
		"name": "Tim",
		"age":  48,
	})
	g.MergeNode("person.clara", map[string]interface{}{
		"name": "Clara",
		"age":  18,
	})
	g.MergeNode("person.jimbo", map[string]interface{}{
		"name": "Jimbo",
		"age":  20,
	})

	g.MergeNode("person.john", map[string]interface{}{
		"name": "john",
		"age":  55,
	})
	g.MergeNode("person.elliott", map[string]interface{}{
		"name": "elliott",
		"age":  80,
	})

	// Now, add the edges (relationships) between these nodes
	g.MergeEdge(
		"patrick.worksin.OnePlus", "WORKS_IN",
		"person.patrick", "company.OnePlus",
		map[string]interface{}{},
	)
	g.MergeEdge(
		"john.worksin.OnePlus", "WORKS_IN",
		"person.john", "company.OnePlus",
		map[string]interface{}{},
	)
	g.MergeEdge(
		"tim.worksin.OnePlus", "WORKS_IN",
		"person.tim", "company.OnePlus",
		map[string]interface{}{},
	)
	g.MergeEdge(
		"clara.worksin.OnePlus", "WORKS_IN",
		"person.clara", "company.OnePlus",
		map[string]interface{}{},
	)
	g.MergeEdge(
		"jimbo.worksin.OnePlus", "WORKS_IN",
		"person.jimbo", "company.OnePlus",
		map[string]interface{}{},
	)
	g.MergeEdge(
		"john.is_father_of.patrick", "IS_FATHER_OF",
		"person.john", "person.patrick",
		map[string]interface{}{},
	)
	g.MergeEdge(
		"tim.is_father_of.clara", "IS_FATHER_OF",
		"person.tim", "person.clara",
		map[string]interface{}{},
	)
	g.MergeEdge(
		"tim.is_father_of.jimbo", "IS_FATHER_OF",
		"person.tim", "person.jimbo",
		map[string]interface{}{},
	)
	g.MergeEdge(
		"elliott.is_father_of.john", "IS_FATHER_OF",
		"person.elliott", "person.john",
		map[string]interface{}{},
	)

	return g
}
