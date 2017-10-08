package sarsa

import (
	"math/rand"
)

func (s *Sarsa) GetRandomAction(actions []int) int {
	return actions[rand.Intn(len(actions))]
}
