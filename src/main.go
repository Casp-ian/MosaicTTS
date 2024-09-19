package main

import "fmt"
import "github.com/spf13/cobra"

func main() {
	var rootCmd = &cobra.Command{
		Use:   "<string>",
		Short: "input a string, as output a audio file of the given string will be output",
		// Long: `looooong`,
		Run: func(cmd *cobra.Command, args []string) {
			// TODO handle unwanted cases like more than one input, and stupid args
			// TODO add args like --lyrics-dir --music-dir --greedy --loose
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

	var output = "output.mp3"

	// parse lyrics files
	var lyricsList LyricsList
	lyricsList = Parse()

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
	doFFMPEGMagic(output, spliceList)

	fmt.Println("Done!")
	// output sound file
}
