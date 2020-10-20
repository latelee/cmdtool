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
	"os"
	"bufio"
	"encoding/base64"
	"encoding/hex"
	"io"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/md5"
	"errors"
)

// base64 encode
func Base64Encode(str string) string {
	return base64.StdEncoding.EncodeToString([]byte(str))
}

// base64 decode
func Base64Decode(str string) (string, error) {
	s, e := base64.StdEncoding.DecodeString(str)
	return string(s), e
}


// AESGCMEncrypt encrypts plaintext with the given key using AES in GCM mode.
func AESGCMEncrypt(key, plaintext []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err := rand.Read(nonce); err != nil {
		return nil, err
	}

	ciphertext := gcm.Seal(nil, nonce, plaintext, nil)
	return append(nonce, ciphertext...), nil
}

// AESGCMDecrypt decrypts ciphertext with the given key using AES in GCM mode.
func AESGCMDecrypt(key, ciphertext []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	size := gcm.NonceSize()
	if len(ciphertext)-size <= 0 {
		return nil, errors.New("Ciphertext is empty")
	}

	nonce := ciphertext[:size]
	ciphertext = ciphertext[size:]

	plainText, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return nil, err
	}

	return plainText, nil
}

func Md5(str string) string {
	return Md5ForByte([]byte(str))
}

func Md5ForByte(data []byte) string {
	h := md5.New()
	h.Write(data)
	return hex.EncodeToString(h.Sum(nil))
}

func Md5ForFile(filename string) string {
	file, err := os.Open(filename)
	if err != nil {
		return ""
	}
	defer file.Close()
	freader := bufio.NewReader(file)
	h := md5.New()
	io.Copy(h, freader)
	return hex.EncodeToString(h.Sum(nil))
}
