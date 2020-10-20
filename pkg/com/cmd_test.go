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
    "os/exec"
	"strings"
    //"bytes"
	"testing"
)

func TestExecCmd1(t *testing.T) {
	stdout, stderr, err := ExecCmd("go", "help", "get")
	if err != nil {
		t.Errorf("ExecCmd:\n Expect => %v\n Got => %v\n", nil, err)
	} else if len(stderr) != 0 {
		t.Errorf("ExecCmd:\n Expect => %s\n Got => %s\n", "", stderr)
	} else if !strings.HasPrefix(stdout, "usage: go get") {
		t.Errorf("ExecCmd:\n Expect => %s\n Got => %s\n", "usage: go get", stdout)
	}
    
    stdout, stderr, err = ExecCmd("sh", "-c", "ls | grep .go")
    if err != nil {
        fmt.Println(stderr, err)
    } else {
        fmt.Println(stdout)
    }
}

func TestExecCmd2(t *testing.T) {
    c := &Command{Cmd: exec.Command("sh", "-c", "ls | grep go")}
	c.Execute()
	fmt.Println(c.GetStdOut())
    
    stdout, err := RunCommandWithShell("ls | grep go")
	if err != nil {
        fmt.Println(err)
	}
	fmt.Println(stdout)
}

func TestExecCmd3(t *testing.T) {
    file1 := "201909020001"
    file2 := "201909020001"
    cmdarg := fmt.Sprintf("lame -h %s.wav %s.mp3", file1, file2)
    //c := &Command{Cmd: exec.Command("sh", "-c", "lame -h 201909020001.wav 201909020001.mp3")}
    c := &Command{Cmd: exec.Command("sh", "-c", cmdarg)}
	c.Execute()
    fmt.Println("cmd: ", c.Cmd.Args)
    errInfo := c.GetStdErr()
    if errInfo != "" {
        fmt.Println("error ", errInfo)
    }
    fmt.Println("==================")

    // 注：此方式亦不行
    stdout, stderr, err := ExecCmd("lame", "-h", "201909020001.wav 201909020002.mp3")
    fmt.Println(stdout, stderr, err)

// 原始API，失败
/*
    cmd := exec.Command("sh", "-c", "lame", "-h", "201909020001.wav 201909020002.mp3")
    fmt.Println("cmd: ", cmd)
    var out bytes.Buffer
    var stderr bytes.Buffer
    cmd.Stdout = &out
    cmd.Stderr = &stderr
    err = cmd.Run()
    if err != nil {
        fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
    } else {
        fmt.Println(out.String())
    }
*/
}


func BenchmarkExecCmd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ExecCmd("go", "help", "get")
	}
}
