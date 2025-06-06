package usecase

import (
	"guitar_tuner/utils"
	"math"
	"math/cmplx"

	"github.com/mjibson/go-dsp/fft"
)

func normalizeBuffer(inputBuffer []int16) []float64 {
	maxVal := float64(1 << 15) // Normalize to [-1, 1]
	normalized := make([]float64, len(inputBuffer))
	for i, v := range inputBuffer {
		normalized[i] = float64(v) / maxVal
	}
	return normalized
}

func applyHanningWindow(inputBuffer []float64) []float64 {
	windowed := make([]float64, len(inputBuffer))
	N := float64(len(inputBuffer))
	for i := range inputBuffer {
		windowed[i] = inputBuffer[i] * 0.5 * (1 - math.Cos(2*math.Pi*float64(i)/N))
	}
	return windowed
}

func FFRFromAudioInputBuffer(inputBuffer []int16) []complex128 {
	audioData := normalizeBuffer(inputBuffer)
	windowedData := applyHanningWindow(audioData)
	fftResult := fft.FFTReal(windowedData)
	return fftResult
}

func FindDominantFrequency(fftResult []complex128) float64 {
	var maxMag float64
	var maxIndex int

	magnitudeSpectrum := make([]float64, len(fftResult)/2)

	for i := 1; i < len(magnitudeSpectrum); i++ {
		magnitude := cmplx.Abs(fftResult[i])
		magnitudeSpectrum[i] = magnitude

		frequency := float64(i) * utils.SAMPLE_RATE / float64(utils.BUFFER_SIZE)

		if magnitude > maxMag && frequency >= utils.MIN_FREQUENCY && frequency <= utils.MAX_FREQUENCY {
			maxMag = magnitude
			maxIndex = i
		}
	}

	detectedFrequency := float64(maxIndex) * utils.SAMPLE_RATE / float64(utils.BUFFER_SIZE)

	// check if the detected frequency is a harmonic for a lower frequency
	for i := maxIndex / 2; i > 1; i-- {
		if magnitudeSpectrum[i] > (maxMag * 0.3) { // fundamental have at least 30% of Harmonics
			detectedFrequency = float64(i) * utils.SAMPLE_RATE / float64(utils.BUFFER_SIZE)
			break
		}
	}

	return detectedFrequency
}
