package main

import (
	"crypto/md5"
	"fmt"
	"strings"
)

func main() {
	hearerStr := "Digest username=\"root\", realm=\"pmplatform\", nonce=\"MTU0MjY5NzMzNDIxODo0NjU4OGE2Nzk3NDk0NDVlNmNjMmU5MDM3OGEyOGQyYQ==\", uri=\"/VIID/System/Register\", cnonce=\"MmI0ZjRhOTEwZDI1YmI5OGY1ODZkODMxYTY5MDMxMzk=\", nc=00000001, qop=auth, response=\"3fcde23656ea985e30240866b2186e74\""

	fmt.Println(strings.Index(hearerStr, " "))
	fmt.Println(hearerStr[strings.Index(hearerStr, " ") + 1 :])
	headerStrScheme := strings.TrimSpace(hearerStr[strings.Index(hearerStr, " ") + 1 :])
	m := make(map[string]string)
	keyValues := strings.Split(headerStrScheme, ",")
	for _, v := range keyValues {
		if strings.Contains(v, "=") {
			poxi := strings.Index(v, "=")
			key := strings.TrimSpace(v[0:poxi])
			value := strings.TrimSpace(  strings.Replace(v[poxi+1:], "\"", "", len(keyValues)))
			m[key] = value
		}
	}

	fmt.Println(m)

	fmt.Println(string(md5String("aaaa")))

	//a := "   dd dd  "
	//fmt.Println(len(a))
	//b := strings.TrimSpace(a)
	//fmt.Println(len(b))
	//fmt.Println(headerStrScheme)
	//s := "我是中国侠"
	//fmt.Println(s[1:3])

}

func md5String(data string)[]byte{
	h := md5.New()
	h.Write([]byte(data))
	return h.Sum(nil)
}