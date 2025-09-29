package main

import (
	"fmt"
	"getmp3/ytdl"
	"os"

	// "github.com/therecipe/qt/core"
	// "github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/widgets"
)


func createTempDir() string {
	tempDir, err := os.MkdirTemp("", "getmp3")
	if err != nil {
		fmt.Println("Failed to create temp dir.")
		os.Exit(1)
	}

	fmt.Printf("Download to <%s>\n", tempDir)
	return tempDir
}

func main() {
	// link := "https://www.youtube.com/watch?v=mw7Djj4m9Hk"
	link := "https://www.youtube.com/watch?v=2oL7h7MY7AA"
	ytdl.Download(link)
	tempDir := createTempDir()
	defer os.RemoveAll(tempDir)
	// getmp3(link, tempDir)

	widgets.NewQApplication(len(os.Args), os.Args)

	var window = widgets.NewQMainWindow(nil, 0)
	window.SetWindowTitle("Get mp3")

	window.Show()

	widgets.QApplication_Exec()
}
