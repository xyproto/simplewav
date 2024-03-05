package main

import (
	"time"

	"github.com/xyproto/sawtooth"
	"github.com/xyproto/simplewav"
)

func main() {
	const (
		duration           = 1 * time.Second
		frequency  float64 = 220
		sampleRate         = 44100
	)

	// Generate a sawtooth signal
	wave := sawtooth.GenerateSawtoothParticle(frequency, sampleRate, duration)

	// Write the wave to test.wav
	if err := simplewav.Write(wave, "test.wav", sampleRate); err != nil {
		panic(err)
	}
}
