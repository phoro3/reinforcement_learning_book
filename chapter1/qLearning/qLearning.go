package qLearning

import (
	"../env"
	"../valueCommon"
	"fmt"
)

type qLearning struct {
	*valueCommon.ValueIterateCommon
}

func NewQLearning(states []int, actions []int, alpha float64, gamma float64) *qLearning {
	c := valueCommon.ValueIterateCommon{}
	c.InitQTable(states, actions, 10)
	c.SetAlpha(alpha)
	c.SetGammma(gamma)
	q := qLearning{&c}
	return &q
}

func (q *qLearning) getMaxQValFromState(state int) float64 {
	maxVal := 0.0
	for _, val := range q.QTableMap(state) {
		if val > maxVal {
			maxVal = val
		}
	}
	return maxVal
}

func (q *qLearning) UpdateQTable(state int, action int, reward float64, nextState int) {
	value := (1-q.Alpha())*q.QTable(state, action) + q.Alpha()*(reward+q.Gamma()*q.getMaxQValFromState(nextState))
	q.SetQTable(state, action, value)
}

func (q *qLearning) runEpisode(e *env.Env) {
	var action int
	var nextState int
	var reward float64
	for state := e.InitState; ; {
		action = q.GetRandomAction(e.Actions)
		reward = e.GetReward(state, action)
		nextState = e.GetNextState(state, action)
		q.UpdateQTable(state, action, reward, nextState)
		state = nextState
		if state == e.EndState {
			break
		}
	}
}

func (q *qLearning) Learn(e *env.Env, loopNum int) {
	for i := 1; i <= loopNum; i++ {
		if i%10 == 0 {
			fmt.Println("Episode: ", i)
		}
		q.runEpisode(e)
	}
}
