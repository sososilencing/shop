package service

import (
	"encoding/json"
	"fmt"
	"github.com/streadway/amqp"
	"PowerShop/model"
)

func failError(err error,msg string){
	if err!=nil{
		fmt.Println(msg,err)
	}
}

func Produce(order model.Order){
	fmt.Println(order)
	con,err:=amqp.Dial("amqp://guest:guest@127.0.0.1:5672/")
	if err != nil {
		fmt.Println(err.Error())
	}
	defer con.Close()
	/**
	*仅仅只有拿来建立管道
	 */
	ch,err:=con.Channel()

	defer ch.Close()

	//得到的ch 管道 才是 可以用来搞事

	que,err:=ch.QueueDeclare(
		"order",
		true,
		false,
		false,
		false,
		nil,
	)

	body,_:=json.Marshal(order)

	err = ch.Publish("", que.Name, false, false, amqp.Publishing{
		DeliveryMode: amqp.Persistent,
		ContentType:  "text/plain",
		Body:         body,
	})

	if err != nil {
		failError(err,"The publish is failed")
	}
}