package main

import (
	"flag"
	"fmt"
	"hermes-service/engine"

	"hermes-service/config"
	"log"
)

func main() {
	env := flag.String("env", "dev", "Execution environment")
	flag.Parse()
	log.Printf("Starting application  - %s", *env)

	c := config.BuildConfig(*env)

	e, err := engine.New(c)

	if err != nil {
		panic(fmt.Sprintf("can't start engine: %v", err))
	}

	if c.Sqs.Listener.Enabled {
		worker, err := engine.NewQueueWorker(c.AWS, &c.Sqs.Listener)
		if err != nil {
			panic(fmt.Sprintf("Couldn't initialize Sqs worker: %s", err.Error()))
		}
		err = worker.Start(e)
		if err != nil {
			panic(fmt.Sprintf("Error starting Sqs worker: %s", err.Error()))
		}
	}
}
