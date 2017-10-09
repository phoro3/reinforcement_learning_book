package qLearning

import (
	"math/rand"
)

func (q *qLearning) GetRandomAction(actions []int) int {
	return actions[rand.Intn(len(actions))]
}
