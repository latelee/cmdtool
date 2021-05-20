package cmd

import (
	"fmt"
	"os"

	// "io/ioutil"
	"bytes"
	"path/filepath"

	//"golang.org/x/net/context"
	"github.com/spf13/cobra"
	// "github.com/spf13/afero"
	"github.com/spf13/viper"

	"github.com/fsnotify/fsnotify"

	test "cmdtool/cmd/test"
	conf "cmdtool/common/conf"
	"cmdtool/pkg/klog"
)

/*
#include <stdio.h>
#include "version.h"

char* GetVersion()
{
    static char buffer[64] = {0};

    //snprintf(buffer, 64, " %s build: %s %s\r\n", "v1.0", __DATE__, __TIME__);
    snprintf(buffer, 64, " %s\r\n", VERSION_NUMBER);

    return buffer;
}
*/
import "C"

var (
	cfgFile   string
	BuildTime string
	Version   string

	runmode string

	longDescription = `  cmd test tool.
  【中文样例】命令终端测试示例工具。
`
	example = `  comming soon...
`
)

func getVersion1() string {
	return fmt.Sprintf("  %v build: %v\n", Version, BuildTime)
}

// c调用
func getVersion() string {
	name1 := C.GetVersion()
	name := C.GoString(name1)
	return fmt.Sprintf(" %v", name)
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:     filepath.Base(os.Args[0]),
	Short:   "cmd tool",
	Long:    getVersion() + longDescription,
	Example: example,
	Version: getVersion(), //, //Version,
	// PreRun: func(cmd *cobra.Command, args []string) {
	// },
	// 一级命令
	Run: func(cmd *cobra.Command, args []string) {
		klog.Printf("cobra demo program, with args: %v\n", args)
		if runmode == "auto" {
			klog.Println("in mode auto")
		} else if runmode == "upgrade" {
			klog.Println("in mode upgrade")
		} else {
			klog.Println("default mode")
		}
		// for {
		// }
	},
	// PostRun: func(cmd *cobra.Command, args []string) {
	// },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() error {
	rootCmd.AddCommand(test.RegisterCmd())

	return rootCmd.Execute()
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "config file (config.yaml)")
	rootCmd.PersistentFlags().StringVarP(&runmode, "mode", "m", "auto", "mode: auto|upgrade")

	// 只支持长命令，默认为false，输入--print即为true BoolVarP可加短选项
	//rootCmd.PersistentFlags().BoolVar(&conf.FlagPrint, "print", false, "will print sth")
	rootCmd.PersistentFlags().BoolVarP(&conf.FlagPrint, "print", "p", false, "verbose output")
	// cmd.PersistentFlags().IntVarP(&port, "port", "p", 89, "port number")
	// cmd.PersistentFlags().DurationVarP(&timeout, "timeout", "t", 10*time.Second, "http request timeout")
	rootCmd.PersistentFlags().StringVarP(&conf.FlagOutputFile, "output", "o", "", "specify the output file name")

}

var yamlExample = []byte(
	`cmdserver:
  dbstr: helloooooo
  timeout:
    connect: 67s
    singleblock: 2s
  name:
    name: firstblood
`)

func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		viper.AddConfigPath("./")
		viper.SetConfigName("config")
		viper.SetConfigType("yaml")
	}

	viper.AutomaticEnv() // read in environment variables that match

	err := viper.ReadInConfig()
	if err != nil {
		klog.Println("not found config file. using default")

		// yamlStr := fmt.Sprintf("cmdserver:\n %s %s %s \n",
		// 						conf.FlagDBServer, conf.FlagTimeout, conf.FlagName)
		// ioutil.WriteFile("config.yaml", []byte(yamlStr), 0666)
		viper.ReadConfig(bytes.NewBuffer(yamlExample))
		viper.SafeWriteConfig()

	}
	conf.FlagDBServer = viper.GetString("cmdserver.dbstr")
	conf.FlagTimeout = viper.GetString("cmdserver.timeout.connect")
	conf.FlagName = viper.GetString("cmdserver.name.name")
	klog.Println(conf.FlagDBServer, conf.FlagTimeout, conf.FlagName)

	//设置监听回调函数 似乎调用了2次
	viper.OnConfigChange(func(e fsnotify.Event) {
		//klog.Printf("config is change :%s \n", e.String())
		conf.FlagTimeout = viper.GetString("cmdserver.timeout.connect")
	})

	viper.WatchConfig()

}
