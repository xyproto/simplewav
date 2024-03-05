package simplewav

import (
	"encoding/binary"
	"os"
)

const (
	bitsPerSample = 16
	numChannels   = 1
)

// writeHeader writes a minimal WAV header to a file
func writeHeader(file *os.File, numSamples, sampleRate int) error {
	// WAV file header size: 44 bytes
	var (
		fileSize      int32 = 44 + int32(numSamples)*numChannels*bitsPerSample/8
		audioFormat   int16 = 1 // PCM = 1
		byteRate      int32 = int32(sampleRate) * int32(numChannels) * bitsPerSample / 8
		blockAlign    int16 = numChannels * bitsPerSample / 8
		subChunk2Size int32 = int32(numSamples) * int32(numChannels) * bitsPerSample / 8
	)

	// RIFF header
	if _, err := file.Write([]byte("RIFF")); err != nil {
		return err
	}

	binary.Write(file, binary.LittleEndian, fileSize-8) // File size minus RIFF header size

	if _, err := file.Write([]byte("WAVE")); err != nil {
		return err
	}

	// fmt subchunk
	if _, err := file.Write([]byte("fmt ")); err != nil {
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
	if _, err := file.Write([]byte("data")); err != nil {
		return err
	}
	binary.Write(file, binary.LittleEndian, subChunk2Size)
	return nil
}

func Write(wave []float64, filename string, sampleRate int) error {
	// Create a new file
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	// Write the WAV header
	if err := writeHeader(file, len(wave), sampleRate); err != nil {
		return err
	}

	// Write the audio data
	for _, sample := range wave {
		// Convert the sample from -1.0 <= sample <= 1.0 to 16-bit signed integer
		var intSample int16 = int16(sample * 32767)
		binary.Write(file, binary.LittleEndian, intSample)
	}

	return nil
}
