package chromosome

// Chromosomes represents a list of IChromosomes. This allows us to tap into
// the sort api in go.
type Chromosomes []IChromosome

func (c Chromosomes) Len() int {
	return len(c)
}

func (c Chromosomes) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}

func (c Chromosomes) Less(i, j int) bool {
	return c[i].GetFitness() < c[j].GetFitness()
}
