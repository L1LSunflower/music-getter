package main

import (
	"flag"
	"mapper/utils"
)

func main() {
	//	sourcepath is variable of source path to get music from there
	sourcepath := flag.String("source", "./test", "Source folder to get music")
	//	folderpath is variable of destination path to save music there
	folderpath := flag.String("destination", "./music", "Destination folder to copy there music")
	flag.Parse()
	//	utils.GrabAllFiles(source and destination) copying music files from source to destination
	utils.GrabAllFiles(*sourcepath, *folderpath)
}
