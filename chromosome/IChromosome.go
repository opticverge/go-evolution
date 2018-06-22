package chromosome

import (
	"github.com/opticverge/goevolution/generator"
)

// IChromosome is the interface which encourages generality and specificity
// allowing for our evolutionary optimiser to optimise any chromosome. If
// a problem is to be solved then embedding the IChromosome into the problem
// chromosome will encourage the implementation of the Generate, Mutate and
// Clone functions which are specific to the problem.
type IChromosome interface {

	// Functions to be implemented by the specific chromosome
	Generate()
	Mutate(float64)
	Clone(generator.IGenerator) IChromosome

	// Generic functions to be implemented by the base Chromosome
	// struct that will be embedded into all Chromosome variants
	SetFitness(float64)
	SetGenerator(generator.IGenerator)
	SetDimensions(int)

	GetFitness() float64
	GetDimensions() int
	GetGenerator() generator.IGenerator
	GetPhenotype() interface{}
}
