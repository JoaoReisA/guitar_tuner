package main

import (
	"guitar_tuner/cmd"
	"guitar_tuner/internal/domain/service"
)

func main() {
	cmd.Execute()
	service.ReceiveAudioInput()
}
