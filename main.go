package main


import (
    "./res"
    "fmt"
    "os"
    "os/exec"
)

func WriteResBin(fn, bin string) {
    f, err := os.Create(fn)
    if (err != nil) {
        fmt.Println(err.Error())
        return
    }
    defer f.Close()

    resBin, err := res.Asset(bin)
    if (err != nil) {
        fmt.Println(err.Error())
        return
    }
    len, err := f.Write(resBin)
    if (err != nil) {
        fmt.Println(err.Error())
        return
    }
    fmt.Printf("Write %d bytes\n", len)
}

func main() {
    WriteResBin("HelloWorld.exe", "HelloWorld.exe")
    out, err := exec.Command("HelloWorld.exe", "").Output()
    if (err != nil) {
        fmt.Println(err.Error())
        os.Exit(0)
    }
    fmt.Println(fmt.Sprintf("%s", out))
}
