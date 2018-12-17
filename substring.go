package main

import (
	"fmt"
	"net"
	"net/textproto"
	"time"
)

var(
	GATimeFormat = "197012011012201"
	GATimeFormat1 = "20060102150405"
)

func main() {
	//len 函数在获取带有汉字长度时也会出现问题
	//通过将字符串rune化后再进行截取操作
	var a = "历史会记住今天"
	var  r = []rune(a)
	fmt.Println(string(r[0]))
	fmt.Println(string(r))
	fmt.Println(len(a))
	fmt.Println(len(r))

	fmt.Println(substring(a, 1, len(r)))

	fmt.Println(	textproto.CanonicalMIMEHeaderKey("www-Authrecation"))
	fmt.Println(	textproto.CanonicalMIMEHeaderKey("Content-type"))
	fmt.Println(	textproto.CanonicalMIMEHeaderKey("content-type"))
	fmt.Println(time.Now().Format(GATimeFormat1))

	fmt.Println(net.InterfaceAddrs())

}

func substring(source string, start int, end int) string {

	var  r = []rune(source)
	length := len(r)

	if start < 0 || end > length || start > end{
		return ""
	}

	if start == 0 && end == length {
		return source
	}

	var substring = ""

	for i := start; i < length ; i++ {
		substring += string(r[i])
	}
	return substring
}