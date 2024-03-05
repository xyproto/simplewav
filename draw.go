package simplewav

import (
	"image"
	"image/color"
	"image/png"
	"os"
)

// Draw creates a PNG image of the waveform represented by the audio samples.
// The waveform image is plotted with the X-axis representing time and the Y-axis amplitude.
func Draw(wave []float64, filename string, sampleRate int) error {
	width := 1024 // Width of the output image
	height := 256 // Height of the output image
	img := image.NewRGBA(image.Rect(0, 0, width, height))

	// Determine the number of samples per pixel to scale the waveform to the image width.
	samplesPerPixel := len(wave) / width
	if samplesPerPixel == 0 {
		samplesPerPixel = 1 // Avoid division by zero for very short samples
	}

	// Set a more visually appealing background and waveform color
	bgColor := color.RGBA{230, 230, 230, 255}  // Light grey background
	waveColor := color.RGBA{30, 144, 255, 255} // Dodger blue waveform

	// Fill the background
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			img.Set(x, y, bgColor)
		}
	}

	// Draw the waveform
	centerY := height / 2
	for x := 0; x < width; x++ {
		startSample := x * samplesPerPixel
		endSample := (x + 1) * samplesPerPixel
		if endSample > len(wave) {
			endSample = len(wave)
		}

		// Find the min and max amplitudes for this segment
		minSample, maxSample := 0.0, 0.0
		for _, sample := range wave[startSample:endSample] {
			if sample < minSample {
				minSample = sample
			}
			if sample > maxSample {
				maxSample = sample
			}
		}

		// Scale the min/max to the image height
		minY := centerY + int(minSample*float64(centerY-1))
		maxY := centerY + int(maxSample*float64(centerY-1))
		if minY < 0 {
			minY = 0
		}
		if maxY >= height {
			maxY = height - 1
		}

		// Ensure minY <= maxY for the loop
		if minY > maxY {
			minY, maxY = maxY, minY
		}

		// Draw a vertical line for this segment of the waveform
		for y := minY; y <= maxY; y++ {
			img.Set(x, y, waveColor)
		}
	}

	// Save the image to disk
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	return png.Encode(file, img)
}
