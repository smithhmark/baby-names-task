package main

import "fmt"

//const nameCounts = map[string]int{
var nameCounts = map[string]int{
	"John": 15,
	"Jon": 12,
	"Chris": 13,
	"Kris": 4,
	"Christopher": 19,
}

func printRawCounts(cnts map[string]int, showHeader bool) {
	if showHeader {
		fmt.Printf("\t%s\t%s\n", "Name", "Count")
	}

	for nm, cnt := range cnts {
		fmt.Printf("\t%s\t%v\n", nm, cnt)
	}
}

func main() {
	fmt.Printf("Going to demonstrate the solution here.\n")
	printRawCounts(nameCounts, true)
}

