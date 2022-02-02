package main

import (
	"flag"
	"mapper/utils"
)

func main() {
	//	sourcepath is variable of source path to get music from there
	sourcepath := flag.String("source", "/home/deka/Music/000_UNCLS_MUSIC/Music/", "Source folder to get music")
	//	folderpath is variable of destination path to save music there
	folderpath := flag.String("destination", "/media/deka/8498967F98967002/003_Media/001_Music/", "Destination folder to copy there music")
	flag.Parse()
	//	utils.GrabAllFiles(source and destination) copying music files from source to destination
	utils.GrabAllFiles(*sourcepath, *folderpath)
}
