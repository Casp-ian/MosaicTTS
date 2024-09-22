package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"io/ioutil"
)

func main() {
	var rootCmd = &cobra.Command{
		Use:   "<string>",
		Short: "input a string, as output a audio file of the given string will be output",
		// Long: `looooong`,
		Run: func(cmd *cobra.Command, args []string) {
			// TODO handle unwanted cases like more than one input, and stupid args
			// TODO add args like --lyrics-dir --music-dir --greedy --loose --output
			test(args[0])
		},
	}

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		// os.Exit(1)
	}
}

func test(input string) {
	fmt.Println("starting...")

	// lyricsDir := "/mnt/Data/Projects/cli/MosaicTTS/lyrics/"
	lyricsDir := "../lyrics/"
	lyricsFiles, err := ioutil.ReadDir(lyricsDir)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	var LyricFileNames []string
	for _, lyricsFile := range lyricsFiles {
		LyricFileNames = append(LyricFileNames, lyricsFile.Name())
	}

	// audioDir := "/mnt/Data/Projects/cli/MosaicTTS/audio/"
	audioDir := "../audio/"
	audioFiles, err := ioutil.ReadDir(audioDir)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	var audioFileNames []string
	for _, audioFile := range audioFiles {
		audioFileNames = append(audioFileNames, audioDir, audioFile.Name())
	}

	var outputFile = "output.mp3"

	// parse lyrics files
	var lyricsList LyricsList
	lyricsList = Parse(lyricsDir, LyricFileNames)

	// calculate which splices of songs we want
	spliceList, err := decide(input, lyricsList)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	var spliceListEntry *SpliceListEntry = spliceList.Head
	for spliceListEntry != nil {
		fmt.Println(spliceListEntry)

		spliceListEntry = spliceListEntry.Next
	}
	// splice and concatenate these
	doFFMPEGMagic(outputFile, audioDir, audioFileNames, spliceList)

	fmt.Println("Done!")
	// output sound file
}
