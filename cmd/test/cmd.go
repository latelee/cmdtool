package cmd

import (
    // "fmt"

    "github.com/spf13/cobra"
	_ "github.com/spf13/pflag"
	
	"k8s.io/klog"
	conf "github.com/latelee/cmdtool/common/conf"
	common "github.com/latelee/cmdtool/common"
	
)

var (
    name = `test`
    shortDescription = `test command`
    longDescription  = `test...
`
    example = `  example comming up...
`
)

var theCmd = []conf.UserCmdFunc{
    conf.UserCmdFunc {
        Name: "foo",
        ShortHelp: "just a foo help info",
        Func: foo,
    },
    conf.UserCmdFunc {"watch", "watch config file", testWatch,},
}

var mode int

func NewCmdTest() *cobra.Command{

    var cmd = &cobra.Command{
        Use:     name,
        Short:   shortDescription,
        Long:    longDescription + "\n" + common.GetHelpInfo(theCmd),
        Example: example,
        RunE: func(cmd *cobra.Command, args []string) error {
			// 1 没有参数
			if (len(args) == 0) {
				//klog.Warning("no args found")
				//common.PrintHelpInfo(theCmd)
				cmd.Help()
				return nil
			}

			// 2 遍历是否有合法的参数，如无则提示
			idx := -1
			for idx1, item := range theCmd {
				if (args[0] == item.Name) {
					idx = idx1 // why ???
					break
				}
			}
			if idx == -1 {
				klog.Printf("arg '%v' not support", args[0])
				cmd.Help()
				return nil
			}
			
			// 3 执行公共的初始化
			klog.Printf("bussiness init, mode: %v\n", mode) // just test

			// 4 执行命令
			theCmd[idx].Func(args)

			return nil
        },
    }
    // note：使用子命令形式，下列可在help中展开
    // 命令参数，保存的值，参数名，默认参数，说明
	cmd.Flags().IntVarP(&mode, "mode", "m", 0, "set the test mode")

    return cmd
}
