package utils

import (
	"crypto/md5"
	"crypto/sha1"
	"fmt"
)

var (
	Md5  func(string) string = Md5String
	Sha1 func(string) string = Sha1String
)

func Md5Byte(in []byte) []byte {
	h := md5.New()
	h.Write(in)
	return h.Sum(nil)
}

func Md5String(in string) string {
	return fmt.Sprintf("%x", Md5Byte([]byte(in)))
}

func Sha1Byte(in []byte) []byte {
	h := sha1.New()
	h.Write(in)
	return h.Sum(nil)
}

func Sha1String(in string) string {
	return fmt.Sprintf("%x", Sha1Byte([]byte(in)))
}
