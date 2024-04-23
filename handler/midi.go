package handler

import (
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
)

func Midi() (string, *os.File) {
	wd, _ := os.Getwd()
	K := "CHORD PROGRESSIONS/"
	Keys := filepath.Join(wd, K)

	dir, err := os.Open(Keys)
	if err != nil {
		fmt.Println("Error opening directory: ", err)
		return "", nil
	}
	defer dir.Close()

	files, err := dir.Readdir(0)
	if err != nil {
		fmt.Println("Error reading directory content: ", err)
		return "", nil
	}

	fileCount := len(files)
	if fileCount == 0 {
		fmt.Println("No files found in directory")
		return "", nil
	}

	randNum := rand.Intn(fileCount)

	for i, file := range files {
		if randNum == i {
			choice := file.Name()
			filePath := filepath.Join(Keys, choice)
			file, err := os.Open(filePath)
			if err != nil {
				fmt.Println("Error opening file: ", err)
				return "", nil
			}
			return choice, file
		}
	}

	return "", nil
}
