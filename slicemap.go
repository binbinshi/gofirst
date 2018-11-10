package main

import (
	"fmt"
	"strings"
)
type statusFlag uint16

const(
	a statusFlag = iota
	b
	c
	d
)

type fieldType byte
const(
	e fieldType = iota
	f
	h
	g
)


type clientFlag uint32

const (
	j clientFlag = 1<<iota
	k
	l
)


const  (
	n fieldType = iota + 0xf5
	m
)
func main() {
	str := "h e l l o g o"
	strMap := countStr(str)
	fmt.Println(strMap)
	removeElem("h", strMap)
	fmt.Println(1<<24 - 1)
	fmt.Println(1<<2)
	fmt.Println(1<<3)
	fmt.Println(8>>2)
	fmt.Println(k|l|j)
	fmt.Println(k+l+j)
	fmt.Println(a,b,c,d,e,f,h,g,j,k,l,n,m)
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
