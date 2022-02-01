package utils

import (
	"fmt"
	"io"
	"io/fs"
	"log"
	"os"
	"strings"
)

// artistMap this variable need to remove duplicated artist in list.
var artistMap = map[string]bool{}

// listOfArtist this variable need to collect original artist.
var listOfArtist = []string{}

/*
	GrabAllFiles this function is main function to copy music files(*.mp3).
*/
func GrabAllFiles(sourcepath, folderpath string) {
	checkDir(sourcepath)
	allFiles, err := os.ReadDir(sourcepath)
	if err != nil {
		log.Fatal(err)
	}
	artistSlice := []string{}
	for _, file := range allFiles {
		artist := strings.Split(file.Name(), " - ")
		artistSlice = append(artistSlice, artist[0])
	}
	filteredArtists := filterArtist(artistSlice)
	createDirForArtist(folderpath, filteredArtists)
	for _, value := range allFiles {
		fileInfo, err := value.Info()
		if err != nil {
			log.Fatal(err)
		}
		err = copyFile(sourcepath, folderpath, fileInfo)
		if err != nil {
			fmt.Printf("File copying failed: %q\n", err)
		}
	}
}

/*
	checkDir this function checking directory on existing.
*/
func checkDir(folderpath string) {
	if _, err := os.Stat(folderpath); os.IsNotExist(err) {
		log.Fatal(err)
	} else {
		log.Printf("This path exist: %s", folderpath)
	}
}

/*
	filterArtist this function define collabs and duplicates then remove or adding it.
*/
func filterArtist(artistList []string) []string {
	for _, artist := range artistList {
		artist = strings.ToLower(artist)
		singleArtist, state := filteringMusic(artist)
		if state != false {
			checkAndAddArtist(singleArtist)
		}
	}
	return listOfArtist
}

/*
	filteringMusic this function define this letters like ["feat.", "x", ",", etc...].
*/
func filteringMusic(artist string) (string, bool) {
	letters := strings.Split(artist, " ")
	for index, letter := range letters {
		switch letter {
		case ",":
			singleArtist := strings.Join(letters[:index], " ")
			return singleArtist, true
		case "feat.":
			singleArtist := strings.Join(letters[:index], " ")
			return singleArtist, true
		case "ft.":
			singleArtist := strings.Join(letters[:index], " ")
			return singleArtist, true
		case "x":
			singleArtist := strings.Join(letters[:index], " ")
			return singleArtist, true
		case "×":
			singleArtist := strings.Join(letters[:index], " ")
			return singleArtist, true
		case "X":
			singleArtist := strings.Join(letters[:index], " ")
			return singleArtist, true
		case "_":
			singleArtist := strings.Join(letters[:index], " ")
			return singleArtist, true
		default:
			singleArtist := strings.Join(letters[:index+1], " ")
			if singleArtist[len(singleArtist)-1] == ',' {
				singleArtist = singleArtist[:len(singleArtist)-1]
				return singleArtist, true
			} else if singleArtist == artist {
				return singleArtist, true
			}
		}
	}
	return artist, false
}

/*
	checkAndAddArtist this function uses function(checkArtist) and adding artist to list.
*/
func checkAndAddArtist(artist string) {
	if state := checkArtist(artist); state {
		listOfArtist = append(listOfArtist, artist)
	}
}

/*
	checkArtist this function checking artist with special artist map, to define copies.
*/
func checkArtist(artist string) bool {
	if _, value := artistMap[artist]; !value {
		artistMap[artist] = true
		return true
	} else {
		return false
	}
}

/*
	createDirForArtist this function create artist directory in destination folder.
*/
func createDirForArtist(folder string, artistList []string) {
	for _, artist := range artistList {
		if _, err := os.Stat(folder + artist); os.IsNotExist(err) {
			err = os.Mkdir(folder+artist, 0777)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}

/*
	copyFile this function copying files in special generated artist directory.
*/
func copyFile(sourceFolder, dst string, file fs.FileInfo) error {
	log.Printf("Copying file: %s", file.Name())
	singleArtist, state := filteringMusic(strings.Split(file.Name(), " - ")[0])
	singleArtist = strings.ToLower(singleArtist)
	if state != false {
		fmt.Errorf("\nSomething went wrong!!!\n")

	}
	source, err := os.Open(sourceFolder + file.Name())
	dst = dst + singleArtist + "/" + file.Name()
	if err != nil {
		return err
	}
	defer source.Close()

	if _, err := os.Stat(dst); err == nil {
		return fmt.Errorf("File %s already exists.", dst)
	}
	destination, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer destination.Close()

	if err != nil {
		panic(err)
	}
	buf := make([]byte, file.Size())
	for {
		n, err := source.Read(buf)
		if err != nil && err != io.EOF {
			return err
		}
		if n == 0 {
			break
		}
		if _, err := destination.Write(buf[:n]); err != nil {
			return err
		}
	}

	return err
}

/*
	If you fucked up and created directories in another place, with this function you can remove it.
	Heh :D
*/
// func RemoveDirs(folder string, artistList []string) {
// 	for _, artist := range artistList {
// 		err := os.Remove(folder + artist)
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 	}
// }
