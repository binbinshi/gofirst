package main_test

import (
	"encoding/json"
	"fmt"
	"testing"
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

func TestStruct2Json(t *testing.T) {
	var student Student
	student.Age = 18
	student.Name = "binbin"
	student.StNo = "0001"
	str := Struct2Json(student)
	fmt.Println(str)
	t.Log(str)
}
