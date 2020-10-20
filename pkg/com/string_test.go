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

func TestString(t *testing.T) {

	fmt.Println("isletter: ", IsLetter('1'))
	fmt.Println("reverse ", Reverse("abcd中文"))

	fmt.Println("to snake ", ToSnakeCase("HttpServer"))
}

func TestExpand(t *testing.T) {
	match := map[string]string{
		"domain":    "gowalker.org",
		"subdomain": "github.com",
	}
	s := "http://{domain}/{subdomain}/{0}/{1}"
	sR := "http://gowalker.org/github.com/unknwon/gowalker"
	if Expand(s, match, "unknwon", "gowalker") != sR {
		t.Errorf("Expand:\n Expect => %s\n Got => %s\n", sR, s)
	}
}