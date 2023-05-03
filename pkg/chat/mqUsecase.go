package chat

import (
	"errors"
	"fmt"
)

//공통 메서드: 사용자 입장, 사용자 나감
//특이적 메서드

//exchange 2개 생성
//상담용 exchange는 상담원, 챗봇, 봉사용 exchange
//사용자가 채팅방에 입장할 때부터, 큐가 생성됩니다. 채팅 입장 전에는 구독 상태가 아닙니다.

type Chat interface {
	getExchangeName() string
	getExchangeType() string
	getRoutingKeyForBind() string
	makeRouteKey(string) string
}

var HelperExchangeName = "helperEx"
var VolunteerExchangeName = "volunteerEx"

type HelperChat struct {
	pairType string //bot or consultant
}

func (hc *HelperChat) getExchangeName() string {
	return HelperExchangeName
}

func (hc *HelperChat) getExchangeType() string {
	return "topic"
}

func (hc *HelperChat) getRoutingKeyForBind() string {
	return fmt.Sprintf("%s.*", hc.pairType)
}

func (hc *HelperChat) makeRouteKey(to string) string {
	return fmt.Sprintf("%s.%s", hc.pairType, to)
}

type VolunteerChat struct {
	privacy string //public provate
}

func (vc *VolunteerChat) getExchangeName() string {
	return VolunteerExchangeName
}

func (vc *VolunteerChat) getExchangeType() string {
	return "topic"
}

func (vc *VolunteerChat) getRoutingKeyForBind() string {
	return fmt.Sprintf("vol.%s.*", vc.privacy)
}

func (vc *VolunteerChat) makeRouteKey(to string) string {
	return fmt.Sprintf("vol.%s.%s", vc.privacy, to)
}

// 팩토리 함수
func setChatInfo(chatType string) Chat {
	var chat Chat

	switch chatType {
	case "chatbot":
		chat = &HelperChat{pairType: "bot"}
	case "consult":
		chat = &HelperChat{pairType: "consult"}
	case "volunteerPrivate":
		chat = &VolunteerChat{privacy: "private"}
	case "volunteerPublic":
		chat = &VolunteerChat{privacy: "public"}
	}

	return chat
}

// 사용자 채팅방에 입장합니다
func EnterNewChatRoom(userid, chatType string) error {
	err := checkEmptyUserId(userid)
	if err != nil {
		return err
	}

	q, err := MQIn.DeclareQueue(userid)
	if err != nil {
		return err
	}

	chat := setChatInfo(chatType)

	routingKey := chat.getRoutingKeyForBind()
	exchangeName := chat.getExchangeName()

	err = MQIn.BindQueue(q.Name, routingKey, exchangeName)
	if err != nil {
		return err
	}
	return nil
}

// 사용자가 채팅방을 나갑니다
func LeaveChatRoom(userid, chatType string) error {
	err := checkEmptyUserId(userid)
	if err != nil {
		return err
	}

	chat := setChatInfo(chatType)

	queueName := userid
	routingKey := chat.getRoutingKeyForBind()
	exchangeName := chat.getExchangeName()

	err = MQIn.UnbindQueue(queueName, routingKey, exchangeName)
	if err != nil {
		return err
	}
	return nil
}

func checkEmptyUserId(userid string) error {
	if userid == "" {
		return errors.New("userid cannot be empty")
	}
	return nil
}

func SendMessage(chatType, to, message string) error {
	chat := setChatInfo(chatType)

	exchange := chat.getExchangeName()
	routingKey := chat.makeRouteKey(to)

	err := MQIn.SendMsgToExchange(exchange, routingKey, message)
	if err != nil {
		return err
	}

	return nil
}
