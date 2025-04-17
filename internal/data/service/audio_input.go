package service

import (
	"fmt"
	"guitar_tuner/utils"

	"github.com/gordonklaus/portaudio"
)

func OpenAudioInputBufferStreamChannel(streamCallback func(in []int16), sampleRate float64) (*portaudio.Stream, error) {
	err := portaudio.Initialize()
	if err != nil {
		return nil, fmt.Errorf("falha ao inicializar PortAudio: %v", err)
	}

	inputBuffer := make([]int16, utils.BUFFER_SIZE)

	portAudioCallback := func(in []int16, out []int16) {
		copy(inputBuffer, in)
		streamCallback(inputBuffer)
	}

	stream, err := portaudio.OpenDefaultStream(1, 0, sampleRate, len(inputBuffer), portAudioCallback)
	if err != nil {
		return nil, fmt.Errorf("Error on open Stream: %v", err)
	}

	err = stream.Start()
	if err != nil {
		return nil, fmt.Errorf("Error on start stream: %v", err)
	}

	return stream, nil
}
