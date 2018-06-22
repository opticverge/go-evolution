package benchmark

import (
	"testing"
	"time"

	"github.com/opticverge/goevolution/generator"
)

func BenchmarkRandomGeneratorFloat64(b *testing.B) {
	rng := generator.NewRandomGenerator(time.Now().UnixNano())
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = rng.Float64()
	}
}

func BenchmarkRandomGeneratorIntSmallN(b *testing.B) {
	rng := generator.NewRandomGenerator(time.Now().UnixNano())
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = rng.Intn(2)
	}
}

func BenchmarkRandomGeneratorIntBigN(b *testing.B) {
	rng := generator.NewRandomGenerator(time.Now().UnixNano())
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = rng.Intn(1024)
	}
}

func BenchmarkRandomGeneratorFloatRange(b *testing.B) {
	rng := generator.NewRandomGenerator(time.Now().UnixNano())
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = rng.FloatRange(0.0, 1.0)
	}
}

func BenchmarkRandomGeneratorNormFloat64(b *testing.B) {
	rng := generator.NewRandomGenerator(time.Now().UnixNano())
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = rng.NormFloat64()
	}
}

func BenchmarkWeylGeneratorFloat64(b *testing.B) {
	rng := generator.NewWeylGenerator(time.Now().UnixNano())
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = rng.Float64()
	}
}

func BenchmarkWeylGeneratorIntSmallN(b *testing.B) {
	rng := generator.NewWeylGenerator(time.Now().UnixNano())
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = rng.Intn(2)
	}
}

func BenchmarkWeylGeneratorFloat64OneMax(b *testing.B) {
	rng := generator.NewWeylGenerator(time.Now().UnixNano())
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if rng.Float64() > 0.5 {
			_ = 1
		} else {
			_ = 0
		}
	}
}

func BenchmarkWeylGeneratorIntBigN(b *testing.B) {
	rng := generator.NewWeylGenerator(time.Now().UnixNano())
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = rng.Intn(1024)
	}
}

func BenchmarkWeylGeneratorFloatRange(b *testing.B) {
	rng := generator.NewWeylGenerator(time.Now().UnixNano())
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = rng.FloatRange(0.0, 1.0)
	}
}

func BenchmarkWeylGeneratorFloatRangeLargeMax(b *testing.B) {
	rng := generator.NewWeylGenerator(time.Now().UnixNano())
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = rng.FloatRange(0.0, 10000.0)
	}
}

func BenchmarkWeylGeneratorNormFloat64(b *testing.B) {
	rng := generator.NewWeylGenerator(time.Now().UnixNano())
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = rng.NormFloat64()
	}
}

func BenchmarkWeylGeneratorChoice(b *testing.B) {
	rng := generator.NewWeylGenerator(time.Now().UnixNano())
	size := 1000
	choices := make([]interface{}, size)
	for i := 0; i < size; i++ {
		choices[i] = i
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = rng.Choice(choices, 2)
	}
}
