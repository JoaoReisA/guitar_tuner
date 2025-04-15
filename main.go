package main

import (
	"fmt"
	"guitar_tuner/cmd"
	"guitar_tuner/internal/data/service"
	"guitar_tuner/internal/domain/usecase"
	"guitar_tuner/utils"
	"time"

	"github.com/gordonklaus/portaudio"
)

func main() {
	cmd.Execute()

	stream, err := service.OpenAudioInputBufferStreamChannel(func(in []int16) {
		fftResult := usecase.FFRFromAudioInputBuffer(in)
		dominantFrequency := usecase.FindDominantFrequency(fftResult)
		note := usecase.NoteFromFrequency(dominantFrequency)
		fmt.Println("Current Note:", note.Name.String(), note.CurrentFrequency, note.ExpectedFrequency)
	}, utils.SAMPLE_RATE)

	if err != nil {
		fmt.Println("Erro ao abrir stream:", err)
		return
	}

	time.Sleep(60 * time.Second)

	stream.Close()
	portaudio.Terminate()
}
