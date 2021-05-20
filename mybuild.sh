#!/bin/sh

target=cmdtool.exe
#VER_FILE=cmd/version.h

# 版本和编译时间 TODO：找一个好的方法：
Version="v1.0"
BuildTime=`date +'%Y-%m-%d %H:%M:%S'`

GIT_VERSION=$Version" build: "$BuildTime

#echo "Generated" $VER_FILE "for version:" $GIT_VERSION
#
#echo "#ifndef PROJECT_VERSION_H" > $VER_FILE
#echo "#define PROJECT_VERSION_H" >> $VER_FILE
#echo "" >> $VER_FILE
#echo "#define VERSION_NUMBER \"$GIT_VERSION\"" >> $VER_FILE
#echo "" >> $VER_FILE
#echo "#endif" >> $VER_FILE
#
#echo "Job done!!"

#GO111MODULE=on go build  -mod vendor -o $target main.go || exit 1

GO111MODULE=on go build -ldflags "-X 'cmdtool/cmd.Version=${Version}'" -mod vendor -o $target main.go || exit 1

exit 0

sleep 1
strip $target