package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"sync"
)

func main() {
	// go to the parallel processing
	if os.Args[1] == "parallel" {
		runParallel(os.Args[2:])
		return
	}
	// go to sequential processing
	run(os.Args[1:])
}

// Run urls processing sequentially
func run(urls []string) {
	for _, url := range urls {
		content, err := getContent(url)
		if err != nil {
			fmt.Printf("Error on processing %s url: %v", url, err)
			return
		}

		fmt.Printf("%s %s\n", url, calculateHash(content))
	}
}

// Run urls processing in parallel
func runParallel(urls []string) {
	wg := sync.WaitGroup{}

	for _, url := range urls {
		wg.Add(1)

		go func(url string) {
			defer wg.Done()

			content, err := getContent(url)
			if err != nil {
				fmt.Printf("Error on processing %s url: %v", url, err)
				return
			}

			fmt.Printf("%s %s\n", url, calculateHash(content))
		}(url)
	}

	wg.Wait()
}

func getContent(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", fmt.Errorf("GET error: %v", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("Status error: %v", resp.StatusCode)
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("Read body: %v", err)
	}

	return string(data), nil
}

func calculateHash(content string) string {
	hash := md5.Sum([]byte(content))
	return hex.EncodeToString(hash[:])
}
