package benchmark

import (
	_ "fmt"
	"runtime"
	"sync"
	"testing"
	"time"

	"github.com/opticverge/goevolution/examples/onemax"
	"github.com/opticverge/goevolution/generator"
	"github.com/opticverge/goevolution/problem"
	"github.com/opticverge/goevolution/solver"
)

func BenchmarkEvolutionOneMax(b *testing.B) {
	// This benchmark basically asks how many times we can generate 1,000,000
	// chromosomes. With the values below we are looking at approximately
	// 1,300,000 chromosomes per second.

	runtime.GOMAXPROCS(8)
	dimensions := 64
	populationSize := 100
	epochs := 100

	p := onemax.NewProblem()
	p.SetDimensions(dimensions)
	p.SetGenerator(generator.NewWeylGenerator(time.Now().UnixNano()))

	s := solver.NewSolver()
	s.SetEpochs(epochs)
	s.SetProblem(p)
	s.SetPopulationSize(populationSize)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = s.Run()
	}
}

func BenchmarkEvolutionOneMaxMultiSolver(b *testing.B) {
	// This benchmark basically asks how many times we can generate 1,000,000
	// chromosomes. With the values below we are looking at approximately
	// 1,300,000 chromosomes per second.

	runtime.GOMAXPROCS(8)
	dimensions := 64
	populationSize := 100
	epochs := 100
	multiSolverCount := 2
	var wg sync.WaitGroup
	p := onemax.NewProblem()
	p.SetDimensions(dimensions)
	rng := generator.NewWeylGenerator(time.Now().UnixNano())
	// rng := generator.NewRandomGenerator(time.Now().UnixNano())
	b.ResetTimer()
	for i := 0; i < b.N; i++ {

		for j := 0; j < multiSolverCount; j++ {
			wg.Add(1)
			go func(prob problem.IProblem) {

				defer wg.Done()

				prob.SetGenerator(rng.Clone(time.Now().UnixNano()))

				s := solver.NewSolver()
				s.SetEpochs(epochs)
				s.SetProblem(prob)
				s.SetPopulationSize(populationSize)

				_ = s.Run()
			}(p)
		}
		wg.Wait()
	}
}
