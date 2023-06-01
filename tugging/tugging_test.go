package tugging

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"sync"
	"testing"
)

func TestTuggingUrls(t *testing.T) {
	File, err := os.OpenFile("test.txt", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		t.Errorf("error")
		return
	}
	defer os.Remove("test.txt")
	defer File.Close()

	wg := &sync.WaitGroup{}

	wg.Add(1)
	go TuggingUrls(File, "https://vk.com", wg)

	WaitChan := make(chan int, 1)
	go Wait(wg, WaitChan)
	_, _ = <-WaitChan

	fileScanner := bufio.NewScanner(File)
	for fileScanner.Scan() {
		fmt.Println(fileScanner)
		break
	}

	if !reflect.DeepEqual(fileScanner.Text(), "https://vk.com - 200 OK") {
		t.Errorf("error")
		return
	}
}
