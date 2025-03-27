package service

import (
	"fmt"

	"github.com/gordonklaus/portaudio"
)

// OpenAudioInputBufferStreamChannel abre uma stream de áudio e executa um callback para cada buffer capturado
func OpenAudioInputBufferStreamChannel(streamCallback func(in []int16), sampleRate float64) (*portaudio.Stream, error) {
	err := portaudio.Initialize()
	if err != nil {
		return nil, fmt.Errorf("falha ao inicializar PortAudio: %v", err)
	}

	inputBuffer := make([]int16, int(sampleRate))

	portAudioCallback := func(in []int16, out []int16) {
		copy(inputBuffer, in)
		streamCallback(inputBuffer)
	}

	stream, err := portaudio.OpenDefaultStream(1, 0, sampleRate, len(inputBuffer), portAudioCallback)
	if err != nil {
		return nil, fmt.Errorf("erro ao abrir stream: %v", err)
	}

	fmt.Println("Recording...")
	err = stream.Start()
	if err != nil {
		return nil, fmt.Errorf("erro ao iniciar stream: %v", err)
	}

	return stream, nil
}
