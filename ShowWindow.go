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
    //GetConsoleWindowName
    getWin := syscall.NewLazyDLL("kernel32.dll").NewProc("GetConsoleWindow")
    //ShowWindow
    showWin := syscall.NewLazyDLL("user32.dll").NewProc("ShowWindow")
    hwnd, _, _ := getWin.Call()
    if hwnd == 0 {
        return 1
    }
    showWin.Call(hwnd, 0)
    return 0
}