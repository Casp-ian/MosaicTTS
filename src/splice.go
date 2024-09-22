package main

import (
	"fmt"
	"os/exec"
	"strconv"
	"strings"
)

func doFFMPEGMagic(output string, audioDir string, files []string, splices SplicesList) {
	cmd := exec.Command("ffmpeg", "-version")
	if err := cmd.Run(); err != nil {
		fmt.Println("ffmpeg might be missing")
	}

	var spliceCount = 0

	var currentSplice = splices.Head
	for currentSplice != nil {
		createSplice(currentSplice.Start, currentSplice.End, audioDir+currentSplice.Song+".mp3", fmt.Sprint(spliceCount)+".mp3")

		currentSplice = currentSplice.Next
		spliceCount = spliceCount + 1
	}

	concat(spliceCount, audioDir, output)
	cleanup(spliceCount)
}

func createSplice(start float32, end float32, input string, output string) {
	cmd := exec.Command("ffmpeg", "-y", "-ss", strconv.FormatFloat(float64(start), 'g', -1, 32), "-to", strconv.FormatFloat(float64(end), 'g', -1, 32), "-i", input, output)
	if err := cmd.Run(); err != nil {
		fmt.Println(err)
	}
}

func concat(spliceCount int, audioDir string, output string) {

	var files []string

	for i := 0; i < spliceCount; i++ {
		files = append(files, fmt.Sprintf("%v.mp3", i))
	}

	// TODO use the demuxer method instead of the protocol method, this gets rid of the 'Header missing' errors https://trac.ffmpeg.org/wiki/Concatenate#demuxer
	cmd := exec.Command("ffmpeg", "-i", "concat:"+strings.Join(files, "|"), "-c", "copy", output)
	if err := cmd.Run(); err != nil {
		fmt.Println(err)
	}
}

func cleanup(spliceCount int) {

	for i := 0; i < spliceCount; i++ {
		file := fmt.Sprintf("%v.mp3", i)
		cmd := exec.Command("rm", file)
		if err := cmd.Run(); err != nil {
			fmt.Println(err)
		}
	}

}
