package log

import "math/rand"

type Log struct {
	ID    int
	Value string
}

func NewLog() Log {
	const letters = "olgadubova"

	var s string
	for i := 0; i < 10; i++ {
		s += string(letters[rand.Intn(len(letters))])
	}

	return Log{
		ID:    rand.Intn(100),
		Value: s,
	}
}
