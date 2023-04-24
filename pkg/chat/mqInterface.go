package chat

import (
	"context"
	"errors"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

// MQ와 연결하고 채널 생성, exchange, 큐 선언, 바인딩, 메세지 발행, 큐 삭제 등의 함수를 작성합니다.
type MQInf struct {
	MQConn    *amqp.Connection
	MQChannel *amqp.Channel
}

func (mc *MQInf) Connect(url string) error {
	conn, err := amqp.Dial(url)
	if err != nil {
		return err
	}
	defer conn.Close()
	mc.MQConn = conn
	return nil
}

// 서로 다른 채널은 메세지 공유가 불가합니다!!
func (mc *MQInf) CreateChannel() error {
	if mc.MQConn == nil {
		return errors.New("cannot create a channel with nil Connection")
	}

	ch, err := mc.MQConn.Channel() //커넥션이 닫히면 이 채널도 닫힌다.
	if err != nil {
		return err
	}

	mc.MQChannel = ch
	return nil
}

func (mc *MQInf) DeclareExchange() error {
	if mc.MQChannel == nil {
		return errors.New("cannot declare exchange with nil channel")
	}

	err := mc.MQChannel.ExchangeDeclare(
		"name",  //name
		"topic", // type
		true,    //durable
		false,   //auto-deleteds
		false,   //internal
		false,   //no-wait?
		nil,     //arguments?
	)
	if err != nil {
		return err
	}
	return nil
}

func (mc *MQInf) DeclareQueue(queueName string) error {
	if mc.MQChannel == nil {
		return errors.New("cannot declare queue with nil channel")
	}

	_, err := mc.MQChannel.QueueDeclare(
		queueName, // names
		true,      // durable
		true,      // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	if err != nil {
		return err
	}
	return nil
}

func (mc *MQInf) BindQueue() error {
	if mc.MQChannel == nil {
		return errors.New("cannot bind queue with nil channel")
	}

	err := mc.MQChannel.QueueBind(
		"queueName",
		"routingKey",
		"exchangeName",
		false,
		nil,
	)
	if err != nil {
		return err
	}
	return nil
}

func (mc *MQInf) SendMsgToExchange(exchange, routingKey, message string, ch *amqp.Channel) error {
	//messageType: public, private, disconnect
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if mc.MQChannel == nil {
		return errors.New("cannot publish message with nil channel")
	}

	err := mc.MQChannel.PublishWithContext(ctx,
		exchange,   // exchange
		routingKey, // routing key
		false,      // mandatory
		false,      // immediate
		amqp.Publishing{
			ContentType: "application/json", //이걸로 바꿈
			Body:        []byte(message),
		})
	if err != nil {
		return err
	}

	return nil
}
