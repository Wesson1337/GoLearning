package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func main() {
	rpsLimit := 50000
	semaphore := make(chan struct{}, rpsLimit)

	ticker := time.NewTicker(time.Second / time.Duration(rpsLimit))
	defer ticker.Stop()
	req, err := http.NewRequest("GET", "http://localhost:8080", nil)
	if err != nil {
		log.Fatalf("Failed to create request: %v", err)
	}

	messageCount := 0
	startTime := time.Now()

	for {
		go func(r *http.Request) {
			semaphore <- struct{}{}
			sendRequest(r)
			<-semaphore
		}(req)

		messageCount++
		if messageCount%10000 == 0 {
			elapsed := time.Since(startTime)
			log.Printf("Processed %d messages in %v. Current RPS: %f", messageCount, elapsed, float64(10000)/elapsed.Seconds())
			startTime = time.Now()
		}
	}
}

func sendRequest(req *http.Request) {
	http.DefaultTransport.(*http.Transport).MaxConnsPerHost = 1000

	client := &http.Client{
		Transport: http.DefaultTransport,
		Timeout:   10 * time.Second,
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("Failed to send request: %v", err)
		return
	}
	ioutil.ReadAll(resp.Body)
	resp.Body.Close()
}
