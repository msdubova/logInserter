package batch

import (
	"fmt"
	"goroutine/inserter"
	"goroutine/log"
)

type Batch struct {
	logs []log.Log
}

type Inserter struct {
	source inserter.Inserter
	ch     chan<- Batch
}

func NewInserter(inserter inserter.Inserter) *Inserter {
	ch := make(chan Batch, 1)

	in := &Inserter{
		source: inserter,
		ch:     ch,
	}

	go runInsertion(in.source, ch)

	return in
}

type MyInserter struct{}

func (MyInserter) Insert(logs []log.Log) {
	fmt.Println("Вставка логов - ", logs)
}
func runInsertion(inserter inserter.Inserter, ch <-chan Batch) {
	fmt.Println("start run insertion - Розпочато")

	var logsToInsert []log.Log

	for {
		fmt.Println("Iterating in runInsertion Ітерація в процесі")

		select {
		case b := <-ch:
			fmt.Println("Got new batch of logs")
			logsToInsert = append(logsToInsert, b.logs...)
			if len(logsToInsert) >= 15 {
				fmt.Println("Inserting logs due to batch size")
				inserter.Insert(logsToInsert)
				logsToInsert = nil
			}
		}
	}
}

func (in Inserter) Insert(logs []log.Log) {
	fmt.Println("відправка батчу логів у канал")
	in.ch <- Batch{logs: logs}
	fmt.Println("відправлено батчу логів у канал")

}
