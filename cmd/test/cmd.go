package cmd

import (
    //"fmt"

    "github.com/spf13/cobra"
	_ "github.com/spf13/pflag"
	
	"k8s.io/klog"
)

var (
    name = `test`
    shortDescription = `  test command`
    longDescription  = `  test...
`
    example = `  example comming up...
`
)

type UserCmdFunc struct {
	name string
	fn func(args []string)
}

// theCmd := &UserCmdFunc {
// 	"foo", foo(args []string)
// }

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
				return nil
			}

            if (args[0] == "foo"){
                foo(args)
            } else if (args[0] == "watch"){
                testWatch(args)
            } else {
				klog.Printf("cmd '%v' not support", args[0])
				return nil
			} 
            return nil
        },
    }
    // note：使用子命令形式，下列可在help中展开
    // 命令参数，保存的值，参数名，默认参数，说明
    //cmd.Flags().StringVar(&mode, "db", "-", "set the database name")

    return cmd
}
