package consume_wx_pay_notice

import (
	"encoding/json"
	"mall_server/internal/models/wx"
	"mall_server/internal/services"
	"mall_server/store"
)

var rabbitmqService = new(store.Rabbitmq)
var orderService = new(services.OrderService)

func Do() (err error) {
	mq := rabbitmqService.Get()
	mqChan, e := mq.Channel()
	if e != nil {
		err = e
		return
	}
	defer mqChan.Close()
	q, e := mqChan.QueueDeclare(
		"wxPayNotice", // name
		false,         // durable
		false,         // delete when unused
		false,         // exclusive
		false,         // no-wait
		nil,           // arguments
	)
	if e != nil {
		err = e
		return
	}
	msgs, err := mqChan.Consume(
		q.Name, // queue
		"",     // consumer
		false,  // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	forever := make(chan bool)
	go func() {
		for d := range msgs {
			wxPayReq := new(wx.WxpayReq)
			err := json.Unmarshal(d.Body, wxPayReq)
			if err != nil {
				continue
			}
			err = orderService.ConfirmOrder(wxPayReq)
			if err == nil {
				_ = d.Ack(true)
			}
		}
	}()
	<-forever
	return
}
