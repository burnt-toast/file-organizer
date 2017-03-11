//Organize files in a given directory into a new location with the files organized in folders
//by creation date
package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

func main() {

	directoryToOrgainze, err := getAndValidateDirectoryToOrganzie()
	if err != nil {
		panic(err)
	}
	fmt.Println("Directory: ", directoryToOrgainze)
	//Loop over directory and create an array of files and non files
	itemsInDirectory, _ := ioutil.ReadDir(directoryToOrgainze)

	mapByDate := make(map[string][]string)
	fmt.Println("Overall length: ", len(itemsInDirectory))
	var numberAdded int
	var numberCreated int
	for _, item := range itemsInDirectory {
		if item.IsDir() {
			errText := "Directories found in path provided " + item.Name()
			panic(errors.New(errText))
		}
		pathToSpecificFile := directoryToOrgainze + "\\" + item.Name()
		year, month, day := item.ModTime().Date()
		yyyyMMDD := strconv.Itoa(year) + "-" + month.String() + "-" + strconv.Itoa(day)
		existingList, exists := mapByDate[yyyyMMDD]
		//fmt.Println(item.ModTime().Date())
		if exists {
			//add to existing slice
			numberAdded++
			existingList = append(existingList, pathToSpecificFile)
		} else {
			//create new slice and add
			numberCreated++
			mapByDate[yyyyMMDD] = []string{pathToSpecificFile}
		}

	}
	fmt.Println("Added: ", numberAdded)
	fmt.Println("Created: ", numberCreated)
	//fmt.Println(mapByDate)
}

//Validate that a path was provided via command line and that the
//directory exists
func getAndValidateDirectoryToOrganzie() (string, error) {
	if len(os.Args) < 2 {
		return "", errors.New("Missing directory command line arg")
	}

	directoryToOrganize := os.Args[1]
	_, err := os.Stat(directoryToOrganize)
	if err == nil {
		return directoryToOrganize, nil
	}
	if os.IsNotExist(err) {
		return "", err
	}
	return directoryToOrganize, nil
}
