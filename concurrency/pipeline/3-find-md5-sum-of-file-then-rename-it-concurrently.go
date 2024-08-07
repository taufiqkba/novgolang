package main

import (
	"crypto/md5"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sync"
	"time"
)

var tempPathPipeline = filepath.Join(os.Getenv("TEMP"), "chapter-A.59-pipeline-temp")

type FileInfo struct {
	FilePath  string //file location
	Content   []byte //file content
	Sum       string //md5 sum of content
	IsRenamed bool   //indicator wheter the particular file is renamed already or not
}

// Pipeline 1: Read file
func readFiles() <-chan FileInfo {
	chanOut := make(chan FileInfo)

	go func() {
		err := filepath.Walk(tempPathPipeline, func(path string, info os.FileInfo, err error) error {
			// if there is an error return err
			if err != nil {
				return err
			}

			// if it is a sub directory, return immediately
			if info.IsDir() {
				return nil
			}

			buf, err := os.ReadFile(path)
			if err != nil {
				return err
			}

			chanOut <- FileInfo{
				FilePath: path,
				Content:  buf,
			}
			return nil
		})

		if err != nil {
			log.Println("ERROR: ", err.Error())
		}
		close(chanOut)
	}()
	return chanOut
}

// Pipeline 2: MD5 hash file content
func getSum(chanIn <-chan FileInfo) <-chan FileInfo {
	chanOut := make(chan FileInfo)

	go func() {
		for fileInfo := range chanIn {
			fileInfo.Sum = fmt.Sprintf("%x", md5.Sum(fileInfo.Content))
			chanOut <- fileInfo
		}
		close(chanOut)
	}()

	return chanOut
}

// merge md5sum
func mergeChanFileInfo(chanInMany ...<-chan FileInfo) <-chan FileInfo {
	wg := new(sync.WaitGroup)
	chanOut := make(chan FileInfo)

	wg.Add(len(chanInMany))
	for _, eachChan := range chanInMany {
		go func(eachChan <-chan FileInfo) {
			for eachChanData := range eachChan {
				chanOut <- eachChanData
			}
			wg.Done()
		}(eachChan)
	}

	go func() {
		wg.Wait()
		close(chanOut)
	}()

	return chanOut
}

// rename files
func rename(chanIn <-chan FileInfo) <-chan FileInfo {
	chanOut := make(chan FileInfo)

	go func() {
		for fileInfo := range chanIn {
			newPath := filepath.Join(tempPathPipeline, fmt.Sprintf("file-%s.txt", fileInfo.Sum))
			err := os.Rename(fileInfo.FilePath, newPath)
			fileInfo.IsRenamed = err == nil
			chanOut <- fileInfo
		}
		close(chanOut)
	}()
	return chanOut
}

func main() {
	fmt.Println("Concurrency pattern: Pipeline")
	fmt.Println("")

	log.Println("start")
	start := time.Now()

	// pipeline 1: loop all files and read it
	chanFileContent := readFiles()

	// pipeline 2: Calculate md5sum
	chanFileSum1 := getSum(chanFileContent)
	chanFileSum2 := getSum(chanFileContent)
	chanFileSum3 := getSum(chanFileContent)
	chanFileSum := mergeChanFileInfo(chanFileSum1, chanFileSum2, chanFileSum3)

	// pipeline 3: rename files
	chanRename1 := rename(chanFileSum)
	// chanRename2 := rename(chanFileSum)
	// chanRename3 := rename(chanFileSum)
	// chanRename4 := rename(chanFileSum)
	// chanRename5 := rename(chanFileSum)
	// chanRename6 := rename(chanFileSum)
	chanRename := mergeChanFileInfo(chanRename1)

	counterRenamed := 0
	// pipeline 4: print output
	counterTotal := 0
	for fileInfo := range chanRename {
		if fileInfo.IsRenamed {
			counterRenamed++
		}
		counterTotal++
	}
	log.Printf("%d/%d files renamed", counterRenamed, counterTotal)

	duration := time.Since(start)
	log.Println("Done in", duration.Seconds(), "Seconds")
}
