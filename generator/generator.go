package generator

import (
	"fmt"
	"math/rand"
)

type Gen struct {
	Bpm    int
	Key    [12]string
	Chords [7]string
}

func (b *Gen) GetBpm() int {
	b.Bpm = rand.Intn(100) + 60 //hardcoded numbers so you won't get like 3 bpm
	return b.Bpm
}

func (b *Gen) GetKey() string {

	// i really used map on this thing b4...
	b.Key = [12]string{"C", "C#", "D", "D#", "E", "F", "F#", "G", "G#", "A", "A#", "B"}
	key := rand.Intn(12) + 1
	gnedKey := b.Key[key]

	scale := [2]string{
		"Minor", "Major",
	}

	s := rand.Intn(2)
	gnedScale := scale[s]

	output := fmt.Sprintf("%v %v\n", gnedKey, gnedScale)

	return output
}

func (b *Gen) GetChords() []string {
	var chords []string // create new random chor progression

	b.Chords = [7]string{"I", "II", "III", "IV", "V", "VI", "VII"}
	ranI := rand.Intn(4) + 1 // choose amount of chords

	for i := 0; i < ranI; i++ {
		r := rand.Intn(7)
		chords = append(chords, b.Chords[r])
	}
	return chords

}

func New() {
	gned := Gen{}
	bpm := gned.GetBpm()
	fmt.Printf("Bpm: %v\n", bpm)

	key := gned.GetKey()
	fmt.Printf("Scale: %v", key)

	chords := gned.GetChords()
	fmt.Printf("Chords: %v", chords)

}
