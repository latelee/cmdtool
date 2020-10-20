/*
网络相关接口

注：获取IP地址存在bug：当有多个IP时，只返回最后一个。理论上获取活动的那一个。
*/

package com

import (
    "fmt"
	"net"
)

// 取得本机IP地址
func GetLocalIp() (ip string) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}
	for _, value := range addrs {
        fmt.Println("value: ", value)
		if ipnet, ok := value.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				ip = ipnet.IP.String()
			}
		}
	}

	return ip
}