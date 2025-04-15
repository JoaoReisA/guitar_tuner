package usecase

import (
	domain "guitar_tuner/internal/domain/entities"
	"math"
	"sort"
)

func getFrequenciesToNoteTable() map[domain.NoteSymbol][]float64 {
	notesMap := map[domain.NoteSymbol][]float64{
		domain.C:      {32.70, 65.41, 130.81, 261.63, 523.25},
		domain.CSharp: {34.65, 69.30, 138.59, 277.18, 554.37},
		domain.D:      {36.71, 73.42, 146.82, 293.66, 587.33},
		domain.DSharp: {38.89, 77.78, 155.56, 311.13, 622.25},
		domain.E:      {41.20, 82.41, 164.82, 329.63, 659.25},
		domain.F:      {43.65, 87.31, 174.61, 349.23, 698.46},
		domain.FSharp: {46.25, 92.50, 185.0, 369.99, 739.99},
		domain.G:      {49, 98, 196, 392, 783.99},
		domain.GSharp: {51.91, 103.83, 207.65, 415.30, 830.61},
		domain.A:      {55, 110.0, 220, 440, 880},
		domain.ASharp: {58.27, 116.54, 233.08, 466.16, 932.33},
		domain.B:      {61.74, 123.47, 246.94, 493.88, 987.77},
	}

	return notesMap
}

func NoteFromFrequency(frequency float64) domain.Note {

	notesMap := getFrequenciesToNoteTable()
	sortedFrequencies := buildSortedFrequencies(notesMap)
	closestFrequency := findClosest(sortedFrequencies, frequency)

	for i, v := range notesMap {
		for _, freq := range v {
			if freq == closestFrequency {
				return domain.Note{
					Name:              i,
					CurrentFrequency:  frequency,
					ExpectedFrequency: freq,
				}
			}
		}
	}

	return domain.Note{
		Name:             domain.NoteSymbol(0),
		CurrentFrequency: frequency,
	}
}

func buildSortedFrequencies(notesMap map[domain.NoteSymbol][]float64) []float64 {
	var all []float64
	for _, freqs := range notesMap {
		all = append(all, freqs...)
	}
	sort.Float64s(all)
	return all
}

func findClosest(sorted []float64, target float64) float64 {
	i := sort.SearchFloat64s(sorted, target)
	if i == 0 {
		return sorted[0]
	}
	if i == len(sorted) {
		return sorted[len(sorted)-1]
	}
	if math.Abs(sorted[i-1]-target) < math.Abs(sorted[i]-target) {
		return sorted[i-1]
	}
	return sorted[i]
}
