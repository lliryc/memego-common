package common

import (
	"math/rand"
	"sort"
)

type Generation []Instance

func (gen Generation) Next(r int, policy Policy) Generation {	
	n := len(gen)
	sort.Sort(gen)
	nextGen := make(Generation, n)
	var cursor int = 0 
	// reproduction
	reprN := policy.GetReproductionN()
	var reproducibles Generation = gen[:reprN]
	copy(nextGen[cursor:], reproducibles)	
	cursor = reprN
	// crossover
	var crossoverN int = policy.GetCrossoverN()
	parents1 := gen[:crossoverN]
	parents2 := gen[crossoverN: 2 * crossoverN]	
	children := make(Generation, crossoverN)	
	for i := 0; i < crossoverN; i++ {
		parent1 := parents1[i]
		parent2 := parents2[i]
		children[i] = parent1.CrossOver(parent2)
	}
	copy(nextGen[cursor:], children)
	cursor += crossoverN
	// mutations
	var mutationN int = policy.GetMutationN()
	eliteSize := int(0.75 * float32(n))
	mutants := make(Generation, mutationN)	
	for i := 0; i < mutationN; i++ {
		picked := rand.Intn(eliteSize)
		mutants[i] = gen[picked].Mutation()		
	}
	copy(nextGen[cursor:], mutants)
	cursor += mutationN
	for i := 0; i < n; i++ {
		nextGen[i].ComputeFitness()
	}
	sort.Sort(nextGen)
	return nextGen
}

func (gen Generation) Len() int { 
	return len(gen) 
}

func (pq Generation) Less(i, j int) bool {
	return pq[i].Less(pq[j])
}

func (pq Generation) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq Generation) BestFit() float32 {
	return pq[0].Fitness()
}

func (pq Generation) BestInstance() Instance {
	return pq[0]
}
