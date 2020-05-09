package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
	"io"
)

func main() {
	c := getCode("test@gmail.com")
	fmt.Println(c)
	c = getCode("test@gmail.com")
	fmt.Println(c)
}
func getCode(str string) string{
	h := hmac.New(sha256.New,[]byte("private key"))
	io.WriteString(h,str)
	return fmt.Sprintf("%x",h.Sum(nil))
}