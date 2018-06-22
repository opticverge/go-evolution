package generator

import "math/rand"

// IGenerator represents the interface for generators of random numbers. This
// is intended to enable different types of random number generators for the
// evolutionary process.
type IGenerator interface {
	// Functions related to generation
	Float64() float64
	Intn(int) int
	NormFloat64() float64
	IntRange(int, int) int
	FloatRange(float64, float64) float64
	Permutation(int) []int
	Choice([]interface{}, int) []interface{}

	// Functions related to the attributes of a generator

	// Getters
	GetSeed() int64
	GetGenerator() IGenerator

	// Setters
	SetSource(rng rand.Rand)
	SetSeed(int64)

	Clone(int64) IGenerator
}
