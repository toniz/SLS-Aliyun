
package main

import (
    "fmt"
    "context"
    "time"

    "go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
    tracelog "github.com/toniz/otel"
)

func main() {
    err := tracelog.SetGrpcExport(context.TODO(), "trace_ali_bbthis.json", "OrderService", "v0.3.20")
    if err != nil {
        panic(err)
    }

    res, err := otelhttp.Head(context.TODO(), "http://127.0.0.1:8080/hello")

    fmt.Println(res)
    fmt.Println(err)
    res.Body.Close()

    time.Sleep(30*time.Second)
}

