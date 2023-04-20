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
		Image:        "rabbitmq:3.11.13",
		ExposedPorts: []string{"5672/tcp", "15672/tcp", "15674/tcp"},
		Env: map[string]string{
			"RABBITMQ_DEFAULT_USER": "admin",
			"RABBITMQ_DEFAULT_PASS": "admin",
		},
		Files:      []testcontainers.ContainerFile{{HostFilePath: "./rabbit_enabled_plugins", ContainerFilePath: "/etc/rabbitmq/enabled_plugins", FileMode: 700}},
		WaitingFor: wait.ForLog("Ready to accept connections"),
	}
	mqContainer, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	if err != nil {
		fmt.Println(err)
		t.Error(err)
	}
	defer func() {
		if err := mqContainer.Terminate(ctx); err != nil {
			t.Fatalf("failed to terminate container: %s", err.Error())
		}
	}()

	//test functions.

}
