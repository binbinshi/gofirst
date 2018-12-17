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
var (
	nonce = ""
)

func main() {
	start()
	fmt.Println("Ticker stopped")
}


func start(){
	ticker := time.NewTicker(time.Hour * 1)
	caculateNonce()
	for  {
		select {
		case <-ticker.C:
			go func() {
				fmt.Println("start ....")
				caculateNonce()
				}()
		}
	}
}

func caculateNonce(){
	fmt.Println("start caculate ....")
	time := time.Now().UnixNano()/1e6
	timeStr := strconv.Itoa(int(time))
	data := timeStr + strconv.Itoa(rand.Intn(100000))
	h := md5.New()
	h.Write([]byte(data))
	dataMd5 := h.Sum(nil)
	nonce = base64.URLEncoding.EncodeToString([]byte(timeStr + ":" + hex.EncodeToString(dataMd5)))
	fmt.Printf("end .... createNonce %s %s\n", timeStr, nonce)
}