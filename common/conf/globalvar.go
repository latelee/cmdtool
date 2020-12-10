/*
全局变量

参数控制 

*/
package conf

import (
	// "fmt"
)

var FlagDBServer string
var FlagTimeout string
var FlagName string
var FlagPrint bool
var FlagOutputFile string

// 命令列表，包括名称，帮助信息
type UserCmdFunc struct {
    Name string
    ShortHelp string
    // LongHelp string
    Func func(args []string)
}
