package main

import (
	"encoding/base64"
	"fmt"
	"strings"
)

func main() {
	userId := "admin"
	password := "123456a"
	basic := createBasicAuth(userId, password)
	fmt.Println(basic)
	fmt.Println(checkBasicAuth(basic))
}

func createBasicAuth(userId, password string) string{
	baseStr := make([]string, 0)
	baseStr = append(baseStr, userId)
	baseStr = append(baseStr, password)
	baseAuth := strings.Join(baseStr, ":")
	fmt.Println(baseAuth)
	return base64.StdEncoding.EncodeToString([]byte(baseAuth))
}


func checkBasicAuth(auth string) bool{
	b, err := base64.StdEncoding.DecodeString(auth)
	if err != nil {
		fmt.Println(err)
		return false
	}
	if string(b) != "admin:123456a" {
		return false
	}
	return true
}