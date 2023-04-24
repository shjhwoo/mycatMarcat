package chat

import (
	"context"
	"fmt"
	"testing"

	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

// e2e test.
func TestMQInfra(t *testing.T) {
	ctx := context.Background()
	req := testcontainers.ContainerRequest{
		Image:        "rabbitmq:3.9.16-management",
		ExposedPorts: []string{"5672/tcp", "15674/tcp"},
		Files:        []testcontainers.ContainerFile{{HostFilePath: "./rabbit_enabled_plugins", ContainerFilePath: "/etc/rabbitmq/enabled_plugins", FileMode: 700}},
		WaitingFor:   wait.ForLog("plugins started."),
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

	//test functions.
	name, _ := mqContainer.Host(ctx)
	portMap, _ := mqContainer.Ports(ctx)
	fmt.Println(portMap["5672/tcp"][0].HostIP, portMap["5672/tcp"][0].HostPort, "이거머아")
	endPoint, _ := mqContainer.Endpoint(ctx, "amqp")
	fmt.Println(name, endPoint, "뭔대")

	//채팅 서버 시작 시 MQ에 연결하고, 단일 채널도 생성되어야 합니다
	// testMQInf := MQInf{}

	// err = testMQInf.Connect("amqp://guest:guest@" + name)
	// if err != nil {
	// 	t.Error(err)
	// }

	//익스체인지외 큐가 선언되어야 합니다

	//사용자가 로그인을 하게 되면 해당 사용자는 스톰프를 통해 특정 큐를 구독해야 합니다.

	//사용자가 채팅 메세지를 입력하게 되면 해당 메세지는 바인딩 규칙을 통해서 익스체인지에 전달해야 합니다

	//누군가가 자신이 구독한 큐에 메세지를 보낸 경우 이를 컨슘할 수 있어야 합니다

}
