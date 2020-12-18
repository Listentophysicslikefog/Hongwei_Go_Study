package main

import (
	"fmt"
	"sort"
)


func main() {

	var test map[string]bool
	var sortstest []string
	test = make(map[string]bool, 0)

	test["vdb"] = true
	test["vdf"] = true
	test["vdd"] = true
	test["vdh"] = true
	test["vdc"] = true
	test["vdk"] = true
	for i, v := range test {
		fmt.Println(i, ":    ", v)
		sortstest = append(sortstest, i)

	}

	sort.Strings(sortstest)
	for _, v := range sortstest {
		fmt.Println(v, ":     ", test[v])
	}

}
