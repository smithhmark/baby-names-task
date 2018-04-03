# baby-names-task
A solution the Baby Names interview question from "Cracking the Coding Interview"

## Problem
Some names are practically synonyms for others, creating a "name-family". Given data about child names, birth totals, and name equivalencies, calculate the numbers of births by name-family.

eg: Mike and Michael, Mic, Myke. 
eg: Chris, Christopher, Kris, Christofer

## Solution
I thought of this as a graph problem. The name equivalencies are adjacenies between nodes. The totals are the total births for a sub-graph.

## Tour
### Packages
#### main
The main package is just a runner over fixed example data at this point.
#### namegraph
The namegraph package implements the graph structure.

