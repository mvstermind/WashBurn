package handler

import (
	"fmt"
	"math/rand"
	"os"
)

func Midi() (string, *os.File) {

	Keys := "/Users/kneehead/work/github-bot/CHORD PROGRESSIONS/"

	dir, err := os.Open(Keys)
	if err != nil {
		fmt.Println("Error opening directory: ", err)
		return "", nil
	}

	defer dir.Close()

	files, err := dir.Readdir(0)
	if err != nil {
		fmt.Println("Error reading directory contetn: ", err)
		return "", nil
	}

	fileCount := 0
	for range files {
		fileCount++
	}
	randNum := rand.Intn(fileCount)

	for i, file := range files {
		if randNum == i {
			choice := file.Name()
			file, err := os.Open(Keys + choice)
			if err != nil {
				fmt.Println("Error opening file: ", err)
				return "", nil
			}
			return choice, file
		}
	}

	return "", nil
}
