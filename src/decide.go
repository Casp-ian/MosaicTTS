package main

import (
	"fmt"
	"strings"
)

func decide(input string, thing LyricsList) SplicesList {
	var result = SplicesList{nil, nil}

	var wordList = strings.Split(input, " ")

	// TODO draw this algorithm out on a whiteboard with samme :)
	// goal of this part is to get every possible slice that could be used to make up the output
	var test *LyricsListEntry = thing.Head
	for test != nil {
		var temp = getIndexesOf(wordList, test.Text)
		if len(temp) == 0 {
			// TODO
		} else {
			fmt.Println(test.Start, test.End, test.Text)
		}
		test = test.Next
	}

	return result
}

func getIndexesOf(array []string, string string) []int {
	var result []int

	for index, element := range array {
		if purify(element) == purify(string) {
			result = append(result, index)
		}
	}

	return result
}

// dramatic ah function name
func purify(impure string) string {
	// TODO at some later time we could leave these in and match them if possible to match the tone and pacing of the sentance better
	impure = strings.Replace(impure, "?", "", -1)
	impure = strings.Replace(impure, "!", "", -1)
	impure = strings.Replace(impure, ",", "", -1)
	impure = strings.Replace(impure, ".", "", -1)
	impure = strings.Replace(impure, " ", "", -1)
	impure = strings.ToLower(impure)
	return impure
}
