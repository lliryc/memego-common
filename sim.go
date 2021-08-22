package common

import (
	"fmt"
	"time"
)

type Simulation struct {
	Generation Generation
}

func (sim Simulation) Run(maxSameIterations int, policy Policy) Instance{		
	terminateCnt := 0 
	generation := sim.Generation
	var oldEval float32 = generation.BestFit()
	currentTime := time.Now()     
	fmt.Printf("Simulation started %s", currentTime.String())
	for i := 0; ; i++ {
		fmt.Printf("Step %d", i)
		newGeneration := generation.Next(1, policy)
		eval := newGeneration.BestFit()
		fmt.Printf("Best score = %f", eval)
		if oldEval >= eval{
			fmt.Printf("No improvements since the previous step. Go backward")
			terminateCnt++
			if terminateCnt >= maxSameIterations{
				break
			}
		} else {
			oldEval = eval
			generation = newGeneration
			terminateCnt = 0
		}		
	}
	currentTime = time.Now()
	fmt.Printf("Simulation ended %s", currentTime.String())
	return generation.BestInstance()
}



