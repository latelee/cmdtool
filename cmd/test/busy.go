package cmd

import (
	//"fmt"
	
	"k8s.io/klog"
	conf "cmdtool/common/conf"
	com "cmdtool/pkg/com"
)

func foo(args []string) {
	klog.Println("test foo.....");
}

// 监听配置参数变化
func testWatch(args []string) {
	timeout := conf.FlagTimeout
	for {
		if timeout != conf.FlagTimeout {
			klog.Printf("param changed: %v\n", conf.FlagTimeout)
			timeout = conf.FlagTimeout
		}
		com.Sleep(1000)
	}
}