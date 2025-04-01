package main

import (
	"fmt"
	"guitar_tuner/cmd"
	"guitar_tuner/internal/domain/service"
	"guitar_tuner/utils"
	"time"

	"github.com/gordonklaus/portaudio"
)

func main() {
	cmd.Execute()

	stream, err := service.OpenAudioInputBufferStreamChannel(func(in []int16) {
		fftResult := service.FFRFromAudioInputBuffer(in)
		service.FindDominantFrequency(fftResult)
		// fmt.Println("Dominant frequency:", frequency)
	}, utils.SAMPLE_RATE)

	if err != nil {
		fmt.Println("Erro ao abrir stream:", err)
		return
	}

	time.Sleep(60 * time.Second)

	stream.Close()
	portaudio.Terminate()
}
