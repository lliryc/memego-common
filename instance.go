package common

type Instance interface{
	CrossOver (Instance) Instance
	Mutation () Instance
	Reproduce () Instance
	Improve () Instance
	ComputeFitness() float32
	Fitness() float32
	Less(Instance) bool
}



