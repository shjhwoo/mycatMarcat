package chat

import (
	"context"
	"fmt"
	"testing"

	"github.com/docker/go-connections/nat"
	"github.com/stretchr/testify/assert"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

func TestMQInfra(t *testing.T) {
	ctx := context.Background()
	req := testcontainers.ContainerRequest{
		Image:        "rabbitmq:3.9.16-management",
		ExposedPorts: []string{"5672/tcp", "15674/tcp"},
		/*
			When using testcontainers, it creates a temporary network to run the container in,
			and assigns a random available port to the container's exposed ports.
			This means that even though you requested port 5672 to be exposed in your container request,
			the actual port assigned to that container port might be different, in your case it seems to be 54604.
			To access your RabbitMQ instance from your code,
			you should use the port assigned to the container's exposed port
			rather than the original port you requested.
			In your case, you can use the port 54604 to connect to RabbitMQ instance:
		*/
		// HostConfigModifier: func(hc *container.HostConfig) {
		// 	hc.PortBindings = nat.PortMap{
		// 		"5672/tcp": []nat.PortBinding{
		// 			{
		// 				HostIP:   "0.0.0.0",
		// 				HostPort: "5672",
		// 			},
		// 		},
		// 	}
		// },
		Files:      []testcontainers.ContainerFile{{HostFilePath: "./rabbit_enabled_plugins", ContainerFilePath: "/etc/rabbitmq/enabled_plugins", FileMode: 700}},
		WaitingFor: wait.ForLog("plugins started."),
		Name:       "myTestRabbitMQ",
	}
	mqContainer, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})

	if err != nil {
		t.Error(err)
	}

	defer func() {
		if err := mqContainer.Terminate(ctx); err != nil {
			t.Fatalf("failed to terminate container: %s", err.Error())
		}
	}()

	//test functions. 인터페이스 함수에 대해서만 테스트하기
	endPoint, err := mqContainer.PortEndpoint(ctx, nat.Port("5672"), "")
	if err != nil {
		fmt.Println(err, "에러 빌생")
		return
	}

	//채팅 서버 시작 시 MQ에 연결하고, 단일 채널도 생성되어야 합니다
	err = MQIn.Connect("amqp://guest:guest@" + endPoint)
	if err != nil {
		t.Error(err)
	}

	assert.NotNil(t, MQIn.MQConn)

	err = MQIn.CreateChannel()
	if err != nil {
		t.Error(err)
	}

	assert.NotNil(t, MQIn.MQChannel)

	err = MQIn.DeclareExchange("catExchange", "topic")
	if err != nil {
		t.Error(err)
	}

	q, err := MQIn.DeclareQueue("user1")
	if err != nil {
		t.Error(err)
	}

	err = MQIn.BindQueue((*q).Name, "user1.*.public", "catExchange")
	if err != nil {
		t.Error(err)
	}

	err = MQIn.SendMsgToExchange("catExchange", "user1.vol.public", "user1님이 참여하셨습니다")
	if err != nil {
		t.Error(err)
	}

	err = MQIn.UnbindQueue((*q).Name, "user1.*.public", "catExchange")
	if err != nil {
		t.Error(err)
	}

	//사용지가 채팅방 입장 시 함수 구성요소만 테스트. 통합테스트아님
	chat := setChatInfo("chatbot")

	assert.Equal(t, chat.getExchangeName(), "helperEx")
	assert.Equal(t, chat.getExchangeType(), "topic")
	assert.Equal(t, chat.getRoutingKeyForBind(), "bot.*")

	chat = setChatInfo("consult")

	assert.Equal(t, chat.getExchangeName(), "helperEx")
	assert.Equal(t, chat.getExchangeType(), "topic")
	assert.Equal(t, chat.getRoutingKeyForBind(), "consult.*")

	chat = setChatInfo("volunteerPrivate")

	assert.Equal(t, chat.getExchangeName(), "volunteerEx")
	assert.Equal(t, chat.getExchangeType(), "topic")
	assert.Equal(t, chat.getRoutingKeyForBind(), "vol.private.*")

	chat = setChatInfo("volunteerPublic")

	assert.Equal(t, chat.getExchangeName(), "volunteerEx")
	assert.Equal(t, chat.getExchangeType(), "topic")
	assert.Equal(t, chat.getRoutingKeyForBind(), "vol.public.*")

	//====채팅방 들어가기, 나가기 테스트
	err = MQIn.DeclareExchange("helperEx", "topic")
	if err != nil {
		t.Error(err)
	}

	err = EnterNewChatRoom("user1", "chatbot")
	if err != nil {
		t.Error(err)
	}

	err = SendMessage("chatbot", "user1", "expected delivery date will be 5/12")
	if err != nil {
		t.Error(err)
	}

	err = LeaveChatRoom("user1", "chatbot")
	if err != nil {
		t.Error(err)
	}

	err = checkEmptyUserId("")
	assert.Equal(t, err.Error(), "userid cannot be empty")
}
