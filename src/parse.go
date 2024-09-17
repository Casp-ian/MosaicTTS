package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type WhisperTimestampedJson struct {
	Text     string                      `json:"name"`
	Segments []WhisperTimestampedSegment `json:"segments"`
}

type WhisperTimestampedSegment struct {
	Text       string                    `json:"text"`
	Start      float32                   `json:"start"`
	End        float32                   `json:"end"`
	Confidence float32                   `json:"confidence"`
	Words      []WhipserTimestampedWords `json:"words"`
}

type WhipserTimestampedWords struct {
	Text       string  `json:"text"`
	Start      float32 `json:"start"`
	End        float32 `json:"end"`
	Confidence float32 `json:"confidence"`
}

func Parse() LyricsListEnty {
	var fileName = "../lyrics/hello.json"

	jsonFile, err := os.Open(fileName)
	defer jsonFile.Close()

	if err != nil {
		fmt.Printf("cant open %v...\n", fileName)
	}

	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		fmt.Printf("invalid json in %v...\n", fileName)
	}

	var jsonData WhisperTimestampedJson
	json.Unmarshal(byteValue, &jsonData)

	var head *LyricsListEnty

	// filling this shit backwards, so when walking through the linked list it will be in order
	for i := len(jsonData.Segments) - 1; i >= 0; i-- {
		for j := len(jsonData.Segments[i].Words) - 1; j >= 0; j-- {
			head = &LyricsListEnty{
				fileName,                            // Song  string
				jsonData.Segments[i].Words[j].Start, // Start float32
				jsonData.Segments[i].Words[j].End,   // End   float32
				jsonData.Segments[i].Words[j].Text,  // Text  string
				head,                                // Next  *LyricsListEnty
			}
		}
	}
	return *head
}
