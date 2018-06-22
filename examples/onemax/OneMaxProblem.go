package onemax

import (
	"time"

	"github.com/opticverge/goevolution/chromosome"
	"github.com/opticverge/goevolution/objective"
	"github.com/opticverge/goevolution/problem"
)

// Problem represents the OneMax problem is a problem type which evaluates the
// sum of the values in a Chromosome.
type Problem struct {
	problem.Problem
}

// ObjectiveFunction evaluates the chromosome and sets the fitness
func (p *Problem) ObjectiveFunction(chromo *chromosome.IChromosome) {
	chromos := (*chromo).(*Chromosome)
	fitness := 0
	for _, val := range chromos.Phenotype {
		fitness += val
	}
	chromos.SetFitness(float64(fitness))
}

// GenerateChromosome creates a new OneMax Chromosome
func (p *Problem) GenerateChromosome() chromosome.IChromosome {
	return NewChromosome(p.GetDimensions(), p.GetGenerator().Clone(time.Now().UnixNano()))
}

// NewProblem creates a new instance of the OneMax Problem
func NewProblem() problem.IProblem {
	p := &Problem{}
	p.SetName("One Max")
	p.SetObjective(objective.Maximisation)
	return p
}
