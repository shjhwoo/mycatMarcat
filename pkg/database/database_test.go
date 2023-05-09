package database

import (
	"context"
	"path/filepath"
	"testing"

	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/mysql"
)

func TestDBInf(t *testing.T) {
	ctx := context.Background()
	dbContainer, err := mysql.RunContainer(ctx,
		mysql.WithScripts(filepath.Join("./createTable.sql")),
		mysql.WithDatabase("mycatmarcat"),
		mysql.WithUsername("root"),
		mysql.WithPassword("1234"),
		testcontainers.WithImage("mysql:8.0.33"),
	)
	if err != nil {
		t.Fatal(err)
	}

	defer func() {
		if err := dbContainer.Terminate(ctx); err != nil {
			t.Fatalf("failed to terminate container: %s", err.Error())
		}
	}()

}
