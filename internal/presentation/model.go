package presentation

import (
	"fmt"
	"guitar_tuner/internal/data/service"
	"guitar_tuner/internal/domain/usecase"
	"guitar_tuner/utils"
	"math"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	currentFreq  float64
	expectedFreq float64
	note         string
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		}

	case freqUpdateMsg:
		m.currentFreq = msg.current
		m.expectedFreq = msg.expected
		m.note = msg.note
	}
	return m, nil
}

func (m model) View() string {
	asciiLogo := `
  ________      .__  __                 ___________                         
 /  _____/ __ __|__|/  |______ _______  \__    ___/_ __  ____   ___________ 
/   \  ___|  |  \  \   __\__  \\_  __ \   |    | |  |  \/    \_/ __ \_  __ \
\    \_\  \  |  /  ||  |  / __ \|  | \/   |    | |  |  /   |  \  ___/|  | \/
 \______  /____/|__||__| (____  /__|      |____| |____/|___|  /\___  >__|   
        \/                    \/                            \/     \/       
	`
	logo := logoStyle.Render(asciiLogo)
	bar := renderBar(m.currentFreq, m.expectedFreq)
	header := headerStyle.Render(fmt.Sprintf("   %s (%.2f Hz) Expected", m.note, m.expectedFreq))
	value := fmt.Sprintf(" %.2f Hz", m.currentFreq)
	footer := footerStyle.Render("[Q] to quit")

	return fmt.Sprintf("%s\n\n%s\n%s %s\n\n%s", logo, header, bar, value, footer)
}

func renderBar(current, expected float64) string {
	const barWidth = 20
	delta := current - expected
	shift := int(math.Round(delta * 2))

	if shift < -barWidth/2 {
		shift = -barWidth / 2
	} else if shift > barWidth/2 {
		shift = barWidth / 2
	}

	left := barWidth/2 + shift
	right := barWidth - left

	return fmt.Sprintf("[%s|%s]", strings.Repeat("<", left), strings.Repeat(">", right))
}

type freqUpdateMsg struct {
	current  float64
	expected float64
	note     string
}

func sendFreqUpdate(p *tea.Program, current, expected float64, note string) {
	p.Send(freqUpdateMsg{current, expected, note})
}

func StartTuner() (tea.Model, error) {
	p := tea.NewProgram(model{})

	go func() {
		_, err := service.OpenAudioInputBufferStreamChannel(func(in []int16) {
			fft := usecase.FFRFromAudioInputBuffer(in)
			dominant := usecase.FindDominantFrequency(fft)
			note := usecase.NoteFromFrequency(dominant)

			sendFreqUpdate(p, note.CurrentFrequency, note.ExpectedFrequency, note.Name.String())
		}, utils.SAMPLE_RATE)

		if err != nil {
			p.Send(freqUpdateMsg{0, 0, "Error"})
		}
	}()

	return p.Run()
}
