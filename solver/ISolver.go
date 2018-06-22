package solver

import (
	"github.com/opticverge/goevolution/chromosome"
	"github.com/opticverge/goevolution/problem"
)

// ISolver interface encapsulates the behaviours required for a solver to
// be integrated into the evolutionary pipeline
type ISolver interface {

	// Functions which must be implemented by each Solver
	Run() chromosome.IChromosome
	Evolve()
	Mutate()
	Replace()

	// Functions that will be implemented in the base solver struct. This
	// struct will normally be embedded within a new Solver struct. See
	// examples to learn more.
	Initialise()
	GenerateChromosome() chromosome.IChromosome
	GenerateChromosomes(int) []chromosome.IChromosome
	SortChromosomes(*[]chromosome.IChromosome)
	EvaluateChromosomes(*[]chromosome.IChromosome)

	// SETTERS
	SetProblem(problem.IProblem)
	SetEpochs(int)
	SetPopulationSize(int)
	SetPopulation([]chromosome.IChromosome)

	// GETTERS
	GetGeneration() int
	GetPopulation() []chromosome.IChromosome

	// LIFECYCLE MANAGEMENT
	Setup()
	TearDown()
}
