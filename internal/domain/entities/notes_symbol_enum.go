package domain

type NoteSymbol int

const (
	C NoteSymbol = iota
	CSharp
	D
	DSharp
	E
	F
	FSharp
	G
	GSharp
	A
	ASharp
	B
)

func (s NoteSymbol) String() string {
	return [...]string{"C", "C#", "D", "D#", "E", "F", "F#", "G", "G#", "A", "A#", "B"}[s]
}
