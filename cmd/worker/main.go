package main

import (
	"log"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"

	"github.com/desolaafujah/DependencyUpgrader/workflows"
)

func main() {
	// connect to temporal server
	c, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatalln("unable to create temporal client", err)
	}
	defer c.Close()

	// register worker on task queue
	w := worker.New(c, "UPGRADER_TASK_QUEUE", worker.Options{})

	// register workflow + activities
	w.RegisterWorkflow(workflows.DummyWorkflow)

	// start worker
	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalln("unable to start worker", err)
	}
}
