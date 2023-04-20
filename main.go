package main

import (
	"golearning/examples"
	"golearning/experiments"
	"time"
)

// go run main.go
func main() {
	if false {
		examples.ExampleBase62UUID()
	}
	if false {
		examples.ExampleCas()
	}
	if false {
		examples.ExampleKvs()
	}
	if false {
		examples.ExampleZipAndUnzip()
	}
	if false {
		go examples.RunGrpcServer()
		time.Sleep(2 * time.Second)
		examples.RunGrpcClient()
		time.Sleep(time.Second)
	}

	if false {
		experiments.ExperimentFilePath()
	}
	if false {
		experiments.ExperimentStringsBuilder()
	}
	if false {
		experiments.ExperimentStringElementType()
	}
}
