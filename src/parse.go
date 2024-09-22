package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
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

func Parse(lyricsDir string, lyricsFiles []string) LyricsList {
	var result = LyricsList{
		nil, nil,
	}

	for _, file := range lyricsFiles {
		getLyrics(&result, file)
	}

	return result
}

func getLyrics(result *LyricsList, file string) {

	jsonFile, err := os.Open("../lyrics/" + file)
	defer jsonFile.Close()

	if err != nil {
		fmt.Printf("cant open %v...\n", file)
		return
	}

	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		fmt.Printf("invalid json in %v...\n", file)
		return
	}

	fmt.Println("parsed " + file)

	var jsonData WhisperTimestampedJson
	json.Unmarshal(byteValue, &jsonData)

	// filling this shit backwards, so when walking through the linked list it will be in order
	for _, segment := range jsonData.Segments {
		for _, word := range segment.Words {
			result.AddTail(
				&LyricsListEntry{
					strings.Split(file, ".")[0], // Song  string
					word.Start - 0.01,           // Start float32
					word.End + 0.01,             // End   float32
					word.Text,                   // Text  string
					nil,                         // Next  *LyricsListEnty
				},
			)

		}
	}
}
