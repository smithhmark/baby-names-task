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

## Next Steps
 * It would be interesting to scrape name data sources to push this from pure algorithm exercise to more functional tool
   * https://www.ssa.gov/oact/babynames/limits.html for name data
   * http://electricka.com/etaf/muses/languagearts/words/people_first_name_equivalents/people_name_equivalents_table/people_name_equivalents_list.asp?pagesize=500 would be an interesting source of name equivalencies
   
