package sarsa

import (
	"../env"
	"fmt"
)

func InitSarsa(states []int, actions []int, alpha float64, gamma float64) *Sarsa {
	s := Sarsa{}
	s.InitQTable(states, actions, 10)
	s.SetAlpha(alpha)
	s.SetGammma(gamma)
	return &s
}

func runEpisode(e *env.Env, s *Sarsa) {
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

func Learn(e *env.Env, s *Sarsa, loopNum int) {
	for i := 1; i <= loopNum; i++ {
		if i%10 == 0 {
			fmt.Println("Episode: ", i)
		}
		runEpisode(e, s)
	}
}
