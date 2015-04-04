package main

import (
    "fmt"
    "log"
    "os"

    "github.com/arodland/unofficialnest"
)

func main() {
    if len(os.Args) != 3 {
        log.Fatal("Usage: %s username password", os.Args[0])
    }
    username, password := os.Args[1], os.Args[2]

    nest := unofficialnest.NewSession(unofficialnest.Credentials{username, password})
    status, err := nest.GetStatus()
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println(status)
}
