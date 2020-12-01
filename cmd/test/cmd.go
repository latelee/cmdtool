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

func NewCmdTest() *cobra.Command{

    var cmd = &cobra.Command{
        Use:     name,
        Short:   shortDescription,
        Long:    longDescription,
        Example: example,
        RunE: func(cmd *cobra.Command, args []string) error {
			//klog.Println(common.DBName)
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
			klog.Printf("cmd '%v' not support", args[0])
			common.PrintHelpInfo(theCmd)
            return nil
        },
    }
    // note：使用子命令形式，下列可在help中展开
    // 命令参数，保存的值，参数名，默认参数，说明
    //cmd.Flags().StringVar(&mode, "db", "-", "set the database name")

    return cmd
}
