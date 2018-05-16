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
