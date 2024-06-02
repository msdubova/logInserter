package inserter

import (
	"fmt"
	"goroutine/log"
	"math/rand"
)

type Inserter interface {
	Insert(logs []log.Log)
}

func GenerateAndInsertLogs(inserter Inserter) {

	for i := 0; i < 10; i++ {
		logs := make([]log.Log, 0, rand.Intn(10)+1)
		for j := 0; j < cap(logs); j++ {
			l := log.NewLog(fmt.Sprintf("%d-%d", i, j))
			logs = append(logs, l)
		}

		fmt.Println("Сгенерированные логи: ", logs)
		inserter.Insert(logs)
	}

}
