#!/bin/sh

# 版本和编译时间 TODO：找一个好的方法：
Version="v1.0"
BuildTime=`date +'%Y-%m-%d %H:%M:%S'`
#echo "package cmd" > cmd/ver.go
#echo "const BuildTime1 = \"${BuildTime}\"" >> cmd/ver.go
#echo "const Version1 = \"${Version}\"" >> cmd/ver.go

GO111MODULE=on go build -ldflags "-X 'github.com/latelee/cmdtool/cmd.BuildTime=${BuildTime}' -X 'github.com/latelee/cmdtool/cmd.Version=${Version}'" -mod vendor -o cmdtool.exe main.go
