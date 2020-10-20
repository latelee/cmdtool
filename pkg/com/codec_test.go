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
	"bytes"
	"crypto/rand"
	"testing"
)

func TestAES(t *testing.T) {
	t.Parallel()

	key := make([]byte, 16) // AES-128
	_, err := rand.Read(key)
	if err != nil {
		t.Fatal("Failed to create 128 bit AES key: " + err.Error())
	}

	fmt.Println("key: ", key)
	plaintext := []byte("this will be encrypted")

	ciphertext, err := AESGCMEncrypt(key, plaintext)
	if err != nil {
		t.Fatal("Failed to encrypt plaintext: " + err.Error())
	}
	fmt.Println("ciphertext: ", ciphertext)
	decrypted, err := AESGCMDecrypt(key, ciphertext)
	if err != nil {
		t.Fatal("Failed to decrypt ciphertext: " + err.Error())
	}
	fmt.Println("decrypted: ", decrypted)
	if bytes.Compare(decrypted, plaintext) != 0 {
		fmt.Println("Decryption was not performed correctly")
	} else {
		fmt.Println("ok")
	}
}

func TestMd5(t *testing.T) {
	fmt.Println("md5: ", Md5ForFile("codec_test.go"))
}