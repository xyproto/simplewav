# simplewav

Given a slice of float64 and a sample rate (like 44100), write the audio to a WAV file.

The example in `cmd/sawtooth` generates a 1 second long audio file containing a sawtooth wave at 220Hz.

## Example use

```go
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
```

## General info

License: BSD-3
