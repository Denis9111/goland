package gadged

import "fmt"

type TapePlayer struct {
	Batteries string
}

type Player interface {
	Play(string)
	Stop()
}

func (t TapePlayer) Play(song string) {
	fmt.Println("Playing", song)
}

func (t TapePlayer) Stop() {
	fmt.Println("Stopping")
}

type TapeRecorder struct {
	Microfones string
}

func (t TapeRecorder) Play(song string) {
	fmt.Println("Playing", song)
}

func (t TapeRecorder) Record(song string) {
	fmt.Println("Recording", song)
}

func (t TapeRecorder) Stop() {
	fmt.Println("Stopping")
}
