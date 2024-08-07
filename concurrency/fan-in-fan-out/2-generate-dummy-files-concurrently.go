package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"sync"
	"time"
)

const totalFile2 = 3000
const contentLength2 = 5000

var tempPath2 = filepath.Join(os.Getenv("TEMP"), "chapter-A.60-worker-pool")

type FileInfo struct {
	Index       int
	FileName    string
	WorkerIndex int
	Err         error
}

func randomString2(length int) string {
	randomizer := rand.New(rand.NewSource(time.Now().Unix()))
	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	b := make([]rune, length)
	for i := range b {
		b[i] = letters[randomizer.Intn(len(letters))]
	}
	return string(b)
}

func generateFiles2() {
	os.RemoveAll(tempPath2)
	os.MkdirAll(tempPath2, os.ModePerm)

	// pipeline 1: job distribution
	chanFileIndex := generateFileIndexes()

	// pipeline 2: the main logic (creating files)
	createFilesWorker := 100
	chanFileResult := createFiles2(chanFileIndex, createFilesWorker)

	// track and print output
	counterTotal := 0
	counterSuccess := 0
	for fileResult := range chanFileResult {
		if fileResult.Err != nil {
			log.Printf("error creating file %s. stack trace: %s", fileResult.FileName, &fileResult.Err)
		} else {
			counterSuccess++
		}
		counterTotal++
	}
	log.Printf("%d/%d of total files created", counterSuccess, counterTotal)
}

func generateFileIndexes() <-chan FileInfo {
	chanOut := make(chan FileInfo)

	go func() {
		for i := 0; i < totalFile2; i++ {
			chanOut <- FileInfo{
				Index:    i,
				FileName: fmt.Sprintf("file-%d.txt", i),
			}
		}
		close(chanOut)
	}()

	return chanOut
}

func createFiles2(chanIn <-chan FileInfo, numberOfWorkers int) <-chan FileInfo {
	chanOut := make(chan FileInfo)

	// wait group to control workers
	wg := new(sync.WaitGroup)

	// allocate N of workers
	wg.Add(numberOfWorkers)

	go func() {

		// dispatch N workers
		for workerIndex := 0; workerIndex < numberOfWorkers; workerIndex++ {
			go func(workerIndex int) {

				// listen to `chanIn` channel for incoming jobs
				for job := range chanIn {

					// do the job
					filePath := filepath.Join(tempPath2, job.FileName)
					content := randomString2(contentLength2)
					err := os.WriteFile(filePath, []byte(content), os.ModePerm)

					log.Println("worker", workerIndex, "working on", job.FileName, "file generation")

					// construct the job's result, and send it to `chanOut`
					chanOut <- FileInfo{
						FileName:    job.FileName,
						WorkerIndex: workerIndex,
						Err:         err,
					}
				}
				// if `chanIn` is closed, and the remaining jobs are finished,
				// only then we mark the worker are completed
				wg.Done()
			}(workerIndex)
		}
	}()
	// wait until `chanIn` closed and then all workers are done,
	// because right after that - we need to close the `chanOut` channel.
	go func() {
		wg.Wait()
		close(chanOut)
	}()
	return chanOut
}

func main() {
	log.Println("Start")
	start := time.Now()

	generateFiles2()

	duration := time.Since(start)
	log.Println("Done in", duration.Seconds(), "seconds")
}
