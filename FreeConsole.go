package main

import (
    "fmt"
    "syscall"
    "time"
)

var h = HConsole()

func main() {
    fmt.Println("Hello Hide! ")
    time.Sleep(30*time.Second)
}

func HConsole()int{
    //FreeConsole
    FreeConsole := syscall.NewLazyDLL("kernel32.dll").NewProc("FreeConsole")
    FreeConsole.Call()
    return 0
}