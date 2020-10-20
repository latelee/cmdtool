
package com

import (
	"fmt"
	//"strings"
	"testing"
)

type TestObj struct {
	Name  string
	Value uint64
	Size  int32
	Guard float32
}

func TestConvert(t *testing.T) {
	var s StrTo = StrTo([]byte("100695")[3:])
	var s1 StrTo = StrTo([]byte("000102B7")[4:])
	fmt.Println("str to int ", s.Int(), s1.IntHex())
	fmt.Println("conv ", ToStr(100.56))

	fmt.Println("Tofixed: ", ToFixed(6.5, 2))
	fmt.Println("Round: ", Round(0.65, 4))

	hex := ToHexByte("dd")
	fmt.Printf("hex1 %v \n", hex)
	hex = ToHexByte("deadbeef")
	fmt.Printf("hex2 %v \n", hex)

	str := ToHexString(hex)
	fmt.Printf("hexstr %v \n", str)
	
    object := TestObj{
		Name:  "James",
		Value: 128,
		Size:  256,
		Guard: 56.4,
	}
    
    data, err := Marshal(object)
	if err != nil {
		fmt.Printf("marshal failed")
	}
    fmt.Println("data: ", data)
    
    var o TestObj
    err = Unmarshal(data, &o)
    if err != nil {
		fmt.Printf("Unmarshal failed")
	}
	fmt.Println("data: ", o)
	
	fmt.Println(GetStructFieldName(&TestObj{}))
	fmt.Println(GetStructFieldType(&TestObj{}))
	
}