package main

import (
    "fmt"
    "log"
    "net/http"
    "os"
)

func greetHandler(w http.ResponseWriter, r *http.Request) {
    message := "Hi there - Golang works just fine on Azure Functions Flex Consumption SKU"
    fmt.Fprint(w, message)
}

func main() {
    listenAddr := ":8080"
    if val, ok := os.LookupEnv("FUNCTIONS_CUSTOMHANDLER_PORT"); ok {
        listenAddr = ":" + val
    }
    
    http.HandleFunc("/api/GolangFunc", greetHandler)
    log.Printf("About to listen on %s. Go to https://127.0.0.1%s/", listenAddr, listenAddr)
    log.Fatal(http.ListenAndServe(listenAddr, nil))
}