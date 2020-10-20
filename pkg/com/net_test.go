
package com

import (
	"fmt"
	"testing"
)

func TestIP(t *testing.T) {
    ip :=  GetLocalIp()
    fmt.Println("localip: ", ip)
}