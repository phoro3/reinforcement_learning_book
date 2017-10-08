package env

type Env struct {
	States  []int
	Actions []int
	//stateMatrix[state][action] = nextState
	stateMatrix map[int]map[int]int
	//rewardMatrix[state][action] = reward
	rewardMatrix map[int]map[int]float64
	currentState int
	nextState    int
	InitState    int
	EndState     int
}

func (e *Env) SetStateMatrix(matrix map[int]map[int]int) {
	e.stateMatrix = matrix
}

func (e *Env) SetRewardMatrix(matrix map[int]map[int]float64) {
	e.rewardMatrix = matrix
}

func (e *Env) GetNextState(state int, action int) int {
	nextState := e.stateMatrix[state][action]
	e.nextState = nextState
	return nextState
}

func (e *Env) GetReward(state int, action int) float64 {
	return e.rewardMatrix[state][action]
}

func (e *Env) Next() {
	e.currentState = e.nextState
}

func (e *Env) Init() {
	e.currentState = e.InitState
}
