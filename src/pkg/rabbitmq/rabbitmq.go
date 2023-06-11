package rabbitmq

import (
	"context"
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
	"time"
)

type RMQ struct {
	Username     string
	Password     string
	Host         string
	Port         string
	ExchangeName string
	QueueName    string
	RoutingKey   string
	RMQChannel   *amqp.Channel
}

func NewRMQ(username, password, host, port, exName, qName, rKey string) *RMQ {
	return &RMQ{
		Username:     username,
		Password:     password,
		Host:         host,
		Port:         port,
		ExchangeName: exName,
		QueueName:    qName,
		RoutingKey:   rKey,
	}
}

func (r *RMQ) getRMQConnection() *amqp.Connection {
	uri := fmt.Sprintf("amqp://%s:%s@%s:%s/", r.Username, r.Password, r.Host, r.Port)
	RMQChannel, err := amqp.Dial(uri)
	if err != nil {
		fmt.Println("Failed to connect to RabbitMQ")
	}
	return RMQChannel
}

func (r *RMQ) Initialize() {
	RMQChannel := r.getRMQConnection()
	defer RMQChannel.Close()

	ch, err := RMQChannel.Channel()
	if err != nil {
		fmt.Println("Failed to connect to RabbitMQ")
	}
	defer ch.Close()

	err = ch.ExchangeDeclare(
		r.ExchangeName, // name
		"direct",       // type
		true,           // durable
		false,          // auto-deleted
		false,          // internal
		false,          // no-wait
		nil,            // arguments
	)
	if err != nil {
		fmt.Println("Failed to declare an exchange")
	}

	_, err = ch.QueueDeclare(
		r.QueueName, // name
		false,       // durable
		false,       // delete when unused
		false,        // exclusive
		false,       // no-wait
		nil,         // arguments
	)
	if err != nil {
		fmt.Println("Failed to declare a queue")
	}

	err = ch.QueueBind(
		r.QueueName,    // queue name
		r.RoutingKey,   // routing key
		r.ExchangeName, // exchange
		false,
		nil,
	)
	if err != nil {
		fmt.Println("Failed to bind a queue")
	}
	var forever chan struct{}
	<-forever
}

func (r *RMQ) Send(msg []byte) {
	RMQChannel := r.getRMQConnection()
	defer RMQChannel.Close()

	ch, err := RMQChannel.Channel()
	if err != nil {
		fmt.Println("Failed to connect to RabbitMQ")
	}
	defer ch.Close()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = ch.PublishWithContext(ctx,
		r.ExchangeName, // exchange
		r.RoutingKey,   // routing key
		false,          // mandatory
		false,          // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        msg,
		})

	if err != nil {
		fmt.Println("Failed to Send msg in queue ", r.QueueName)
	}
	//fmt.Printf(" [x] Sent %s\n", msg)
}

func (r *RMQ) Read() {
	RMQChannel := r.getRMQConnection()
	defer RMQChannel.Close()

	ch, err := RMQChannel.Channel()
	if err != nil {
		fmt.Println("Failed to connect to RabbitMQ")
	}
	defer ch.Close()

	msgs, err := ch.Consume(
		r.QueueName,  // queue
		r.RoutingKey, // consumer
		true,         // auto ack
		false,        // exclusive
		false,        // no local
		false,        // no wait
		nil,          // args
	)
	if err != nil {
		fmt.Println("Failed to register a consumer for queue ", r.QueueName, ": ",err)
	}

	var forever chan struct{}

	go func() {
		for d := range msgs {
			fmt.Printf(" [x] Read %s", d.Body)
		}
	}()

	fmt.Printf(" [*] Waiting for logs. To exit press CTRL+C")
	<-forever

}

func (r *RMQ) GetMessage(c *chan []byte) {
	RMQChannel := r.getRMQConnection()
	defer RMQChannel.Close()

	ch, err := RMQChannel.Channel()
	if err != nil {
		fmt.Println("Failed to connect to RabbitMQ")
	}
	defer ch.Close()

	msgs, err := ch.Consume(
		r.QueueName,  // queue
		r.RoutingKey, // consumer
		true,         // auto ack
		false,        // exclusive
		false,        // no local
		false,        // no wait
		nil,          // args
	)
	if err != nil {
		fmt.Println("Failed to register a consumer for queue ", r.QueueName, ": ",err)
	}

	var forever chan struct{}

	go func() {
		for d := range msgs {
			//fmt.Printf(" [x] Sending %s through channel", d.Body)
			*c <- d.Body
		}
	}()
	<-forever

}
