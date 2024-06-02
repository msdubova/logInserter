package main

import (
	"goroutine/batch"
	"goroutine/inserter"
)

func main() {

	var ins = batch.NewInserter(batch.Inserter{})

	inserter.GenerateAndInsertLogs(ins)
}
