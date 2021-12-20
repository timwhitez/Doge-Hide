package main

import (
    "fmt"
    "syscall"
    "time"
    "unsafe"
)

var h = HConsole()

func main() {
    fmt.Println("Hello Hide! ")
    time.Sleep(30*time.Second)
}

func HConsole()int{
    FindWindowA:= syscall.NewLazyDLL("user32.dll").NewProc("FindWindowA")
    lpClassName := "ConsoleWindowClass"
    Stealth,_,_ := FindWindowA.Call(uintptr(unsafe.Pointer(syscall.StringBytePtr(lpClassName))), 0)
    ShowWindow := syscall.NewLazyDLL("user32.dll").NewProc("ShowWindow")
    ShowWindow.Call(Stealth,0)
    return 0
}

