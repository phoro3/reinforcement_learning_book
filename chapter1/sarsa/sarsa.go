package sarsa

import (
	"../env"
	"../valueCommon"
	"fmt"
)

type Sarsa struct {
	*valueCommon.ValueIterateCommon
}

func NewSarsa(states []int, actions []int, alpha float64, gamma float64) *Sarsa {
	c := valueCommon.ValueIterateCommon{}
	c.InitQTable(states, actions, 10)
	c.SetAlpha(alpha)
	c.SetGammma(gamma)
	s := Sarsa{&c}
	return &s
}

func (s *Sarsa) UpdateQTable(state int, action int, reward float64, nextState int, nextAction int) {
	value := (1-s.Alpha())*s.QTable(state, action) + s.Alpha()*(reward+s.Gamma()*s.QTable(nextState, nextAction))
	s.SetQTable(state, action, value)
}

func (s *Sarsa) runEpisode(e *env.Env) {
	var action int
	var nextState int
	var nextAction int
	var reward float64
	for state := e.InitState; ; {
		action = s.GetRandomAction(e.Actions)
		reward = e.GetReward(state, action)
		nextState = e.GetNextState(state, action)
		nextAction = s.GetRandomAction(e.Actions)
		s.UpdateQTable(state, action, reward, nextState, nextAction)
		state = nextState
		if state == e.EndState {
			break
		}
	}
}

func (s *Sarsa) Learn(e *env.Env, loopNum int) {
	for i := 1; i <= loopNum; i++ {
		if i%10 == 0 {
			fmt.Println("Episode: ", i)
		}
		s.runEpisode(e)
	}
}
