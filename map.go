package main

import(
	"fmt"
	"reflect"
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

	basicMap := make(map[interface{}]interface{})
	fmt.Println(basicMap)
	basicMap = map[interface{}]interface{}{
		1: "Alice",
		2: "Lucas",
		3: "Ronald",
		4: "Messi",
		5: "Paul",
	}
	slice1 := []float64{1.1,7.8,9.0,3.4}
	fmt.Println(index(slice1, 9.0))
	fmt.Print()

	slice3 := []interface{}{0,3,4,5,"abc",6}
	fmt.Println(index(slice3, "abc"))
	mapPrint(basicMap)
}

func mapPrint(m map[interface{}]interface{}){
	//遍历map并打印key-value


}

func index (slice interface{}, v interface{}) int {
	if slice := reflect.ValueOf(slice); slice.Kind() == reflect.Slice {
		for i := 0; i < slice.Len(); i++  {
			if reflect.DeepEqual(v, slice.Index(i).Interface()) {
				return i
			}
		}
	}
	return -1
}



/*for i := 0; i < 5; i++  {
	mapPrint(basicMap)
	fmt.Println()
}*/
