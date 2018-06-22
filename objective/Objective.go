package objective

// Objective is a type which defines whether an optimiser is solving a
// maximisation or minimisation
type Objective string

const (
	// Maximisation is an objective that is used by an optimiser when trying to
	// maximise the result of evaluating a chromosome. It is typically used by
	// the optimiser to sort the chromosomes if there is a population of them
	Maximisation Objective = "Maximisation"

	// Minimisation defines a type of objective that aims to minimise the
	// evaluation function.
	Minimisation Objective = "Minimisation"
)
