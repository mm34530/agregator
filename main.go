package main

import (
	"agregator/producer"
	"agregator/publisher"
	"agregator/receiver"
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/msales/streams/v6"
)

func main() {
	producer := producer.New()
	publisher := publisher.New(producer)
	receiver := receiver.New()

	builder := streams.NewStreamBuilder()
	builder.Source("business-partners", publisher).
		Process("stats", receiver)

	tp, _ := builder.Build()
	task := streams.NewTask(tp)
	task.OnError(func(err error) {
		log.Fatal(err.Error())
	})

	task.Start(context.Background())
	defer task.Close()

	go generateMessages(context.Background(), publisher)
	go printStats(context.Background(), receiver)

	waitForShutdown()
}

func generateMessages(ctx context.Context, publisher *publisher.Publisher) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
			time.Sleep(time.Second)

			publisher.Publish(100)
		}
	}
}

func printStats(ctx context.Context, receiver *receiver.Receiver) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
			time.Sleep(time.Second)

			log.Println(receiver.GetStats())
		}
	}
}

func waitForShutdown() {
	quit := make(chan os.Signal, 1)

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	defer signal.Stop(quit)

	<-quit
}
