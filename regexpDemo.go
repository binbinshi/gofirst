package main

import (
	"fmt"
	"regexp"
)

func main() {
	fmt.Println(IsIp("0.199.1.179"))
}

func IsIp(ip string) bool{
	pattern := "^((25[0-5]|2[0-4]\\d|[01]?\\d\\d?)\\.){3}(25[0-5]|2[0-4]\\d|[01]?\\d\\d?)$"
	if mr, _ := regexp.MatchString(pattern , ip) ; !mr {
		return false
	}
	return true
}