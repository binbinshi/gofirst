package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	Name string
	Age  int
	StNo string
}

func Struct2Json(student Student) string{
	str, err := json.Marshal(student)
	if err != nil {
		fmt.Println("convert error ")
	}
	return string(str)
}

