package main

import (
    "fmt"
    "syscall"
    "time"
)

var h = HConsole()

func main() {
    fmt.Println("Hello Hide! ")
    time.Sleep(10*time.Second)
}

func HConsole()int{
    getWin := syscall.NewLazyDLL("kernel32.dll").NewProc("GetConsoleWindow")
    hwnd,_,_ := getWin.Call()
    if hwnd != 0 {
        showWindowAsync := syscall.NewLazyDLL("user32.dll").NewProc("ShowWindowAsync")
        showWindowAsync.Call(hwnd,0)
    }
    return 0
}
