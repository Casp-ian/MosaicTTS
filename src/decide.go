package main

import (
	"errors"
	"strings"
)

type IncompleteSplice struct {
	Index     int
	StartTime float32
	EndTime   float32
	Song      string
}

func decide(input string, thing LyricsList) (SplicesList, error) {
	var result = SplicesList{nil, nil}

	var wordList = strings.Split(input, " ")

	// collect all lyrics usableSpliceList entries we can use
	var usableSpliceList []IncompleteSplice

	var lyricsListEntry *LyricsListEntry = thing.Head
	for lyricsListEntry != nil {
		// were loading list absolutely full of duplicates, depending on how slices or arrays or whatever this is are implemented, this might be bad for performance TODO
		var acquiredIndexes = getIndexesOf(wordList, lyricsListEntry.Text)
		for _, acquiredIndex := range acquiredIndexes {
			usableSpliceList = append(usableSpliceList, IncompleteSplice{
				acquiredIndex,
				lyricsListEntry.Start,
				lyricsListEntry.End,
				lyricsListEntry.Song,
			})
		}
		lyricsListEntry = lyricsListEntry.Next
	}

	// pick and choose out of collected lyric list entries to define our splices
	for index, word := range wordList {
		listSubset := itemsThatHaveMatchingIndex(usableSpliceList, index)
		if len(listSubset) == 0 {
			return result, errors.New("THE WORD '" + word + "' DOES NOT APPEAR IN LYRICS, FUCKED UP")
		}
		result.AddTail(&SpliceListEntry{
			listSubset[0].Song,      // Song  string
			listSubset[0].StartTime, // Start float32
			listSubset[0].EndTime,   // End   float32
			nil,                     // Next  *SpliceListEntry
		})

		// TODO some fucking how give a preference for words in sequence
		// also conjoin these into 1 splice
	}

	return result, nil
}

func getIndexesOf(list []string, string string) []int {
	var result []int

	for index, element := range list {
		if purify(element) == purify(string) {
			result = append(result, index)
		}
	}

	return result
}

func itemsThatHaveMatchingIndex(list []IncompleteSplice, index int) []IncompleteSplice {
	var result []IncompleteSplice

	for _, element := range list {
		if element.Index == index {
			result = append(result, element)
		}
	}

	return result
}

// unused
func Contains(list []int, int int) bool {
	for _, element := range list {
		if element == int {
			return true
		}
	}
	return false
}
func getIndexOf(list []string, string string) int {
	for index, element := range list {
		if element == string {
			return index
		}
	}
	return -1
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
