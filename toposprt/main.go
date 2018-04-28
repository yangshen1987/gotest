package toposprt

import (
	"sort"
	"fmt"
)

var Prereqs = map[string][]string{
	"algorithms":{"data struchtures"},
	"calculus":  {"linear algebra"},
	"comilers":{
		"data structures",
		"formal languages",
		"compuser organization",
	},
	"data structures": {"discrete math"},
	"databases":       {"data struchtures"},
	"discrere math" :  {"intro to programming"},
	"formal languages" : {"discrete math"},
	"networks"      : {"operating systems"},
	"operating systems" : {"data structures", "computer organization"},
	"programming languages" : {"data structures", "computer organization"},
}

func main()  {
	for i, coure := range Topsort(Prereqs){
		fmt.Printf("%d:\t%s\n", i+1,coure)
	}
}
func Topsort(m map[string][]string) []string  {
	var order []string
	seen := make(map[string]bool)
	var visitAll  func(items []string)
	visitAll = func(items []string) {
		for _,item := range items{
			if !seen[item]{
				seen[item] = true
				visitAll(m[item])
				order = append(order, item)
			}
		}
	}
	var keys  []string
	for key := range m{
		keys = append(keys, key)
	}
	sort.Strings(keys)
	visitAll(keys)
	return order
}
