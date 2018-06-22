package generator

import (
	"math/rand"

	"github.com/tidwall/weyl"
)

// WeylGenerator is another type of random number generator. Please see the
// repo for more information. Since the RandomGenerator and Weyl hook into the
// same API for Rand we can embed it.
type WeylGenerator struct {
	RandomGenerator
}

// Clone creates a new instance of the WeylGenerator provided a seed
func (r *WeylGenerator) Clone(seed int64) IGenerator {
	return NewWeylGenerator(seed)
}

// NewWeylGenerator creates a new instance of the WeylGenerator
func NewWeylGenerator(seed int64) IGenerator {
	generator := &WeylGenerator{}
	generator.seed = seed
	generator.generator = *rand.New(weyl.NewSource(generator.seed))
	return generator
}
