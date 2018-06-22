package problem

import (
	"github.com/opticverge/goevolution/generator"
	"github.com/opticverge/goevolution/objective"
)

// Problem encapsulates the properties and behaviours required for a Solver
// to produce a solution in accordance with the objective function. Any new
// problem that is to be created should embed the Problem struct. This way
// the new problem will be easily integrated into the evolutionary pipeline.
type Problem struct {
	name       string
	generator  generator.IGenerator
	dimensions int
	objective  objective.Objective
	IProblem
}

///////////////////////////////////////////////////////////////////////////////
// SETTERS ////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////

// SetDimensions sets the number of dimensions for a chromosome to generate.
func (p *Problem) SetDimensions(dimensions int) {
	p.dimensions = dimensions
}

// SetName sets the name of the problem.
func (p *Problem) SetName(name string) {
	p.name = name
}

// SetObjective sets the type of objective for the problem, for example, to
// maximise or minimise the objective function. Also influences how the
// population of a solver will be ordered.
func (p *Problem) SetObjective(objective objective.Objective) {
	p.objective = objective
}

// SetGenerator sets the generator to be used by the problem
func (p *Problem) SetGenerator(rng generator.IGenerator) {
	p.generator = rng
}

///////////////////////////////////////////////////////////////////////////////
// GETTERS ////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////

// GetObjective returns the objective type of the problem
func (p *Problem) GetObjective() objective.Objective {
	return p.objective
}

// GetGenerator returns the generator used for this problem
func (p *Problem) GetGenerator() generator.IGenerator {
	return p.generator
}

// GetName returns the name of the problem
func (p *Problem) GetName() string {
	return p.name
}

// GetDimensions returns the number of dimensions for this problem
func (p *Problem) GetDimensions() int {
	return p.dimensions
}
