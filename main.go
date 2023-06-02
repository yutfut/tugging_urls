package main

import (
	"bufio"
	"fmt"
	"github.com/yutfut/tugging_urls/tugging"
	"os"
	"sync"
)

func main() {
	readFile, err := os.OpenFile("urls.txt", os.O_RDONLY, 0666)
	if err != nil {
		panic(err)
	}
	defer readFile.Close()

	writeFile, err := os.OpenFile("result.txt", os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}
	defer writeFile.Close()

	fileScanner := bufio.NewScanner(readFile)

	wg := &sync.WaitGroup{}

	for fileScanner.Scan() {
		wg.Add(1)
		go tugging.TuggingUrls(writeFile, fileScanner.Text(), wg)
	}

	if err = fileScanner.Err(); err != nil {
		fmt.Printf("Error while reading readFile: %s", err)
	}

	waitChan := make(chan int, 1)

	go tugging.Wait(wg, waitChan)

	_, _ = <-waitChan
}
