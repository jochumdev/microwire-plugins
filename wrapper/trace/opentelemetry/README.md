# OpenTelemetry wrappers

OpenTelemetry wrappers propagate traces (spans) accross services.

## Usage

```go
service, _ := micro.NewService(
    micro.Name("go.micro.srv.greeter"),
    micro.WrapClient(opentelemetry.NewClientWrapper()),
    micro.WrapHandler(open.NewHandlerWrapper()),
    micro.WrapSubscriber(opentelemetry.NewSubscriberWrapper()),
)
```