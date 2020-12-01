package cmd

import (
    "fmt"
    
    "k8s.io/klog"
    "github.com/spf13/cobra"
	_ "github.com/spf13/pflag"
	conf "github.com/latelee/cmdtool/common/conf"
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

func printHelpInfo() {
	klog.Println("valid cmd: ");
	for _, item:=range theCmd {
        fmt.Println(item.Name, "\t:", item.ShortHelp)
    }
}

func NewCmdMisc() *cobra.Command {
    var cmd = &cobra.Command{
        Use:     name,
        Short:   shortDescription,
        Long:    longDescription,
        Example: example,
        RunE: func(cmd *cobra.Command, args []string) error {
           // klog.Printf("argv: %s\n", args)
            if (len(args) != 0 && args[0] == "foo") {
                foo(args)
            } else {
				klog.Printf("cmd '%v' not support", args[0])
				printHelpInfo()
				return nil
			}
            return nil
        },
    }

    // 命令参数，保存的值，参数名，默认参数，说明
    //cmd.Flags().StringVar(&mode, "mode", "", "set the test mode")

    return cmd
}