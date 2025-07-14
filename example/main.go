package main

import (
    "fmt"
    "net/http"

    "github.com/go-chi/chi/v5"
    "github.com/mmycin/intralb/config"
    "github.com/mmycin/intralb/balancer"
)

func main() {
    router := chi.NewRouter()
    router.Get("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Println("Handler triggered")
        _, err := w.Write([]byte("Hello World"))
        if err != nil {
            fmt.Println("Write error:", err)
        }
    })
    
    router.Get("/say", func(w http.ResponseWriter, r *http.Request) {
    	w.Write([]byte("Hello From Mycin"))
    })


    cfg := &config.Options{
        MaxConcurrentPerRouter: 100,
        QueueSize:              100,
        EnableLogging:          true,
        GracefulTimeoutSeconds: 5,
    }

    lb := balancer.New(router, cfg)
    server := &http.Server{
        Addr:    ":8080",
        Handler: lb.BalanceLoad(),
    }

    balancer.GracefulShutdown(server)

    fmt.Println("Serving on http://localhost:8080")
    server.ListenAndServe()
}
