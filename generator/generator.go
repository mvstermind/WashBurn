package main

import (
	"fmt"
	"math/rand"
)

type Gen struct {
	Bpm    int
	Key    []string
	Scale  []string
	WScale map[string]string
	Chords [7]string
}

func (b *Gen) GetBpm() int {
	b.Bpm = rand.Intn(100) + 60
	return b.Bpm
}

func main() {
	gned := Gen{} // Bpm is updated
	bpm := gned.GetBpm()
	fmt.Printf("Get Bpm %v", bpm)

}
