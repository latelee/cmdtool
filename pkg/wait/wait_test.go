package mywait

import (
    "fmt"
    "time"
    "testing"
)

// TODO：如何停止？
func TestWait(t *testing.T) {
    fmt.Println("testing wait...")
    go Until(func() {
        fmt.Println(".")
    }, 5 * time.Second, NeverStop)
    
    for {
        
    }
}