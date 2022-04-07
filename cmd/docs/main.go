package main

import (
	"github.com/AnteWall/go-finansinspektionen/cmd"
	"github.com/spf13/cobra/doc"
	"log"
	"os"
)

func main() {
	f, err := os.OpenFile("README.md", os.O_WRONLY, os.ModeAppend)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	err = doc.GenMarkdown(cmd.RootCmd, f)
	if err != nil {
		log.Fatal(err)
	}
	err = doc.GenMarkdown(cmd.DownloadCmd, f)
	if err != nil {
		log.Fatal(err)
	}
}
