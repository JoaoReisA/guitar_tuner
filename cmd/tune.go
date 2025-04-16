package cmd

import (
	"fmt"
	"guitar_tuner/internal/presentation"

	"github.com/spf13/cobra"
)

var tuneCmd = &cobra.Command{
	Use:   "tune",
	Short: "Start Guitar Tuner",
	Run: func(cmd *cobra.Command, args []string) {
		if _, err := presentation.StartTuner(); err != nil {
			fmt.Println("Error on start tuner:", err)
		}
	},
}
