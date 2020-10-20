// Copyright 2013 com authors
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

package com

import (
	"fmt"
	"testing"
)

func TestHtml(t *testing.T) {
	htm := "<div id=\"button\" class=\"btn\">Click me</div>\n\r"
	js := string(Html2JS([]byte(htm)))

	fmt.Println(htm, js)

	fmt.Println("encode: ", UrlEncode("hello world"))
	dec, err := UrlDecode("hello world")
	if err != nil {
		fmt.Println("decode err: ", err)
	} else {
		fmt.Println("decode: ", dec)
	}

}