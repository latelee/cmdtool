package com

import (
	"fmt"
	"testing"
)

/*
文件、路径等测试
*/
func TestFile(t *testing.T) {
	
	// 当前目录
	fmt.Println("pwd: ", GetPWD())
	// 文件大小
	fmt.Println("filesize: ", FileSize("foo.txt"))

	// 文件目录判断
	fmt.Println("IsFile: ", IsFile("file.go"), IsFile("file1.go"), IsFile("testdata"), IsExist("file_test.go"))

	fmt.Println("IsDir: ", IsDir("file.go"), IsDir("file1.go"), IsDir("testdata"), IsDir("file_test.go"))

	// 创建目录
	fmt.Println("mkdir ", MkDir("foo"), MkDir("foo"))

	// 删除目录
	fmt.Println("rmdir ", RmDir("foo1"), RmDir("foo"))

	// 列出子文件或子目录，返回文件名称列表（true为包含子目录）
	files, err := StatDir("testdata", true)
	if (err == nil) {
		fmt.Println(files, err)
	}

	// 递归拷贝目录，如目标目录存在，则失败
	err = CopyDir("testdata", "testdata2")
	if err == nil {
		fmt.Println("copy dir ok")
	} else {
		fmt.Println(err)
	}

	// 文件拷贝
	err = CopyFile("file.go.new", "file.go")
	if err != nil {
		fmt.Println("copy file error:", err)
	} else {
		fmt.Println("copy file ok")
	}

}

