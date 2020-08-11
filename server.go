import (
    "fmt"
    "log"
    "net/http"
)

func handleHelloRequest(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello from GitLab!")
}

func runServer(port string) {
    http.HandleFunc("/hello", handleHelloRequest)

    fmt.Printf("Starting server")

    if err := http.ListenAndServe(":" + port, nil); err != nil {
        log.Fatal(err)
    }
}
