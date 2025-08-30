package main

import (
	"context"
	"log"
	"time"

	"go.temporal.io/sdk/client"

	"github.com/desolaafujah/DependencyUpgrader/workflows"
)

func main() {
	// connect to temporal
	c, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatalln("unable to create temporal client", err)
	}
	defer c.Close()

	// start workflow execution
	workflowOptions := client.StartWorkflowOptions{
		ID:        "dummy_workflow_" + time.Now().Format("150405"),
		TaskQueue: "UPGRADER_TASK_QUEUE",
	}

	we, err := c.ExecuteWorkflow(context.Background(), workflowOptions, workflows.DummyWorkflow, "Desola")
	if err != nil {
		log.Fatalln("unable to execute workflow", err)
	}

	log.Println("started workflow", "WorkflowID", we.GetID(), "RunID", we.GetRunID())

	// get result
	var result string
	err = we.Get(context.Background(), &result)
	if err != nil {
		log.Fatalln("unable to get workflow result", err)
	}

	log.Println("workflow result: ", result)
}
