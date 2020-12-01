package cmd

import (
	//"fmt"
	
	"k8s.io/klog"

)

func foo(args []string) {
	klog.Println("misc foo.....");
	for _, item:=range args {
        klog.Println("args: ", item)
    }
}