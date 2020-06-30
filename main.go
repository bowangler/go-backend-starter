package main

import (
    "fmt"
    "go-backend-starter/router"
)

func main() {
    fmt.Println("This works")
    r := router.InitRouter()
    r.Run()
}
