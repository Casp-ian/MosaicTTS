package main

import "fmt"

func decide(input string, thing LyricsList) SplicesList {
	var result = SplicesList{nil, nil}

	var test *LyricsListEntry = thing.Head
	for test != nil {
		fmt.Println(test.Text)
		test = test.Next
	}

	return result
}
