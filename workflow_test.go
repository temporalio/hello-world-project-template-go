// @@@SNIPSTART hello-world-project-template-go-workflow-test
package app

import (
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"go.temporal.io/sdk/testsuite"
)

func Test_Workflow(t *testing.T) {
	// Set up the test suite and testing execution environment
	testSuite := &testsuite.WorkflowTestSuite{}
	env := testSuite.NewTestWorkflowEnvironment()

	// Mock activity implementation
	env.OnActivity(ComposeGreeting, mock.Anything, "World").Return("Hello World!", nil)

	env.ExecuteWorkflow(GreetingWorkflow, "World")
	require.True(t, env.IsWorkflowCompleted())
	require.NoError(t, env.GetWorkflowError())

	var greeting string
	require.NoError(t, env.GetWorkflowResult(&greeting))
	require.Equal(t, "Hello World!", greeting)
}

// @@@SNIPEND
