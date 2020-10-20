package com

/*
默认为小端格式
TODO：有符号数、浮点数
*/
import (
    "fmt"
    //"io"
    "strconv"
    //"bytes"
    "encoding/hex"
    //"encoding/binary"
    
)

type BufferReader struct {
    Buffer []byte
    offset int
    length int
}


func NewBufferReader(buf []byte) *BufferReader {
	return &BufferReader{
        Buffer: buf,
        length: len(buf),
        offset: 0,
    };
}

func NewBufferReader1(buf []byte, len int) *BufferReader {
	return &BufferReader{
        Buffer: buf,
        length: len,
        offset: 0,
    };
}

func (b *BufferReader) Init(buf []byte) {
    b.Buffer = buf;
    b.offset = 0;
}

func (b *BufferReader) Dump() {
    fmt.Printf("in BufferReader (len %d)\n", b.length)
    Dump(b.Buffer, b.length)
    //fmt.Printf("%v\n", hex.Dump(b.Buffer))
}

func (b *BufferReader) ReadUint8() (o uint8) {
    //fmt.Printf("%x", b.Buffer);
    
    o = b.Buffer[b.offset];
    b.offset += 1;
    return;
}

func (b *BufferReader) SkipBytes(n int) {
    b.offset += n;

    return;
}

func (b *BufferReader) backBytes(n int) {
    b.offset -= n;

    return;
}

// 返回当前的字符，但偏移量不递增
func (b *BufferReader) peekByte(o byte) {
    o = b.Buffer[b.offset];

    return;
}

func (b *BufferReader) ReadUint16() (o uint16) {
    var u1, u2 uint16;
    u1 = uint16(b.Buffer[b.offset]);
    u2 = uint16(b.Buffer[b.offset+1]);
    o = u2 << 8 | u1;
    b.offset += 2;

    return;
}

func (b *BufferReader) ReadUint32() (o uint32) {
    var u1, u2, u3, u4 uint32;
    u1 = uint32(b.Buffer[b.offset]);
    u2 = uint32(b.Buffer[b.offset+1]);
    u3 = uint32(b.Buffer[b.offset+2]);
    u4 = uint32(b.Buffer[b.offset+3]);
    o = (u4<<24) | (u3<<16) | (u2<<8) | u1;
    
    b.offset += 4;

    return;
}

func (b *BufferReader) ReadUint64() (o uint64) {
    var u1, u2, u3, u4, u5, u6, u7, u8 uint64;
    u1 = uint64(b.Buffer[b.offset]);
    u2 = uint64(b.Buffer[b.offset+1]);
    u3 = uint64(b.Buffer[b.offset+2]);
    u4 = uint64(b.Buffer[b.offset+3]);
    u5 = uint64(b.Buffer[b.offset]);
    u6 = uint64(b.Buffer[b.offset+1]);
    u7 = uint64(b.Buffer[b.offset+2]);
    u8 = uint64(b.Buffer[b.offset+3]);
    o = (u8<<65) | (u7<<48) | (u6<<40) | (u5<<32) | (u4<<24) | (u3<<16) | (u2<<8) | u1;
    
    b.offset += 4;

    return;
}

func (b *BufferReader) ReadBytes(n int) (o []byte) {
    o = b.Buffer[b.offset: b.offset+n];
    
    b.offset += n;
    
    return;
}

// 转为十六进制格式的字符串
func (b *BufferReader) ReadHexString(n int) (o string) {
    buf := b.Buffer[b.offset: b.offset+n];
    o = hex.EncodeToString(buf);

    b.offset += n;

    return;
}

// 只能读取正常的字符串，即可打印的
func (b *BufferReader) ReadString(n int) (o string) {
    o = string(b.Buffer[b.offset: b.offset+n]);

    b.offset += n;

    return;
}

// 读取bcd码，如十六进制的0x20，转换十进制的20
func (b *BufferReader) ReadBCD() (o int) {
    b1 := hex.EncodeToString(b.Buffer[b.offset: b.offset+1]);
    o1, _ := strconv.ParseInt(b1, 10, 8);
    o = int(o1);

    b.offset += 1;

    return;
}

// 与HexString相同，只是原始字符表现形式不同
func (b *BufferReader) ReadBCDString(n int) (o string) {
    o = hex.EncodeToString(b.Buffer[b.offset: b.offset+n]);
    b.offset += n;

    return;
}

// n为1、2、4
func (b *BufferReader) ReadBCDNumber(n int) (o int) {
    b1 := hex.EncodeToString(b.Buffer[b.offset: b.offset+n]);
    o1, _ := strconv.ParseInt(b1, 10, 8*n);
    
    o = int(o1);

    b.offset += n;

    return;
}

func None() {
    fmt.Println("none...");
}

/*
// 转入十六进制形式的字符，输出为byte类型
func ToHexByte(str string) (ob []byte) {
    ob, _ = hex.DecodeString(str);
    
    return;
}

// 转入十六进制形式的数组，输出为对应的字符
// 如 4c 77数组，将转换成4c77字符串，可保存到文件
func ToHexString(b []byte) (ostr string) {
    ostr = hex.EncodeToString(b);
    return;
}
*/

