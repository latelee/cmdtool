
package com

import (
	"fmt"
    "os"
	"testing"
)

func getNameAndVer(addr, port string, buf []byte) string {
    idx1, idx2 := GetContinuedIdx(buf)
    return fmt.Sprintf("%s", buf[idx1:idx2])
}

func TestBinbuffer(t *testing.T) {
    
    fmt.Println("test...")

    strData := "4C456400640001002009110768656C6C6F776F726C64992019"

    //var reader BufferReader; //{[]byte(ToHexByte(strData)), 0
    reader := NewBufferReader([]byte(ToHexByte(strData)))

    reader.Init([]byte(ToHexByte(strData)))
    
    fmt.Printf("0x%x\n", reader.ReadUint8())
    reader.Dump()
    
    //var writer BufferWriter
    //writer.Init(100)
    
    writer := NewBufferWriter(100)
    writer.WriteUint8(0x4c)
    writer.Dump()
    
    fmt.Println("--------------")

    
    // dada010100230144525253303256332e310bece798621a2f3389c973ea215d9d2ee5a4a48cfd7178ab8459
    strData = "dada010100230144525253303256332e310bece798621a2f3389c973ea215d9d2ee5a4a48cfd7178ab8459"
    hexByte := []byte(ToHexByte(strData))
    
    idx1, idx2 := GetContinuedIdx(hexByte)
    Dump1(os.Stdout, hexByte, len(hexByte))
    fmt.Println("----")
    

    ver := getNameAndVer("j100", "ddd", hexByte)
    fmt.Println("fuc", ver)
    fmt.Printf("ver1: %d %d %s\n", idx1, idx2, hexByte[idx1:idx2])
    
    
    foo := map[string]int{
         "hello": 100,
         "money": 250,
     }

    file, err := os.Open("file.go3") // For read access.
    if err != nil {
        PrintByLine(os.Stderr, err)
    } else {
        file.Close()
    }
    
    PrintByLine(os.Stdout, foo)
    fmt.Println(foo)
    fmt.Printf("%+v\n", foo)
    
    
}

