package valueCommon

import (
	"fmt"
)

type ValueIterateCommon struct {
	//Qtable[state][action] = Q value
	qTable map[int]map[int]float64
	alpha  float64
	gammma float64
}

func (c *ValueIterateCommon) InitQTable(states []int, actions []int, initVal float64) {
	qTable := make(map[int]map[int]float64)
	for _, state := range states {
		qTable[state] = make(map[int]float64)
		for _, action := range actions {
			qTable[state][action] = initVal
		}
	}
	c.qTable = qTable
}

func (c *ValueIterateCommon) SetAlpha(alpha float64) {
	c.alpha = alpha
}

func (c *ValueIterateCommon) SetGammma(gammma float64) {
	c.gammma = gammma
}

func (c *ValueIterateCommon) Alpha() float64 {
	return c.alpha
}

func (c *ValueIterateCommon) Gamma() float64 {
	return c.gammma
}

func (c *ValueIterateCommon) QTable(state int, action int) float64 {
	return c.qTable[state][action]
}

func (c *ValueIterateCommon) SetQTable(state int, action int, value float64) {
	c.qTable[state][action] = value
}

func (c *ValueIterateCommon) PrintQTable() {
	for state, _ := range c.qTable {
		for action, qValue := range c.qTable[state] {
			fmt.Println(state, action, qValue)
		}
	}
}
