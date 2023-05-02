package chat

import (
	"errors"
	"fmt"
)

//공통 메서드: 사용자 입장, 사용자 나감
//특이적 메서드

//exchange 2개 생성
//상담용 exchange는 상담원, 챗봇, 봉사용 exchange

//사용자가 로그인을 하면 기본 큐가 생성됩니다. 채팅 입장 전에는 구독 상태가 아닙니다.

type Chat interface {
	getExchangeName() string
	getExchangeType() string
	getRoutingKey() string
}

var chat Chat

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

func (hc *HelperChat) getRoutingKey() string {
	return fmt.Sprintf("%s.*", hc.pairType)
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

func (vc *VolunteerChat) getRoutingKey() string {
	return fmt.Sprintf("vol.%s.*", vc.privacy)
}

// 팩토리 함수
func setChatInfo(chatType string) {
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
}

func CreateUserMsgBox(userid string) error {
	err := checkEmptyUserId(userid)
	if err != nil {
		return err
	}

	_, err = MQIn.DeclareQueue(userid)
	if err != nil {
		return err
	}

	return nil
}

// 사용자 채팅방에 입장합니다
func EnterNewChatRoom(userid, chatType string) error {
	err := checkEmptyUserId(userid)
	if err != nil {
		return err
	}
	setChatInfo(chatType)

	queueName := userid
	routingKey := chat.getRoutingKey()
	exchangeName := chat.getExchangeName()

	err = MQIn.BindQueue(queueName, routingKey, exchangeName)
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

	setChatInfo(chatType)

	queueName := userid
	routingKey := chat.getRoutingKey()
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
