package problem

import (
	"github.com/opticverge/goevolution/chromosome"
	"github.com/opticverge/goevolution/generator"
	"github.com/opticverge/goevolution/objective"
)

// IProblem represents the interface by which all problems implement. Anything
// that implements the IProblem interface will be able to included within the
// evolutionary pipeline.
type IProblem interface {

	// Functions to be implemented by the specific problem
	ObjectiveFunction(*chromosome.IChromosome)
	GenerateChromosome() chromosome.IChromosome

	// Generic functions to be implemented by the base Problem
	// struct that will be embedded into all Problem variants

	// Setters
	SetDimensions(int)
	SetName(string)
	SetGenerator(generator.IGenerator)

	// Getters
	GetDimensions() int
	GetName() string
	GetObjective() objective.Objective
	GetGenerator() generator.IGenerator
}
