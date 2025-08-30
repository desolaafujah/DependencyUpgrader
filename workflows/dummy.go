package workflows

import (
	"time"

	"go.temporal.io/sdk/workflow"
)

func DummyWorkflow(ctx workflow.Context, name string) (string, error) {
	logger := workflow.GetLogger(ctx)
	logger.Info("DummyWorklfow started", "name", name)

	// simulate some work
	workflow.Sleep(ctx, time.Second*2)

	result := "hello, " + name + "! temporal is working"
	logger.Info("DummyWorkflow finished", "result", result)

	return result, nil
}
