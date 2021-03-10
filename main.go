package main

import (
    "fmt"
    "html"
    "log"
    "net/http"
    "os"
)

func main() {

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        _, _ = fmt.Fprintf(w, "Hello, %q!!", html.EscapeString(r.URL.Path))
    })

    http.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request){
        _, _ = fmt.Fprintf(w, "Hi")
    })

    port := os.Getenv("PORT")
    if port == "" {
        port = "8081"
        log.Printf("Defaulting to port %s", port)
    }
    log.Fatal(http.ListenAndServe(":" + port, nil))

}

func Calculate(x int) (result int) {
    result = x + 2
    return result
}