package ytdl

import (
	"fmt"
	"os"
	"os/exec"
)

func Download(path string) {
	fmt.Println("Downloading ...")
}

func GetMp3(link string, tempDir string) { // link string, output string) {
	args := []string{"-x", "--audio-format", "mp3", link}

	cmd := exec.Command("yt-dlp", args...)
	cmd.Dir = tempDir

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()

	if err != nil {
		fmt.Printf("%s", err)
	}

}