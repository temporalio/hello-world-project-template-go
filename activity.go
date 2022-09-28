// @@@SNIPSTART hello-world-project-template-go-activity
package app

import (
	"fmt"

	"go.temporal.io/sdk/workflow"
)

func ComposeGreeting(ctx workflow.Context, name string) (string, error) {
	greeting := fmt.Sprintf("Hello %s!", name)
	return greeting, nil
}

// @@@SNIPEND
