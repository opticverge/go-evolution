package test

import (
	"testing"
	"time"

	"github.com/opticverge/goevolution/generator"
)

func TestRandomGeneratorChoiceLength(t *testing.T) {

	// GIVEN
	rng := generator.NewRandomGenerator(time.Now().UnixNano())
	choices := []interface{}{1, 2, 3, 4, 5, 6, 7}

	// WHEN
	selected := rng.Choice(choices, 2)

	// THEN
	if len(selected) != 2 {
		t.Errorf("Expected selected length to be %v, Actual %v", 2, len(selected))
	}
}

func TestRandomGeneratorChoiceSizeGreaterThanChoices(t *testing.T) {

	// GIVEN
	rng := generator.NewRandomGenerator(time.Now().UnixNano())
	choices := []interface{}{1, 2, 3, 4, 5, 6, 7}

	// WHEN
	selected := rng.Choice(choices, 8)

	// THEN
	if len(selected) != len(choices) {
		t.Errorf("Expected selected length to be %v, Actual %v", len(choices), len(selected))
	}
}

func TestRandomGeneratorChoiceNegativeSize(t *testing.T) {

	// GIVEN
	rng := generator.NewRandomGenerator(time.Now().UnixNano())
	choices := []interface{}{1, 2, 3, 4, 5, 6, 7}

	// WHEN
	selected := rng.Choice(choices, -1)

	// THEN
	if len(selected) != 0 {
		t.Errorf("Expected selected length to be %v, Actual %v", 0, len(selected))
	}
}

func TestRandomGeneratorIntRange(t *testing.T) {

	// GIVEN
	rng := generator.NewRandomGenerator(time.Now().UnixNano())
	min := 1
	max := 10

	// WHEN
	value := rng.IntRange(min, max)

	// THEN
	if value < min || value > max {
		t.Errorf("Expected value to be within range %v and %v, Actual %v ", min, max, value)
	}
}

func TestRandomGeneratorIntRangeNegative(t *testing.T) {

	// GIVEN
	rng := generator.NewRandomGenerator(time.Now().UnixNano())
	min := -100
	max := -50

	// WHEN
	value := rng.IntRange(min, max)

	// THEN
	if value < min || value > max {
		t.Errorf("Expected value to be within range %v and %v, Actual %v ", min, max, value)
	}
}

func TestRandomGeneratorFloatRangeNegative(t *testing.T) {

	//GIVEN
	rng := generator.NewRandomGenerator(time.Now().UnixNano())
	min := -100.0
	max := -50.0

	// WHEN
	actual := rng.FloatRange(min, max)

	// THEN
	if actual < min || actual > max {
		t.Errorf("Expected value to be within range %v and %v, Actual %v ", min, max, actual)
	}
}

func TestRandomGeneratorFloatRangePositive(t *testing.T) {

	//GIVEN
	rng := generator.NewRandomGenerator(time.Now().UnixNano())
	min := 10.0
	max := 50.0

	// WHEN
	actual := rng.FloatRange(min, max)

	// THEN
	if actual < min || actual > max {
		t.Errorf("Expected value to be within range %v and %v, Actual %v ", min, max, actual)
	}
}

func TestRandomGeneratorFloatRangeNegativePositive(t *testing.T) {

	//GIVEN
	rng := generator.NewRandomGenerator(time.Now().UnixNano())
	min := -100.0
	max := 100.0

	// WHEN
	actual := rng.FloatRange(min, max)

	// THEN
	if actual < min || actual > max {
		t.Errorf("Expected value to be within range %v and %v, Actual %v ", min, max, actual)
	}
}
