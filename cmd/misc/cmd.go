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
    name = `misc`
    shortDescription = `misc command`
    longDescription  = `misc command for test or other use...
  杂项命令，如临时测试使用...
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
}

func RegisterCmd() *cobra.Command {
    var cmd = &cobra.Command{
        Use:     name,
        Short:   shortDescription,
        Long:    longDescription,
        Example: example,
        RunE: func(cmd *cobra.Command, args []string) error {
			if (len(args) == 0) {
				klog.Warning("no args found")
				common.PrintHelpInfo(theCmd)
				return nil
			}
			for _, item:=range theCmd {
				if (args[0] == item.Name) {
					item.Func(args)
					return nil
				}
			}
			klog.Printf("arg '%v' not support", args[0])
			common.PrintHelpInfo(theCmd)
            return nil
        },
    }

    // 命令参数，保存的值，参数名，默认参数，说明
    //cmd.Flags().StringVar(&mode, "mode", "", "set the test mode")

    return cmd
}