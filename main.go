package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func che(file io.Writer, url string) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(url + " " + err.Error())
	}

	_, err = fmt.Fprintln(file, url+" - "+resp.Status)
	if err != nil {
		fmt.Println(url + " " + err.Error())
	}
}

func main() {
	file, err := os.OpenFile("urls.txt", os.O_RDONLY, 0666)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	f2, err := os.OpenFile("result.txt", os.O_WRONLY, 0666)
	if err != nil {
		panic(err)
	}
	defer f2.Close()

	fileScanner := bufio.NewScanner(file)

	for fileScanner.Scan() {
		go che(f2, fileScanner.Text())
	}

	if err = fileScanner.Err(); err != nil {
		log.Fatalf("Error while reading file: %s", err)
	}
	time.Sleep(2 * time.Second)
}
