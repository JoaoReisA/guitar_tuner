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
		//TODO: Handle audio input with FFR
		fmt.Println("Input buffer (first 100 samples):", in[:100]) // Exibe os primeiros 100 valores
	}, utils.SAMPLE_RATE)

	if err != nil {
		fmt.Println("Erro ao abrir stream:", err)
		return
	}

	time.Sleep(5 * time.Second)

	stream.Close()
	portaudio.Terminate()
}
