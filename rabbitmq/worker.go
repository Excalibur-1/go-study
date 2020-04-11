package main

import (
	"bytes"
	"github.com/streadway/amqp"
	"log"
	"time"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func main() {
	//获取连接
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	//获取管道
	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	//设置队列
	/*q, err := ch.QueueDeclare(
		"hello",
		false,
		false,
		false,
		false,
		nil,
	)*/
	//不允许使用不同的参数定义已有队列
	/*q, err := ch.QueueDeclare(
		"hello",
		true,//设置队列为持久化的
		false,
		false,
		false,
		nil,
	)*/
	//声明名称不同的队列，生产者也需要对应修改
	q, err := ch.QueueDeclare(
		"task_queue",
		true, //设置队列为持久化的
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "Failed to declare a queue")

	//设置预计取数为1，控制不要给消费者发送一个以上的消息
	err = ch.Qos(
		1,
		0,
		false,
	)
	failOnError(err, "Failed to set Qos")

	//消费者
	/*msgs, err := ch.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)*/
	//通过设置autoAck为false来实现消息确认功能
	msgs, err := ch.Consume(
		q.Name,
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "Failed to register a consumer")

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf("Reveived a message: %s", d.Body)
			dotCount := bytes.Count(d.Body, []byte("."))
			t := time.Duration(dotCount)
			time.Sleep(t * time.Second)
			log.Printf("Done")
			//进行消息确认
			d.Ack(false)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")

	<-forever
}
