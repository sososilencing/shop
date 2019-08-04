package service

import (
	"PowerShop/dao"
	"PowerShop/model"
	"encoding/json"
	"fmt"
	"github.com/streadway/amqp"
)

func failOnError(err error,msg string)  {
	if err!=nil{
		fmt.Println(err,msg)
	}
}

func Use(){
	fmt.Println("?")
	con,err:=amqp.Dial("amqp://guest:guest@localhost:5672/")

	failError(err,"amqp connection is failed")
	defer con.Close()

	ch,err:=con.Channel()
	failError(err,"channel is failed")

	que,err:=ch.QueueDeclare(
		"order",
		true,
		false,
		false,
		false,
		nil,
		)
	failError(err,"que is failed")

	ch.Qos(1,0,false)

	ever :=make(chan bool)

	msg,err:=ch.Consume(que.Name,"",false,false,false,false,nil)

	failError(err,"consume is failed")


	go func() {
		for thing := range msg{

			good:=&model.Order{}

			err:=json.Unmarshal(thing.Body,&good)


			if err != nil {
				fmt.Println(err.Error())
			}

			thing.Ack(false)
			//处理 业务 逻辑  存入数据库

			dao.Create(&good)

			dao.Update(good.ShopId,good.GoodId,good.Number)

			fmt.Println(good,"...........................................")
		}
	}()

	<-ever
}