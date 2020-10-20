
package com

import (
	"fmt"
    //"os/exec"
	//"strings"
	"testing"
)

func TestBytes(t *testing.T) {
    // 数组转十六进制，追加形式
    b := []byte{'L', 'a', 't', 'e', 100}
    c := ToHex(b)
    c1 := Bytes2Hex(b)
	fmt.Println("c: ", c, c1)
    
    d := FromHex(c)
    d1 := Hex2Bytes(c)
    e := FromHex("0x123")
    fmt.Println("d: ", d, d1, e)
}

