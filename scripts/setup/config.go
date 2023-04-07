package setup

import (
	"fmt"
	"strconv"
)

const GolangCiLintVersion = "1.51.2"

const GolangCiLintFolder = "third_party/golangci-lint"

const GolangDynamoDBLocalFolder = "third_party/dynamodb-local"
const GolangDynamoDBLocalJar = "scripts/third_party/dynamodb-local"

const DynamoDBPort = 8184

type ServerProcess struct {
	Command        CommandInBackground
	HealthEndpoint string
	Name           string
	Port           uint32
}

func BackendConfig() ServerProcess {
	return ServerProcess{
		Command: CommandInBackground{
			Command: "go",
			Dir:     "backend/",
			Args:    []string{"run", "../backend/local/main.go"},
		},
		Name:           "backend",
		Port:           8080,
		HealthEndpoint: "http://localhost:8080/api/health",
	}
}

func DynamodbConfig() ServerProcess {
	return ServerProcess{
		Command: CommandInBackground{
			Command: "java",
			Args:    []string{"-Djava.library.path=./DynamoDBLocal_lib", "-jar", "DynamoDBLocal.jar", "-inMemory", "-sharedDb", "-port", strconv.Itoa(DynamoDBPort)},
			Dir:     GolangDynamoDBLocalJar,
		},
		Name:           "DynamoDB",
		Port:           DynamoDBPort,
		HealthEndpoint: fmt.Sprintf("http://localhost:%d", DynamoDBPort),
	}
}

func FrontendConfig() ServerProcess {
	return ServerProcess{
		Command: CommandInBackground{
			Command: "npm",
			Dir:     "frontend/",
			Args:    []string{"run", "start"},
		},
		Name:           "frontend",
		Port:           3000,
		HealthEndpoint: "http://localhost:3000/",
	}
}
