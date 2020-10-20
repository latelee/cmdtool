
package com

import (
	"fmt"
	"time"
	"testing"
)

func TestTime(t *testing.T) {
	fmt.Println("time test: ", DateT(time.Now(), "YYYY-MM-DD HH:00:00"))
	fmt.Println("time test: ", DateT(time.Now(), "YYYY-MM-DD HH:mm:ss"))
	Sleep(100)
	
}
