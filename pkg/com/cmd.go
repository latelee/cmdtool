// +build go1.3

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

// Package com is an open source project for commonly used functions for the Go programming language.
package com

import (
	"bytes"
	"fmt"
	"os/exec"
	"io"
	"os"
	"strings"
	"sync"
)

// ExecCmdDirBytes executes system command in given directory
// and return stdout, stderr in bytes type, along with possible error.
func ExecCmdDirBytes(dir, cmdName string, args ...string) ([]byte, []byte, error) {
	bufOut := new(bytes.Buffer)
	bufErr := new(bytes.Buffer)

	cmd := exec.Command(cmdName, args...)
	cmd.Dir = dir
	cmd.Stdout = bufOut
	cmd.Stderr = bufErr

	err := cmd.Run()
	return bufOut.Bytes(), bufErr.Bytes(), err
}

// ExecCmdBytes executes system command
// and return stdout, stderr in bytes type, along with possible error.
func ExecCmdBytes(cmdName string, args ...string) ([]byte, []byte, error) {
	return ExecCmdDirBytes("", cmdName, args...)
}

// ExecCmdDir executes system command in given directory
// and return stdout, stderr in string type, along with possible error.
func ExecCmdDir(dir, cmdName string, args ...string) (string, string, error) {
	bufOut, bufErr, err := ExecCmdDirBytes(dir, cmdName, args...)
	return string(bufOut), string(bufErr), err
}

// ExecCmd executes system command
// and return stdout, stderr in string type, along with possible error.
func ExecCmd(cmdName string, args ...string) (string, string, error) {
	return ExecCmdDir("", cmdName, args...)
}


/////////////////////////////////////////////////////

//Command defines commands to be executed and captures std out and std error
type Command struct {
    Cmd    *exec.Cmd
    StdOut []byte
    StdErr []byte
}

//Execute executes the command and captures the output in stdOut
func (cm *Command) Execute() {
    var err error
    cm.StdOut, err = cm.Cmd.Output()
    if err != nil {
        fmt.Println("Output failed: ", err)
        cm.StdErr = []byte(err.Error())
    }
}

//GetStdOut gets StdOut field
func (cm Command) GetStdOut() string {
    if len(cm.StdOut) != 0 {
        return strings.TrimRight(string(cm.StdOut), "\n")
    }
    return ""
}

//GetStdErr gets StdErr field
func (cm Command) GetStdErr() string {
    if len(cm.StdErr) != 0 {
        return strings.TrimRight(string(cm.StdErr), "\n")
    }
    return ""
}

//ExecuteCmdShowOutput captures both StdOut and StdErr after exec.cmd().
//It helps in the commands where it takes some time for execution.
func (cm Command) ExecuteCmdShowOutput() error {
    var stdoutBuf, stderrBuf bytes.Buffer
    stdoutIn, _ := cm.Cmd.StdoutPipe()
    stderrIn, _ := cm.Cmd.StderrPipe()

    var errStdout, errStderr error
    stdout := io.MultiWriter(os.Stdout, &stdoutBuf)
    stderr := io.MultiWriter(os.Stderr, &stderrBuf)
    err := cm.Cmd.Start()
    if err != nil {
        return fmt.Errorf("failed to start '%s' because of error : %s", strings.Join(cm.Cmd.Args, " "), err.Error())
    }

    var wg sync.WaitGroup
    wg.Add(1)

    go func() {
        _, errStdout = io.Copy(stdout, stdoutIn)
        wg.Done()
    }()

    _, errStderr = io.Copy(stderr, stderrIn)
    wg.Wait()

    err = cm.Cmd.Wait()
    if err != nil {
        return fmt.Errorf("failed to run '%s' because of error : %s", strings.Join(cm.Cmd.Args, " "), err.Error())
    }
    if errStdout != nil || errStderr != nil {
        return fmt.Errorf("failed to capture stdout or stderr")
    }

    cm.StdOut, cm.StdErr = stdoutBuf.Bytes(), stderrBuf.Bytes()
    return nil
}

// runCommandWithShell executes the given command with "sh -c".
// It returns an error if the command outputs anything on the stderr.
func RunCommandWithShell(command string) (string, error) {
	cmd := &Command{Cmd: exec.Command("sh", "-c", command)}
	err := cmd.ExecuteCmdShowOutput()
	if err != nil {
		return "", err
	}
	errout := cmd.GetStdErr()
	if errout != "" {
		return "", fmt.Errorf("%s", errout)
	}
	return cmd.GetStdOut(), nil
}
