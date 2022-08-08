package tools

import (
	"astro"
	"encoding/json"
	"io"
	"net/http"
	"sync"

	"github.com/sirupsen/logrus"
)

func MakeRequest(wg *sync.WaitGroup, url string, ch chan<- astro.ResponsePair) {
	defer wg.Done()
	resp, err := http.Get(url)
	if err != nil {
		ch <- astro.ResponsePair{Body: nil, Error: err}
		return
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		ch <- astro.ResponsePair{Body: nil, Error: err}
		return
	}

	ch <- astro.ResponsePair{Body: body, Error: nil}
}

func MakeRequests(urls ...string) map[string]astro.ResponsePair {
	ch := make(chan astro.ResponsePair, len(urls))
	wg := sync.WaitGroup{}
	for _, url := range urls {
		wg.Add(1)
		go MakeRequest(&wg, url, ch)
	}

	result := make(map[string]astro.ResponsePair)
	for _, u := range urls {
		result[u] = <-ch
	}

	return result
}

func SendResponse(w http.ResponseWriter, resp astro.Response) {
	raw, _ := json.Marshal(resp)

	w.WriteHeader(resp.Status)
	n, err := w.Write(raw)
	if err != nil || n != len(raw) {
		logrus.Print("err: %q, sent %d/%d bytes", err, n, len(raw))
	}
}
