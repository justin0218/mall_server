package auto_close_order

import (
	"fmt"
	"github.com/streadway/amqp"
	"mall_server/internal/services"
	"mall_server/store"
)

var rabbitmqService = new(store.Rabbitmq)
var orderService = new(services.OrderService)

func Do() {
	ch, err := rabbitmqService.Get().Channel()
	if err != nil {
		panic(err)
	}
	defer ch.Close()
	q, err := ch.QueueDeclare(
		"auto_close_order", // name
		false,              // durable
		false,              // delete when unused
		false,              // exclusive
		false,              // no-wait
		nil,                // arguments
	)
	if err != nil {
		panic(err)
	}

	err = ch.Qos(
		1,     // prefetch count
		0,     // prefetch size
		false, // global
	)

	if err != nil {
		panic(err)
	}

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		false,  // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	if err != nil {
		panic(err)
	}
	forever := make(chan bool)
	go func() {
		for d := range msgs {
			orderService.AutoCloseOrder(nil)
			err = ch.Publish(
				"",        // exchange
				d.ReplyTo, // routing key
				false,     // mandatory
				false,     // immediate
				amqp.Publishing{
					ContentType:   "text/plain",
					CorrelationId: d.CorrelationId,
					Body:          make([]byte, 0),
				})
			if err == nil {
				d.Ack(true)
			}
		}
	}()
	fmt.Println("RPC auto_close_order")
	<-forever
}
