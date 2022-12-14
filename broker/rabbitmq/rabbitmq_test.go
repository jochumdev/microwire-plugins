package rabbitmq_test

import (
	"context"
	"os"
	"testing"

	rabbitmq "github.com/go-micro/microwire-plugins/broker/rabbitmq/v5"
	micro "github.com/go-micro/microwire/v5"
	broker "github.com/go-micro/microwire/v5/broker"
	server "github.com/go-micro/microwire/v5/server"
)

type Example struct{}

func (e *Example) Handler(ctx context.Context, r interface{}) error {
	return nil
}

func TestDurable(t *testing.T) {
	if tr := os.Getenv("TRAVIS"); len(tr) > 0 {
		t.Skip()
	}
	rabbitmq.DefaultRabbitURL = "amqp://rabbitmq:rabbitmq@127.0.0.1:5672"
	brkrSub := broker.NewSubscribeOptions(
		broker.Queue("queue.default"),
		broker.DisableAutoAck(),
		rabbitmq.DurableQueue(),
	)

	b := rabbitmq.NewBroker()
	b.Init()
	if err := b.Connect(); err != nil {
		t.Logf("cant conect to broker, skip: %v", err)
		t.Skip()
	}

	s := server.NewServer(server.Broker(b))

	service, err := micro.NewService(
		micro.Server(s),
		micro.Broker(b),
	)
	if err != nil {
		t.Fatal(err)
	}

	h := &Example{}
	// Register a subscriber
	micro.RegisterSubscriber(
		"topic",
		service.Server(),
		h.Handler,
		server.SubscriberContext(brkrSub.Context),
		server.SubscriberQueue("queue.default"),
	)

	//service.Init()

	if err := service.Run(); err != nil {
		t.Fatal(err)
	}
}
