package test

import (
	"testing"
	"time"

	"github.com/opticverge/goevolution/examples/onemax"
	"github.com/opticverge/goevolution/generator"
	"github.com/opticverge/goevolution/solver"
)

func TestSolverEvolution(t *testing.T) {
	dimensions := 8
	populationSize := 10
	epochs := 2

	p := onemax.NewProblem()
	p.SetGenerator(generator.NewRandomGenerator(time.Now().UnixNano()))
	p.SetDimensions(dimensions)

	s := solver.NewSolver()
	s.SetEpochs(epochs)
	s.SetProblem(p)
	s.SetPopulationSize(populationSize)

	bestChromosome := s.Run()

	if bestChromosome == nil {
		t.Errorf("Expected output of solver.Run() to produce an IChromosome not nil")
	}

	if s.GetGeneration() != epochs {
		t.Errorf("Expected generation to be %v not %v", epochs, s.GetGeneration())
	}

	if len(s.GetPopulation()) != populationSize {
		t.Errorf("Expected population size to be %v not %v", populationSize, len(s.GetPopulation()))
	}
}
