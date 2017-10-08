package sarsa

import (
	"fmt"
)

type Sarsa struct {
	//Qtable[state][action] = Q value
	qTable map[int]map[int]float64
	alpha  float64
	gammma float64
}

func (s *Sarsa) InitQTable(states []int, actions []int, initVal float64) {
	qTable := make(map[int]map[int]float64)
	for _, state := range states {
		qTable[state] = make(map[int]float64)
		for _, action := range actions {
			qTable[state][action] = initVal
		}
	}
	s.qTable = qTable
}

func (s *Sarsa) UpdateQTable(state int, action int, reward float64, nextState int, nextAction int) {
	s.qTable[state][action] =
		(1-s.alpha)*s.qTable[state][action] + s.alpha*(reward+s.gammma*s.qTable[nextState][nextAction])
}

func (s *Sarsa) SetAlpha(alpha float64) {
	s.alpha = alpha
}

func (s *Sarsa) SetGammma(gammma float64) {
	s.gammma = gammma
}

func (s *Sarsa) PrintQTable() {
	for state, _ := range s.qTable {
		for action, qValue := range s.qTable[state] {
			fmt.Println(state, action, qValue)
		}
	}
}
