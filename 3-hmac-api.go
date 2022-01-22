package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

var DB = map[string]string{
	"User1Key": "User1Secret",
	"User2Key": "User2Secret",
}

func Server(apikey, sign string, data []byte) {
	apiSecret := DB[apikey]
	h := hmac.New(sha256.New, []byte(apiSecret))
	h.Write(data)
	expectedHMAC := hex.EncodeToString(h.Sum(nil))
	fmt.Println(sign == expectedHMAC)
}

func main() {
	const apikey = "User1Key"
	const apiSecret = "Usr1Secret"

	h := hmac.New(sha256.New, []byte(apiSecret))
	data := []byte("data")
	h.Write(data)
	sign := hex.EncodeToString(h.Sum(nil))
	fmt.Println(sign)

	Server(apikey, sign, data)
}
