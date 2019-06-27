package utils

import (
	"fmt"
	"log"
	"strconv"
)

func Int64(in string) int64 {
	out, err := strconv.ParseInt(in, 10, 64)
	if err != nil {
		log.Printf("string[%s] covert int64 fail. %s", in, err)
		return 0
	}
	return out
}

func Int(in string) int {
	out, err := strconv.ParseInt(in, 10, 64)
	if err != nil {
		log.Printf("string[%s] covert int fail. %s", in, err)
		return 0
	}
	return int(out)
}

func Int32(in string) int32 {
	out, err := strconv.ParseInt(in, 10, 32)
	if err != nil {
		log.Printf("string[%s] covert int64 fail. %s", in, err)
		return 0
	}
	return int32(out)
}

func Str(v interface{}) string {
	return fmt.Sprintf("%v", v)
}
