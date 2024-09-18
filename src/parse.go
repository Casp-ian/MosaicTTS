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

func Parse() LyricsList {
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

	var list = LyricsList{
		nil, nil,
	}

	// filling this shit backwards, so when walking through the linked list it will be in order
	for _, segment := range jsonData.Segments {

		for _, word := range segment.Words {
			list.AddTail(
				&LyricsListEntry{
					fileName,   // Song  string
					word.Start, // Start float32
					word.End,   // End   float32
					word.Text,  // Text  string
					nil,        // Next  *LyricsListEnty
				},
			)

		}
	}
	return list
}
