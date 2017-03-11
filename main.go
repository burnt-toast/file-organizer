/**
Organize files in a given directory into a new location with the files organized in folders
by creation date
*/
package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {

	directoryToOrgainze, err := getAndValidateDirectoryToOrganzie()
	if err != nil {
		panic(err)
	}
	fmt.Println("Directory: ", directoryToOrgainze)
	//Loop over directory and create an array of files and non files
	itemsInDirectory, _ := ioutil.ReadDir(directoryToOrgainze)
	for _, item := range itemsInDirectory {
		if item.IsDir() {
			errText := "Directories found in path provided " + item.Name()
			panic(errors.New(errText))
		}
		fmt.Println(item.Name())
	}
}

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
