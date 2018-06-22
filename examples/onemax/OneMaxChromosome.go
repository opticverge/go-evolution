package onemax

import (
	"github.com/opticverge/goevolution/chromosome"
	"github.com/opticverge/goevolution/generator"
)

// Chromosome represents the structure for producing a one max chromosome.
type Chromosome struct {
	chromosome.Chromosome
	Phenotype []int
}

// Generate creates a new chromosome for the OneMax problem
func (c *Chromosome) Generate() {
	c.Phenotype = make([]int, c.GetDimensions())
	for i := 0; i < c.GetDimensions(); i++ {
		c.Phenotype[i] = c.GenerateGene()
	}
}

// Mutate applies a mutation to the chromosome
func (c *Chromosome) Mutate(mutationProbability float64) {
	for i := 0; i < c.GetDimensions(); i++ {
		if c.GetGenerator().Float64() < mutationProbability {
			c.Phenotype[i] = c.GenerateGene()
		}
	}
}

// Clone creates a new copy of the chromosome
func (c *Chromosome) Clone(rng generator.IGenerator) chromosome.IChromosome {
	clone := &Chromosome{}
	clone.SetGenerator(rng)
	clone.SetDimensions(c.GetDimensions())
	clone.Phenotype = make([]int, c.GetDimensions())
	copy(clone.Phenotype, c.Phenotype)
	return clone
}

// GetPhenotype returns the phenotype of the chromosome
func (c *Chromosome) GetPhenotype() interface{} {
	return c.Phenotype
}

// GenerateGene defines how an individual value of the chromosome should be
// generated
func (c *Chromosome) GenerateGene() int {
	// return c.GetGenerator().Intn(2)
	if c.GetGenerator().Float64() > 0.5 {
		return 1
	}
	return 0
}

// NewChromosome creates a new instance of the OneMax Chromosome
func NewChromosome(dimensions int, rng generator.IGenerator) chromosome.IChromosome {
	chr := &Chromosome{}
	chr.SetGenerator(rng)
	chr.SetDimensions(dimensions)
	return chr
}
