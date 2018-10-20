package main

import(
	"fmt"
)

func main(){
	/*
	m:=make(map[int]string)
	m[1] = "OK"
	delete(m,1)
	a:=m[1]
	fmt.Println(a)*/

	//每一级的map都需要初始化
	/*var m1 map[int]map[int]string
	m1 = make(map[int]map[int]string)
//	m1[1] = make(map[int]string)
	a, ok := m1[2][1]
	if !ok {
		m1[2] = make(map[int]string)
	}
	m1[2][1] = "okMap"
	b := m1[2][1]
	fmt.Println(b, ok)*/

	m1 := map[int]string{1: "a", 2: "b", 3: "c"}
	fmt.Println(m1)
	m2 := make(map[string]int)
	for k, v := range m1{
		m2[v] = k
	}
	fmt.Println(m2)
}
