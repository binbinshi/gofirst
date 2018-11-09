package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "h e l l o g o"
	strMap := countStr(str)
	fmt.Println(strMap)
	removeElem("h", strMap)

}

func countStr(s string) map[string]int {
	fields := strings.Fields(s)
	strMap := make(map[string]int)
	for _, v := range fields {
		if _, ok := strMap[v]; ok {
			strMap[v] += 1
		} else {
			strMap[v] = 1
		}
	}
	return strMap
}

func removeElem(s string, m map[string]int) {
	if _, ok := m[s]; ok {
		delete(m, s)
		fmt.Printf("map remove elem key %s\n", s)
	}
}
