// Copyright 2016 The laik Authors. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"): you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
// License for the specific language governing permissions and limitations
// under the License.

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
