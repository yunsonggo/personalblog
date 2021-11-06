package common

import (
	"2021/yunsongcailu/yunsong_server/config"
	"fmt"
	"github.com/streadway/amqp"
	"github.com/yakaa/log4g"
)

// 连接地址
var mqAddr = config.Conf.Rabbitmq.RabbitmqAddr

type RabbitMQ struct {
	conn    *amqp.Connection
	channel *amqp.Channel
	// 队列
	QueueName string
	// 交换机
	Exchange string
	// bind key
	Key string
	// conn
	Mqurl string
}

func NewRabbitMQ(queueName string, exchange string, key string) *RabbitMQ {
	return &RabbitMQ{
		QueueName: queueName,
		Exchange:  exchange,
		Key:       key,
		Mqurl:     mqAddr,
	}
}

// 断开 关闭
func (r *RabbitMQ) Destory() {
	_ = r.channel.Close()
	_ = r.conn.Close()
}

// 错误处理
func (r *RabbitMQ) failOnErr(err error, message string) {
	if err != nil {
		log4g.Error("%s:%s", message, err)
		panic(fmt.Sprintf("%s:%s", message, err))
	}
}

// 简单模式下  实例
func NewRabbitMQSimple(queueName string) *RabbitMQ {
	rabbitMQ := NewRabbitMQ(queueName, "", "")
	var err error
	// 拨号
	rabbitMQ.conn, err = amqp.Dial(rabbitMQ.Mqurl)
	rabbitMQ.failOnErr(err, "简单模式消息队列:创建连接失败")
	// 获取channel
	rabbitMQ.channel, err = rabbitMQ.conn.Channel()
	rabbitMQ.failOnErr(err, "简单模式消息队列:打开通道失败")
	return rabbitMQ
}

// 生成队列
func (r *RabbitMQ) PublishSimple(message string) {
	// 申请队列
	_, err := r.channel.QueueDeclare(
		r.QueueName,
		// 是否持久化
		false,
		//是否自动删除
		false,
		//是否具有排他性
		false,
		//是否阻塞处理
		false,
		//额外的属性
		nil,
	)
	if err != nil {
		log4g.Error(err)
		fmt.Println(err)
		return
	}
	// 调用channel 发送消息到队列中
	_ = r.channel.Publish(
		r.Exchange,
		r.QueueName,
		//如果为true，根据自身exchange类型和routekey规则无法找到符合条件的队列会把消息返还给发送者
		false,
		//如果为true，当exchange发送消息到队列后发现队列上没有消费者，则会把消息返还给发送者
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		})
}
// 消费队列
func (r *RabbitMQ) ConsumerSimple() {
	// 申请队列
	q, err := r.channel.QueueDeclare(
		r.QueueName,
		//是否持久化
		false,
		//是否自动删除
		false,
		//是否具有排他性
		false,
		//是否阻塞处理
		false,
		//额外的属性
		nil,
	)
	if err != nil {
		log4g.Error(err)
		fmt.Println(err)
		return
	}
	// 接收消息
	msgs, err := r.channel.Consume(
		q.Name, // queue
		//用来区分多个消费者
		"", // consumer
		//是否自动应答
		false, // auto-ack
		//是否独有
		false, // exclusive
		//设置为true，表示 不能将同一个Conenction中生产者发送的消息传递给这个Connection中 的消费者
		false, // no-local
		//列是否阻塞
		false, // no-wait
		nil,   // args
	)
	if err != nil {
		log4g.Error(err)
		fmt.Println(err)
	}
	// 配合Qos 消费者流控 防止暴库
	_ = r.channel.Qos(
		1,     // 每次消费一个消息
		0,     // 服务器传递的最大容量 八位字节为单位
		false, // false 当前消息队列 true 全局使用
	)
	forever := make(chan bool)
	// 启用协程处理消息
	go func() {
		for d := range msgs {
			// 消息逻辑处理
			log4g.Error("处理消息:%s", d.Body)
			// 配合 autoAck 为false 使用 autoAck 为false 如果不写 d.Ack(false) 后果非常严重 会重新返回消息队列 等待消费
			// 如果d.Ack(true) 一般批量消息处理 表示确认所有未确认的消息
			// 如果d.Ack(false) 表示确认当前消息
			_ = d.Ack(false)
		}
	}()
	log4g.Info(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
