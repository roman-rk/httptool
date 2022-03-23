package main

import (
	"testing"
)

func Test_GetContent(t *testing.T) {
	var TestData = []string{"https://go.dev/dl/", "https://www.sports.ru/", "https://www.google.ru/"}

	for _, test_url := range TestData {
		name := test_url
		t.Run(name, func(t *testing.T) {
			content, err := getContent(test_url)
			if err != nil {
				t.Errorf("Error getting content: %v", err)
			}
			if len(content) == 0 {
				t.Errorf("Empty content")
			}
		})
	}
}
