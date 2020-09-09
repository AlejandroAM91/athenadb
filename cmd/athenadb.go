package main

import (
    "os"
    "os/signal"
    "syscall"

    "github.com/AlejandroAM91/athenadb/internal/app/athenadb/ui/api"
)

func main() {
    api := api.NewAPI()
    api.Start(":8080")

    ch := make(chan os.Signal)
    signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
    <-ch

    api.Stop()
}
