package main

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

var(
	b64 = base64.URLEncoding
	userName = "root"
	password = "123456a"
	realm = "pmplatform"
	authMethod = "auth"
	qop = authMethod
	algorithm = "MD5"
	nonce2 = "MTU0NDQzNDA3NjY1OTplNTE4YzhiZDE4NzdkYWVkN2M3NWI5N2YzNTk1ZjJiNA=="
	//客户端随机数
	nc = "00000001"
	httpMethod="POST"
	uri = "/VIID/System/Register"
	cnonce = ""
)
func main() {
	fmt.Println("getHa1 :"+ getHa1())
	fmt.Println("getHa2 :"+ getHa2(httpMethod, uri))
	cnonce = createCnonce()
	fmt.Println("cnonce :" + cnonce)
	response := createResponse()
	fmt.Println("response :" + response )
}

func createResponse() string{
	data := getHa1() + ":" + nonce2 + ":" + nc + ":" + cnonce + ":" + qop + ":" + getHa2(httpMethod, uri)
	return hexmd5(data)
}

func getHa1() string {
	data := userName + ":" + realm + ":"+ password
	return hexmd5(data)
}

func getHa2(httpMethod string, uri string) string {
	data := httpMethod + ":" + uri
	return hexmd5(data)
}

func hexmd5(data string)string{
	h := md5.New()
	h.Write([]byte(data))
	return hex.EncodeToString(h.Sum(nil))
}

func createCnonce () string{
	time := time.Now().UnixNano()/1e6
	timeStr := strconv.Itoa(int(time))
	//timestamp + random digest = data
	data := timeStr + strconv.Itoa(rand.Intn(100000))
	dataMd5 := hexmd5(data)
	// timestamp + MD5(data)
	return b64.EncodeToString([]byte(dataMd5))
}

