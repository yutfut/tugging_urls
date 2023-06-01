package main

import (
	"bufio"
	"fmt"
	"github.com/yutfut/tugging_urls/tugging"
	"os"
	"sync"
)

func main() {
	ReadFile, err := os.OpenFile("urls.txt", os.O_RDONLY, 0666)
	if err != nil {
		panic(err)
	}
	defer ReadFile.Close()

	WriteFile, err := os.OpenFile("result.txt", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}
	defer WriteFile.Close()

	fileScanner := bufio.NewScanner(ReadFile)

	wg := &sync.WaitGroup{}

	for fileScanner.Scan() {
		wg.Add(1)
		go tugging.TuggingUrls(WriteFile, fileScanner.Text(), wg)
	}

	if err = fileScanner.Err(); err != nil {
		fmt.Printf("Error while reading ReadFile: %s", err)
	}

	WaitChan := make(chan int, 1)

	go tugging.Wait(wg, WaitChan)

	_, _ = <-WaitChan
}
