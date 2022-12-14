# JSON-RPC 2.0 Codec

## Usage

Import the codec and set within the client/server
```go
package main

import (
    "github.com/go-micro/microwire/v5"
    "github.com/micro/go-micro/client"
    "github.com/micro/go-micro/server"
    "github.com/go-micro/microwire-plugins/codec/jsonrpc2/v5"
)

func main() {
    client := client.NewClient(
        client.Codec("application/json", jsonrpc2.NewCodec),
        client.ContentType("application/json"),
    )

    server := server.NewServer(
        server.Codec("application/json", jsonrpc2.NewCodec),
    )

    service, _ := micro.NewService(
        micro.Client(client),
        micro.Server(server),
    )

    // ...
}
```

