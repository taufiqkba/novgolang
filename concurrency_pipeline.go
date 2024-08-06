package main

import (
	"crypto/md5"
	"fmt"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"time"
)

const totalFile = 3000
const contentLength = 5000

var tempPath = filepath.Join(os.Getenv("TEMP"), "chapter-A.59-pipeline-temp")

// Generate Dummy Files
func randomStringConcurrent(length int) string {
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
		fileName := filepath.Join(tempPath, fmt.Sprintf("file-%d.txt", i))
		content := randomStringConcurrent(contentLength)
		err := os.WriteFile(fileName, []byte(content), os.ModePerm)
		if err != nil {
			log.Println("Error writing file", fileName)
		}

		if i%100 == 0 && i > 0 {
			log.Println(i, "files created")
		}
	}
	log.Printf("%d of total files created", totalFile)
}

// Read all files, find MD5 hash-file, and use the hash for rename files
func proceed() {
	counterTotal := 0
	counterRenamed := 0
	err := filepath.Walk(tempPath, func(path string, info os.FileInfo, err error) error {
		// if there is an error, returned immediatelly
		if err != nil {
			return err
		}

		// if it is a sub directory, return immediatelly
		if info.IsDir() {
			return nil
		}
		counterTotal++

		// read file
		buf, err := os.ReadFile(path)
		if err != nil {
			return err
		}

		// sum it
		sum := fmt.Sprintf("%x", md5.Sum(buf))

		// rename file
		destinationPath := filepath.Join(tempPath, fmt.Sprintf("file-%d.txt", sum))
		err = os.Rename(path, destinationPath)
		if err != nil {
			return err
		}

		counterRenamed++
		return nil
	})

	if err != nil {
		log.Println("ERROR:", err.Error())
	}

	log.Printf("%d/%d files renamed", counterRenamed, counterTotal)
}

func main() {
	fmt.Println("Concurrency Pattern: Pipeline")
	fmt.Println("")

	log.Println("Start")
	start := time.Now()

	// generateFiles()
	// duration := time.Since(start)
	// log.Println("done in", duration.Seconds(), "seconds")

	// second program
	// use hash-file to rename file
	proceed()
	duration := time.Since(start)
	log.Println("Done in: ", duration.Seconds(), "seconds")

}
