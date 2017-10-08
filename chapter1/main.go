package main

import (
	"./env"
	"./sarsa"
)

func initStateMatrix() map[int]map[int]int {
	matrix := make(map[int]map[int]int)
	matrix[1] = make(map[int]int)
	matrix[2] = make(map[int]int)
	matrix[3] = make(map[int]int)
	matrix[1][1] = 3
	matrix[1][2] = 2
	matrix[2][1] = 1
	matrix[2][2] = 4
	matrix[3][1] = 4
	matrix[3][2] = 1
	return matrix
}

func initRewardMatrix() map[int]map[int]float64 {
	matrix := make(map[int]map[int]float64)
	matrix[1] = make(map[int]float64)
	matrix[2] = make(map[int]float64)
	matrix[3] = make(map[int]float64)
	matrix[1][1] = 0
	matrix[1][2] = 1
	matrix[2][1] = -1
	matrix[2][2] = 1
	matrix[3][1] = 5
	matrix[3][2] = -100
	return matrix
}

func initEnv() *env.Env {
	states := []int{1, 2, 3, 4}
	actions := []int{1, 2}
	e := env.Env{
		States:    states,
		Actions:   actions,
		InitState: 1,
		EndState:  4}
	stateMatrix := initStateMatrix()
	e.SetStateMatrix(stateMatrix)
	rewardMatrix := initRewardMatrix()
	e.SetRewardMatrix(rewardMatrix)
	e.Init()

	return &e
}

func main() {
	loopNum := 3000
	alpha := 0.05
	gammma := 0.95
	e := initEnv()
	target := sarsa.NewSarsa(e.States, e.Actions, alpha, gammma)
	target.Learn(e, loopNum)
	target.PrintQTable()
}
