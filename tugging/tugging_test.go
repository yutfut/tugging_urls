package tugging

import (
	"fmt"
	"io"
	"os"
	"reflect"
	"sync"
	"testing"
)

func TestTuggingUrls(t *testing.T) {
	file, err := os.OpenFile("test.txt", os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		t.Errorf("error")
		return
	}

	wg := &sync.WaitGroup{}

	wg.Add(1)
	go TuggingUrls(file, "https://vk.com", wg)

	waitChan := make(chan int, 1)
	go Wait(wg, waitChan)
	_, _ = <-waitChan

	file.Close()

	file, err = os.OpenFile("test.txt", os.O_RDONLY, 0666)
	if err != nil {
		t.Errorf("error")
		return
	}
	defer os.Remove("test.txt")
	defer file.Close()

	data := make([]byte, 64)
	var test string

	for {
		n, err := file.Read(data)
		if err == io.EOF { // если конец файла
			break // выходим из цикла
		}
		test = string(data[:n])
		fmt.Print(string(data[:n]))
	}

	if !reflect.DeepEqual(test, "https://vk.com - 200 OK\n") {
		t.Errorf("error")
		return
	}
}
