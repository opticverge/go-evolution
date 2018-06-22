package generator

import (
	"math"
	"math/rand"

	"github.com/opticverge/goevolution/util"
)

// RandomGenerator is the base struct for all random number generators that
// depend on a seed and utilise Go's existing pattern for generating random
// numbers.
type RandomGenerator struct {
	generator rand.Rand
	seed      int64
	IGenerator
}

///////////////////////////////////////////////////////////////////////////////
// GETTERS ////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////

// GetSeed returns the seed of the random number generator
func (r *RandomGenerator) GetSeed() int64 {
	return r.seed
}

///////////////////////////////////////////////////////////////////////////////
// SETTERS ////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////

// SetSeed sets the seed for the random number generator
func (r *RandomGenerator) SetSeed(seed int64) {
	r.seed = seed
}

// SetSource sets the source of the generator
func (r *RandomGenerator) SetSource(rng rand.Rand) {
	r.generator = rng
}

///////////////////////////////////////////////////////////////////////////////
// IGENERATOR IMPLEMENTATIONS /////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////

// Float64 returns a randomly generated float64
func (r *RandomGenerator) Float64() float64 {
	return r.generator.Float64()
}

// Permutation generates a permutation of integers
func (r *RandomGenerator) Permutation(length int) []int {
	return r.generator.Perm(length)
}

// Intn generates a random integer between 0 and value where value cannot be
// less than zero
func (r *RandomGenerator) Intn(value int) int {
	return r.generator.Intn(value)
}

// IntRange generates a random int between the min and max values
func (r *RandomGenerator) IntRange(min int, max int) int {
	scale := int(math.Abs(float64(max - min)))
	index := r.generator.Intn(scale)
	return int(util.Translate(float64(index), 0.0, float64(scale), float64(min), float64(max)))
}

// FloatRange generates a random float64 between the min and max values
func (r *RandomGenerator) FloatRange(min float64, max float64) float64 {
	scale := math.Abs(max - min)
	index := r.generator.Float64() * scale
	return util.Translate(index, 0.0, scale, min, max)
}

// NormFloat64 generates a random float based on the normal distribution
func (r *RandomGenerator) NormFloat64() float64 {
	return r.generator.NormFloat64()
}

// Choice returns an array of selected
func (r *RandomGenerator) Choice(options []interface{}, size int) []interface{} {

	totalOptions := len(options)

	// if the number of selected is greater than the number of options then
	// return all options
	if size >= totalOptions {
		return options
	}

	// if the caller mistakenly provided a negative size then return an empty
	// array
	if size <= 0 {
		return make([]interface{}, 0)
	}

	// the map of selected choices ensuring uniqueness
	selectedMap := make(map[interface{}]interface{})
	selected := make([]interface{}, size)
	added := 0

	for added < size {
		value := r.generator.Intn(totalOptions)
		_, ok := selectedMap[value]
		if !ok {
			selectedMap[value] = true
			selected[added] = value
			added++
		}
	}

	return selected
}

// Clone generates a copy of the generator
func (r *RandomGenerator) Clone(seed int64) IGenerator {
	return NewRandomGenerator(seed)
}

///////////////////////////////////////////////////////////////////////////////
// CONSTRUCTOR ////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////

// NewRandomGenerator returns a new instance of a RandomGenerator
func NewRandomGenerator(seed int64) IGenerator {
	rng := &RandomGenerator{}
	rng.SetSeed(seed)
	rng.SetSource(*rand.New(rand.NewSource(seed)))
	return rng
}
