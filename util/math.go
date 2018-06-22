package util

// Translate translates values from one range to another.
func Translate(value float64, sourceLower float64, sourceUpper float64, targetLower float64, targetUpper float64) float64 {
	return targetLower + (value-sourceLower)*((targetUpper-targetLower)/(sourceUpper-sourceLower))
}
