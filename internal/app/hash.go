package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"log"
)

func getHash(s string) string {
	h := md5.New()
	io.WriteString(h, s)
	hash := fmt.Sprintf("%x", h.Sum(nil))
	hash = hash[0:8]
	log.Println(hash)
	return hash
}