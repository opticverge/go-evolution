package chromosome

import (
	"github.com/opticverge/goevolution/generator"
)

// Chromosome is the base struct for all problem chromosomes and implements
// the most common behaviours for all chromosomes based on the
// IChromosome interface.
type Chromosome struct {
	fitness    float64
	dimensions int
	generator  generator.IGenerator
	IChromosome
}

///////////////////////////////////////////////////////////////////////////////
// SETTERS ////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////

// SetFitness sets the fitness of the chromosome
func (c *Chromosome) SetFitness(fitness float64) {
	c.fitness = fitness
}

// SetGenerator sets the generator of the chromosome
func (c *Chromosome) SetGenerator(generator generator.IGenerator) {
	c.generator = generator
}

// SetDimensions sets the dimensions of the chromosome
func (c *Chromosome) SetDimensions(dimensions int) {
	c.dimensions = dimensions
}

///////////////////////////////////////////////////////////////////////////////
// GETTERS ////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////

// GetFitness returns the fitness of the chromosome.
func (c *Chromosome) GetFitness() float64 {
	return c.fitness
}

// GetDimensions returns the number of dimensions of the chromosome.
func (c *Chromosome) GetDimensions() int {
	return c.dimensions
}

// GetGenerator returns the generator for this chromosome
func (c *Chromosome) GetGenerator() generator.IGenerator {
	return c.generator
}
