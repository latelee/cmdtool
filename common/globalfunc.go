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

func Foo() {

}


func PrintHelpInfo(theCmd []conf.UserCmdFunc) {
	fmt.Println("valid command: ");
	for _, item:=range theCmd {
        fmt.Println(item.Name, "\t:", item.ShortHelp)
    }
}