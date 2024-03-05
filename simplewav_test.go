package simplewav

import (
	"bytes"
	"os"
	"testing"
)

// TestWriteHeader tests the writeHeader function to ensure it writes the correct WAV header.
func TestWriteHeader(t *testing.T) {
	const sampleRate = 44100
	const numSamples = 22050 // 0.5 seconds of audio

	var buf bytes.Buffer
	if err := writeHeader(&buf, numSamples, sampleRate); err != nil {
		t.Fatalf("writeHeader failed: %v", err)
	}

	expectedHeaderSize := 44 // WAV header size
	if buf.Len() != expectedHeaderSize {
		t.Errorf("Expected header size to be %d, got %d", expectedHeaderSize, buf.Len())
	}
}

// TestWrite tests the Write function to ensure it creates a WAV file with the correct format.
func TestWrite(t *testing.T) {
	const sampleRate = 44100
	tempFile, err := os.CreateTemp("", "simplewav_test_*.wav")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(tempFile.Name()) // clean up

	// Generate a short tone for testing
	var wave []float64
	for i := 0; i < sampleRate; i++ {
		wave = append(wave, 0.5)
	}

	if err := Write(wave, tempFile.Name(), sampleRate); err != nil {
		t.Fatalf("Write failed: %v", err)
	}

	// Verify the file content
	fileInfo, err := tempFile.Stat()
	if err != nil {
		t.Fatalf("Failed to stat temp file: %v", err)
	}

	expectedFileSize := int64(44) + int64(len(wave)*2) // header size + data size (2 bytes per sample)
	if fileInfo.Size() != expectedFileSize {
		t.Errorf("Expected file size to be %d, got %d", expectedFileSize, fileInfo.Size())
	}
}
