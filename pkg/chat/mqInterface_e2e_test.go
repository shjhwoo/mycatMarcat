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

// e2e test.
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
	testMQInf := MQInf{}

	err = testMQInf.Connect("amqp://guest:guest@" + endPoint)
	if err != nil {
		t.Error(err)
	}

	assert.NotNil(t, testMQInf.MQConn)

	err = testMQInf.CreateChannel()
	if err != nil {
		t.Error(err)
	}

	assert.NotNil(t, testMQInf.MQChannel)

	err = testMQInf.DeclareExchange("catExchange", "topic")
	if err != nil {
		t.Error(err)
	}

	q, err := testMQInf.DeclareQueue("user1")
	if err != nil {
		t.Error(err)
	}

	err = testMQInf.BindQueue((*q).Name, "user1.*.public", "catExchange")
	if err != nil {
		t.Error(err)
	}

	err = testMQInf.SendMsgToExchange("catExchange", "user1.vol.public", "user1님이 참여하셨습니다")
	if err != nil {
		t.Error(err)
	}

	//서버 시작 시 상담용 익스체인지, 봉사대화용 익스체인지 2개가 생성되어야 합니다
	//서버 시작 시 실행되는 함수를 테스트해야 한다
	//그 함수는 mq 커넥션 객체를 인자로 받아서 익스체인지를 생성할 수 있어야 한다
	//사용자가 로그인을 하게 되면 해당 사용자는 스톰프를 통해 특정 큐를 구독해야 합니다.
	//챗봇 상담
	//상담원 상담
	//봉사 구인, 봉사 신청 1:1채팅

	//사용자가 채팅 메세지를 입력하게 되면 해당 메세지는 바인딩 규칙을 통해서 익스체인지에 전달해야 합니다

	//누군가가 자신이 구독한 큐에 메세지를 보낸 경우 이를 컨슘할 수 있어야 합니다
}
