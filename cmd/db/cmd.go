package cmd

import (
    // "fmt"
    "k8s.io/klog"
    "github.com/spf13/cobra"
	_ "github.com/spf13/pflag"
	conf "cmdtool/common/conf"
	common "cmdtool/common"
)

var (
    name = `db`
    shortDescription = `db command`
	longDescription  = 
`
db command for test or other use...
`

    example = `  example comming up...
`
)

var mode int

var theCmd = []conf.UserCmdFunc{
    conf.UserCmdFunc {
        Name: "foo",
        ShortHelp: "foo help info",
        Func: foo,
    },
}

func myHelp(cmd *cobra.Command, args []string) {
	common.PrintHelpInfo(theCmd)
}

func RegisterCmd() *cobra.Command {
    var cmd = &cobra.Command{
        Use:     name,
        Short:   shortDescription,
        Long:    longDescription + "\n" + common.GetHelpInfo(theCmd),
		Example: example,
		// Args: cobra.ExactArgs(1), // 只接受一个参数，但只提示简单的非法信息，要加上-h才显示帮助信息，权衡后弃之
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

	// 添加自定义的帮助函数，如加，则自己要实现显示flag的信息，麻烦，舍之
	//cmd.SetHelpFunc(myHelp)

    // 命令参数，保存的值，参数名，默认参数，说明
    cmd.Flags().IntVarP(&mode, "mode", "m", 0, "set the test mode")

    return cmd
}