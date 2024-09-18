package main

import "fmt"
import "github.com/spf13/cobra"

func main() {
	var rootCmd = &cobra.Command{
		Use:   "<string>",
		Short: "input a string, as output a audio file of the given string will be output",
		// Long: `A Fast and Flexible Static Site Generator built with
		//                love by spf13 and friends in Go.
		//                Complete documentation is available at http://hugo.spf13.com`,
		Run: func(cmd *cobra.Command, args []string) {
			// TODO handle error
			test(args[0])
		},
	}

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		// os.Exit(1)
	} // read user input string
}

func test(input string) {
	fmt.Println("starting...", input)

	// parse lyrics files
	var lyricsList LyricsList
	lyricsList = Parse()

	// calculate which splices of songs we want
	spliceList := decide("test", lyricsList)

	// splice and concatenate these
	doFFMPEGMagic(spliceList)

	fmt.Println("Done!")
	// output sound file
}
