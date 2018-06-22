package solver

import (
	"math"
	"sort"
	"sync"
	"time"

	"github.com/opticverge/goevolution/chromosome"
	"github.com/opticverge/goevolution/objective"
	"github.com/opticverge/goevolution/problem"
)

// Solver represents the structure necessary to evolve a set of solutions
// against a problem. It implements most of the ISolver interface and acts
// as the base solver for all solvers.
type Solver struct {
	epochs         int
	generation     int
	population     []chromosome.IChromosome
	populationSize int
	problem        problem.IProblem
	ISolver
}

///////////////////////////////////////////////////////////////////////////////
// SETTERS ////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////

// SetProblem sets the problem for a solver to solve.
func (s *Solver) SetProblem(problem problem.IProblem) {
	s.problem = problem
}

// SetEpochs sets the max number of generations the solver should run for.
func (s *Solver) SetEpochs(epochs int) {
	s.epochs = epochs
}

// SetPopulationSize sets the number of chromosomes in the population
func (s *Solver) SetPopulationSize(populationSize int) {
	s.populationSize = populationSize
}

// SetPopulation sets an array of chromosomes as the population of the solver
func (s *Solver) SetPopulation(population []chromosome.IChromosome) {
	s.population = population
}

///////////////////////////////////////////////////////////////////////////////
// GETTERS ////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////

// GetGeneration returns the current generation of the solver.
func (s *Solver) GetGeneration() int {
	return s.generation
}

// GetPopulation returns the population of chromosomes
func (s *Solver) GetPopulation() []chromosome.IChromosome {
	return s.population
}

///////////////////////////////////////////////////////////////////////////////
// INTERFACE METHODS //////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////

// Initialise generates, scores and sorts a population of chromosomes
func (s *Solver) Initialise() {
	s.population = s.GenerateChromosomes(s.populationSize)
	s.EvaluateChromosomes(nil)
	s.SortChromosomes(nil)
}

// EvaluateChromosomes will evaluate the provided list of chromosomes
// or default to evaluating the population of the solver
func (s *Solver) EvaluateChromosomes(chromosomes *[]chromosome.IChromosome) {

	var chromosomesToEvaluate []chromosome.IChromosome

	if chromosomes != nil {
		chromosomesToEvaluate = *chromosomes
	} else {
		chromosomesToEvaluate = s.population
	}

	var wg sync.WaitGroup

	count := len(chromosomesToEvaluate)
	for i := 0; i < count; i++ {
		wg.Add(1)
		go func(toEvaluate *chromosome.IChromosome) {
			defer wg.Done()
			s.problem.ObjectiveFunction(toEvaluate)
		}(&chromosomesToEvaluate[i])
	}

	wg.Wait()
}

// GenerateChromosomes generates an array of chromosomes based on the value of
// count. It presumes that the problem has implemented the GenerateChromosome
// function since the problem is the specific component of the evolutionary
// process.
func (s *Solver) GenerateChromosomes(count int) []chromosome.IChromosome {
	chromosomes := make([]chromosome.IChromosome, count)
	var wg sync.WaitGroup
	for i := 0; i < count; i++ {
		wg.Add(1)
		go func(pos int, generatedChromosomes *[]chromosome.IChromosome) {
			defer wg.Done()
			generatedChromosome := s.problem.GenerateChromosome()
			generatedChromosome.Generate()
			(*generatedChromosomes)[pos] = generatedChromosome
		}(i, &chromosomes)
	}
	wg.Wait()
	return chromosomes
}

// Run is the gateway to initialising the evolutionary optimisation process.
func (s *Solver) Run() chromosome.IChromosome {

	// prepares the solver
	s.Setup()

	// initialises the population of chromosomes to be evolved
	s.Initialise()

	// evolves the chromosomes for the specified number of generations
	for s.epochs == -1 || s.generation < s.epochs {
		s.generation++
		s.Evolve()
	}

	s.TearDown()

	s.SortChromosomes(nil)

	return s.population[0]
}

// Evolve triggers the evolutionary process for mutation and replacement
func (s *Solver) Evolve() {
	s.Mutate()
	s.Replace()
}

// Mutate initiates the mutation process for all of the chromosomes in the
// population
func (s *Solver) Mutate() {
	var wg sync.WaitGroup
	for i := 0; i < s.populationSize; i++ {
		wg.Add(1)
		go func(sourceChromosome chromosome.IChromosome, pos int, population *[]chromosome.IChromosome) {
			defer wg.Done()
			(*population)[pos] = s.MutateChromosomes(sourceChromosome, pos)
		}(s.population[i], i, &s.population)
	}
	wg.Wait()
}

// MutateChromosomes generates mutations of the source chromosome.
func (s *Solver) MutateChromosomes(sourceChromosome chromosome.IChromosome, rank int) chromosome.IChromosome {

	// TODO: Allow for setting mutation strategy to determine clone count
	cloneCount := s.populationSize //int(math.Max(float64((s.populationSize / (rank + 1))), 1.0))

	// TODO: Calculate mutation probability independently
	mutationProbability := math.Exp(-2.4 * float64((s.populationSize-rank)/s.populationSize))

	// placeholder for mutated chromosomes
	clones := make([]chromosome.IChromosome, cloneCount)

	var wg sync.WaitGroup

	for i := 0; i < cloneCount; i++ {
		wg.Add(1)
		go func(source chromosome.IChromosome, pos int, sourceClones *[]chromosome.IChromosome) {
			defer wg.Done()

			// first we clone the chromsome and the selected generator from
			// the problem
			clonedGenerator := s.problem.GetGenerator().Clone(time.Now().UnixNano())
			clone := source.Clone(clonedGenerator)
			clone.Mutate(mutationProbability)
			(*sourceClones)[pos] = clone
		}(sourceChromosome, i, &clones)
	}

	wg.Wait()

	// when complete we evaluate the clones
	s.EvaluateChromosomes(&clones)

	// we then sort based on the objective function
	s.SortChromosomes(&clones)

	// we retrieve the best solution in the population of clones then we
	// replace the original chromosome if the best clone is better
	bestChromosome := clones[0]

	if s.problem.GetObjective() == objective.Maximisation {
		if bestChromosome.GetFitness() < sourceChromosome.GetFitness() {
			bestChromosome = sourceChromosome
		}
	} else {
		if bestChromosome.GetFitness() > sourceChromosome.GetFitness() {
			bestChromosome = sourceChromosome
		}
	}

	return bestChromosome
}

// SortChromosomes sorts a list of IChromosomes according to the objective of
// the problem.
func (s *Solver) SortChromosomes(chromosomes *[]chromosome.IChromosome) {

	var chromosomesToSort []chromosome.IChromosome
	if chromosomes != nil {
		chromosomesToSort = *chromosomes
	} else {
		chromosomesToSort = s.population
	}

	sort.Sort(chromosome.Chromosomes(chromosomesToSort))

	if s.problem.GetObjective() == objective.Maximisation {
		sort.Sort(sort.Reverse(chromosome.Chromosomes(chromosomesToSort)))
	}
}

// Replace uses an empiricist approach to remove the worst in the population
func (s *Solver) Replace() {

	// get the replacement count of the population
	replaceCount := int(0.1 * float64(s.populationSize))

	// remove the last 10%
	s.population = s.population[0 : s.populationSize-replaceCount]

	// generate the replacements
	newChromosomes := s.GenerateChromosomes(replaceCount)

	// sort the generated chromosomes
	s.SortChromosomes(&newChromosomes)

	// add them to the population
	s.population = append(s.population, newChromosomes...)

	// sort the population
	s.SortChromosomes(nil)
}

// Setup provides the solver with the opportunity to prepare the solver for the
// evolutionary process.Typically this will involve initialisation as well as
// opportunities for future integration tasks.
func (s *Solver) Setup() {
	s.generation = 1
}

// TearDown provides the opposite of what Setup provides.
func (s *Solver) TearDown() {}

// NewSolver generates a new Solver which implements the ISolver interface
func NewSolver() ISolver {
	return &Solver{}
}
