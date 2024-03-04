package main

import (
	"encoding/binary"
	"os"
	"time"

	"github.com/xyproto/sawtooth"
)

const (
	sampleRate    = 44100
	bitsPerSample = 16
	numChannels   = 1
)

// writeWAVHeader writes a minimal WAV header to a file.
func writeWAVHeader(file *os.File, numSamples int) error {
	// WAV file header size: 44 bytes
	var fileSize int32 = 44 + int32(numSamples)*numChannels*bitsPerSample/8
	var audioFormat int16 = 1 // PCM = 1
	var byteRate int32 = int32(sampleRate) * int32(numChannels) * bitsPerSample / 8
	var blockAlign int16 = numChannels * bitsPerSample / 8
	var subChunk2Size int32 = int32(numSamples) * int32(numChannels) * bitsPerSample / 8

	// RIFF header
	_, err := file.Write([]byte("RIFF"))
	if err != nil {
		return err
	}
	binary.Write(file, binary.LittleEndian, fileSize-8) // File size minus RIFF header size
	_, err = file.Write([]byte("WAVE"))
	if err != nil {
		return err
	}

	// fmt subchunk
	_, err = file.Write([]byte("fmt "))
	if err != nil {
		return err
	}
	binary.Write(file, binary.LittleEndian, int32(16)) // Subchunk1Size (16 for PCM)
	binary.Write(file, binary.LittleEndian, audioFormat)
	binary.Write(file, binary.LittleEndian, int16(numChannels))
	binary.Write(file, binary.LittleEndian, int32(sampleRate))
	binary.Write(file, binary.LittleEndian, byteRate)
	binary.Write(file, binary.LittleEndian, blockAlign)
	binary.Write(file, binary.LittleEndian, int16(bitsPerSample))

	// data subchunk
	_, err = file.Write([]byte("data"))
	if err != nil {
		return err
	}
	binary.Write(file, binary.LittleEndian, subChunk2Size)

	return nil
}

func main() {
	// Generate the sawtooth wave
	duration := 1 * time.Second
	wave := sawtooth.GenerateSawtoothParticle(220, sampleRate, duration)

	// Create a new file
	file, err := os.Create("test.wav")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Write the WAV header
	numSamples := len(wave)
	if err := writeWAVHeader(file, numSamples); err != nil {
		panic(err)
	}

	// Write the audio data
	for _, sample := range wave {
		// Convert the sample from -1.0<=sample<=1.0 to 16-bit signed integer
		var intSample int16 = int16(sample * 32767)
		binary.Write(file, binary.LittleEndian, intSample)
	}
}
