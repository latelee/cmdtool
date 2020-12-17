/*
共用底层函数
*/
package common

import (
	"fmt"

	conf "github.com/latelee/cmdtool/common/conf"
	//"errors"
	// "os"
	// "strings"
)

// rpad adds padding to the right of a string.
func rpad(s string, padding int) string {
	template := fmt.Sprintf("%%-%ds", padding)
	return fmt.Sprintf(template, s)
}

// 输出到终端的
func PrintHelpInfo(theCmd []conf.UserCmdFunc) {
	var cmdMaxLen int = 0
	fmt.Printf("Available Commands:\n");

	// 找最大的字符长度，方便对齐
	for _, item := range theCmd {
		nameLen := len(item.Name)
		if nameLen > cmdMaxLen {
			cmdMaxLen = nameLen
		}
	}
	
	for _, item := range theCmd {
		fmt.Printf("  %v %v\n", rpad(item.Name, cmdMaxLen), item.ShortHelp)
    }
}

//返回字符串的
func GetHelpInfo(theCmd []conf.UserCmdFunc) (ret string) {
	var cmdMaxLen int = 0
	ret = fmt.Sprintf("Available Commands:\n");

	for _, item := range theCmd {
		nameLen := len(item.Name)
		if nameLen > cmdMaxLen {
			cmdMaxLen = nameLen
		}
	}
	
	for _, item := range theCmd {
		ret += fmt.Sprintf("  %v %v\n", rpad(item.Name, cmdMaxLen), item.ShortHelp)
	}
	
	return
}