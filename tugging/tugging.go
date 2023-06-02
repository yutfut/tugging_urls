package tugging

import (
	"fmt"
	"io"
	"net/http"
	"sync"
)

func TuggingUrls(file io.Writer, url string, wg *sync.WaitGroup) {
	defer func(wg *sync.WaitGroup) {
		wg.Done()
	}(wg)

	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("%s %s", url, err.Error())
	}

	_, err = fmt.Fprintln(file, fmt.Sprintf("%s - %s", url, resp.Status))
	if err != nil {
		fmt.Printf("%s %s", url, err.Error())
	}
}

func Wait(wg *sync.WaitGroup, in chan int) {
	wg.Wait()
	close(in)
}
