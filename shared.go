// @@@SNIPSTART hello-world-project-template-go-shared
package app

import (
	"os"

	"go.temporal.io/sdk/client"
)

const GreetingTaskQueue = "GREETING_TASK_QUEUE"

func ClientDial() (client.Client, error) {
	return client.Dial(client.Options{
		HostPort:  os.Getenv("TEMPORAL_HOST_PORT"),
		Namespace: os.Getenv("TEMPORAL_NAMESPACE"),
	})
}

// @@@SNIPEND
