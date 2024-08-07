package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"time"
)

const totalFile = 3000
const contentLength = 5000

var tempPath = filepath.Join(os.Getenv("TEMP"), "chapter-A.60-worker-pool")

func randomString(length int) string {
	randomizer := rand.New(rand.NewSource(time.Now().Unix()))
	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	b := make([]rune, length)
	for i := range b {
		b[i] = letters[randomizer.Intn(len(letters))]
	}
	return string(b)
}

func generateFiles() {
	os.RemoveAll(tempPath)
	os.MkdirAll(tempPath, os.ModePerm)

	for i := 0; i < totalFile; i++ {
		fileName := filepath.Join(tempPath, fmt.Sprintf("file-%d", i))
		content := randomString(contentLength)
		err := os.WriteFile(fileName, []byte(content), os.ModePerm)
		if err != nil {
			log.Println("Error writing file", fileName)
		}
		log.Println(i, "files created")
	}
	log.Printf("%d of total files created", totalFile)
}

func main() {
	fmt.Println("Concurrency Pattern: Simplified Fan-In Fan-Out Pipeline")

	log.Println("start")
	start := time.Now()

	generateFiles()

	duration := time.Since(start)
	log.Println("Done in", duration.Seconds(), "seconds")
}
