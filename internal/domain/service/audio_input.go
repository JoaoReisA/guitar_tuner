package service

import (
	"fmt"
	"log"
	"time"

	"github.com/gordonklaus/portaudio"
)

func ReceiveAudioInput() {
	portaudio.Initialize()
	defer portaudio.Terminate()

	inputBuffer := make([]int16, 44100) // Buffer for 1 second of audio

	stream, err := portaudio.OpenDefaultStream(1, 0, 44100, len(inputBuffer), inputBuffer)
	if err != nil {
		log.Fatal(err)
	}
	defer stream.Close()

	fmt.Println("Recording...")
	stream.Start()
	time.Sleep(time.Second)
	stream.Stop()

	fmt.Println(inputBuffer)
	fmt.Println("Recording finished.")
	// Process `inputBuffer` (e.g., save to file or analyze)
}
