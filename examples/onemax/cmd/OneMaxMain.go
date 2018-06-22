package main

import (
	"fmt"
	"runtime"
	"time"

	"github.com/opticverge/goevolution/examples/onemax"
	"github.com/opticverge/goevolution/generator"
	"github.com/opticverge/goevolution/solver"
)

func main() {

	// set the max processors for your cpu
	runtime.GOMAXPROCS(runtime.NumCPU())

	// set some of the properties
	dimensions := 128
	populationSize := 100
	epochs := 100

	// Generate the one max problem
	p := onemax.NewProblem()
	p.SetGenerator(generator.NewRandomGenerator(time.Now().UnixNano()))
	p.SetDimensions(dimensions)

	// create the generic solver for the one max problem
	s := solver.NewSolver()
	s.SetEpochs(epochs)
	s.SetProblem(p)
	s.SetPopulationSize(populationSize)

	// initiate the evolutionary process
	bestChromosome := s.Run()

	fmt.Println(bestChromosome.GetFitness(), bestChromosome.GetPhenotype())
}
