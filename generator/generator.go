package generator

import (
	"fmt"
	"math/rand"
)

type Gen struct {
	Bpm int
}

func (b *Gen) GetBpm() int {
	b.Bpm = rand.Intn(100) + 60 //hardcoded numbers so you won't get like 3 bpm
	return b.Bpm
}

func New() {
	gned := Gen{}
	bpm := gned.GetBpm()
	fmt.Printf("Bpm: %v\n", bpm)

}
