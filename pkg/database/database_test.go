package database

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

func TestDBInf(t *testing.T) {
	ctx := context.Background()

	// Start a MySQL container
	mysqlContainer, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: testcontainers.ContainerRequest{
			Image:        "mysql:8.0.33",
			ExposedPorts: []string{"3306/tcp"},
			Env: map[string]string{
				"MYSQL_ROOT_PASSWORD": "1234",
				"MYSQL_DATABASE":      "mycatmarcat",
			},
			WaitingFor: wait.ForLog("port: 3306  MySQL Community Server - GPL"),
		},
		Started: true,
	})
	if err != nil {
		t.Fatal(err)
	}

	defer func() {
		// Terminate and remove the container
		if err := mysqlContainer.Terminate(ctx); err != nil {
			log.Fatalf("failed to terminate container: %s", err.Error())
		}
	}()

	// Get the host and port for connecting to the MySQL container
	mysqlHost, err := mysqlContainer.Host(ctx)
	if err != nil {
		t.Fatal(err)
	}

	mysqlPort, err := mysqlContainer.MappedPort(ctx, "3306")
	if err != nil {
		t.Fatal(err)
	}

	// Create the MySQL connection string
	dsn := fmt.Sprintf("root:1234@tcp(%s:%s)/mycatmarcat", mysqlHost, mysqlPort.Port())

	// Connect to the MySQL container
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	// Perform tests against the MySQL container here...
}
