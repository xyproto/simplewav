package sawtooth

import (
	"math"
	"time"
)

// GenerateSawtoothFormula generates a sawtooth wave using a mathematical formula.
// frequency specifies the frequency of the sawtooth wave in Hz.
// sampleRate specifies the number of samples per second.
// duration specifies the duration of the wave to be generated.
func GenerateSawtoothFormula(frequency float64, sampleRate int, duration time.Duration) []float64 {
	T := 1 / frequency // Calculate the period of the wave.
	durationSeconds := duration.Seconds()
	samples := int(durationSeconds * float64(sampleRate)) // Calculate the total number of samples.
	waveform := make([]float64, samples)                  // Initialize the slice to store the waveform.
	for i := range waveform {
		t := float64(i) / float64(sampleRate) // Calculate the time for the current sample.
		// Calculate the waveform value using the sawtooth formula.
		waveform[i] = 2 * (t/T - math.Floor(0.5+t/T))
	}
	return waveform
}

// GenerateSawtoothParticle generates a sawtooth wave by simulating a particle moving in a space from Y = -1 to Y = 1.
func GenerateSawtoothParticle(frequency float64, sampleRate int, duration time.Duration) []float64 {
	durationSeconds := duration.Seconds()
	samples := int(durationSeconds * float64(sampleRate))
	waveform := make([]float64, samples)
	velocityY := 2 * frequency // Velocity to cover from -1 to 1 in one period
	stepY := (velocityY / float64(sampleRate))
	// Adjust initial positionY by subtracting one increment step
	positionY := 0 - stepY
	for i := range waveform {
		positionY += stepY
		// Boundary check: if Y reaches or exceeds 1, reset it to -1 + stepY
		if positionY >= 1 {
			positionY = -1 + stepY
		}
		waveform[i] = positionY
	}
	return waveform
}
