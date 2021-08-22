package common
type Policy interface {
	GetPopulationSize() int
	GetSolutionN() int
	GetCrossoverN() int
	GetReproductionN() int
	GetMutationN() int
	SetGeneration(int)
}
