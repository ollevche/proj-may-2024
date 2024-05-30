package log

import "math/rand"

type Log struct {
	ID    string
	Value string
}

func NewRandom(id string) Log {
	const allowedLetters = "abcde123"

	var v string

	for i := 0; i < 10; i++ {
		v += string(allowedLetters[rand.Intn(len(allowedLetters))])
	}

	return Log{
		ID:    id,
		Value: v,
	}
}
