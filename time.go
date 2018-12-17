package main

import (
	"crypto/md5"
	b64 "encoding/base64"
	"encoding/hex"
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

var nonce1 = ""
func main() {

//	t := "2006:01:02:03:04:05"
//	time := time.Now()
//	timeStr := fmt.Sprintf(time.Format(t)) strconv.Itoa(rand.Intn(100000))
//	data := "2006:01:02:03:04:05" + "8888"
//	h := md5.New()
//	h.Write([]byte(data))
//	dataMd5 := h.Sum(nil)
	fmt.Println(time.Now().Second())
	fmt.Println(time.Now().Hour())
	fmt.Println(time.Now().Nanosecond())
	fmt.Println(time.Now().Unix())
	fmt.Println(time.Now().UnixNano())
	fmt.Println(time.Now().UnixNano()/1e6)
	fmt.Println(time.Now().UnixNano()/1e9)
	time := time.Now().UnixNano()/1e6
	timeStr := strconv.Itoa(int(time))
	fmt.Println(timeStr)


	a := "ssssssssssssssssssssssssssssssss"
	nonce1 = b64.URLEncoding.EncodeToString([]byte(a))
	fmt.Println(nonce1)
	d, _ := b64.URLEncoding.DecodeString("MTU0MjY5NzMzNDIxODo0NjU4OGE2Nzk3NDk0NDVlNmNjMmU5MDM3OGEyOGQyYQ==")
	fmt.Println(string(d))

	c, _:= b64.StdEncoding.DecodeString("MTU0MjY5NzMzNDIxODo0NjU4OGE2Nzk3NDk0NDVlNmNjMmU5MDM3OGEyOGQyYQ==")
	fmt.Println(string(c))
	e, _ := b64.URLEncoding.DecodeString("MmI0ZjRhOTEwZDI1YmI5OGY1ODZkODMxYTY5MDMxMzk=")
	fmt.Println("ddddd: ",string(e))
	time1 := time
	timeStr1 := strconv.Itoa(int(time1))
	data := timeStr1 + strconv.Itoa(rand.Intn(100000))
	h := md5.New()
	h.Write([]byte(data))
	dataMd5 := h.Sum(nil)
	nonce1 = b64.URLEncoding.EncodeToString([]byte(timeStr1 + ":" + hex.EncodeToString(dataMd5)))
	fmt.Println(nonce1)

}
