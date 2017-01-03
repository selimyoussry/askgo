# Ask [questions to your directed property graphs in] Go

If you have some data stored in a directed property graph, [like this implementation](https://github.com/hippoai/graphgo.git), you can query it using this AskGo. This is a graph traversal library, very much inspired by [Gremlin](http://tinkerpop.apache.org/).

**Requirements** The graph, nodes and edges needs to implement the interfaces described in [Graphgo](https://github.com/hippoai/graphgo.git) in order to navigate it.


## Install

`go get -u github.com/hippoai/askgo.git`

## Fundamental concepts

When asking a question to your graph, you must identify the three pillars:
1. The central node(s) of the query
2. Where to start to get to the central node(s)
3. What information do we want to know for these nodes?

For instance, let's say we have a graph consisting of:
* a company `OnePlus`, with the unique ID `company.OnePlus`
* a few employees with unknown IDs, who have a name and age, and are related to the company by the relationship `WORKS_FOR`
* the `IS_FATHER_OF` relationship. Some employees have their father working with them too.

and we want to answer the question `Which employees of OnePlus have their father working there? We want their name, age, and their father's age.`. We identify:
1. The central nodes are the employees
2. We must start from the company since we don't know exactly the employees' IDs
3. We need to extract the employees' name and age (from their node properties), but also information about their father (from the {node at the end of the `IS_FATHER_OF`'s relationship}'s properties).

Let's try to answer this question with an `askgo` traversal, assuming the graph is in the object `g`.

```go
// Create a traversal on this graph
traversal := askgo.NewTrv(g, "company.OnePlus")

// Query
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
```

and if all went well you get the result

```javascript
{}
```

You can have a look at the [examples folder](https://github.com/hippoai/askgo/tree/master/examples/) for a more in-depth tutorial.

## API

You can check the API by using GoDoc or looking deeper in the code for now. I'll wait for the code to be a bit more tested to publish on [GoDoc](https://godoc.org).
