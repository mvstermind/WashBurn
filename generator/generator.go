package main

import (
	"fmt"
	"math/rand"
)

type Gen struct {
	Bpm    int
	Key    map[int]string
	Chords [7]string
}

func (b *Gen) GetBpm() int {
	b.Bpm = rand.Intn(100) + 60
	return b.Bpm
}

func (b *Gen) GetKey() string {

	b.Key = map[int]string{
		1: "C", 2: "C#", 3: "D", 4: "D#", 5: "E", 6: "F", 7: "F#", 8: "G", 9: "G#",
		10: "A", 11: "A#", 12: "B",
	}
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

func main() {
	gned := Gen{}
	bpm := gned.GetBpm()
	fmt.Printf("Bpm: %v\n", bpm)

	key := gned.GetKey()
	fmt.Printf("Scale: %v", key)

}
