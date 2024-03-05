package simplewav

import (
	"encoding/binary"
	"io"
	"math"
	"os"
)

const (
	bitsPerSample = 16
	numChannels   = 1
)

// writeHeader writes a minimal WAV header to an io.Writer.
func writeHeader(w io.Writer, numSamples, sampleRate int) error {
	// Ensure all calculations use fixed-size types for binary writing.
	fileSize := uint32(44 + numSamples*numChannels*bitsPerSample/8)
	audioFormat := int16(1) // PCM = 1
	byteRate := uint32(sampleRate) * uint32(numChannels) * uint32(bitsPerSample) / 8
	blockAlign := uint16(numChannels * bitsPerSample / 8)
	subChunk2Size := uint32(numSamples) * uint32(numChannels) * uint32(bitsPerSample) / 8

	// Directly write the byte slices and fixed-size data to the writer
	parts := []struct {
		data interface{}
	}{
		{[]byte("RIFF")},
		{fileSize - 8},
		{[]byte("WAVE")},
		{[]byte("fmt ")},
		{uint32(16)}, // Subchunk1Size for PCM
		{audioFormat},
		{int16(numChannels)},
		{uint32(sampleRate)},
		{byteRate},
		{blockAlign},
		{int16(bitsPerSample)},
		{[]byte("data")},
		{subChunk2Size},
	}

	for _, part := range parts {
		if err := binary.Write(w, binary.LittleEndian, part.data); err != nil {
			return err
		}
	}

	return nil
}

// Write creates a WAV file from a slice of float64 audio samples
func Write(wave []float64, filename string, sampleRate int) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	if err := writeHeader(file, len(wave), sampleRate); err != nil {
		return err
	}
	buffer := make([]int16, len(wave))
	const maxInt16 = float64(math.MaxInt16)
	for i, sample := range wave {
		buffer[i] = int16(sample * maxInt16)
	}
	return binary.Write(file, binary.LittleEndian, buffer)
}
