package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

var (
	ErrFileNotFound = errors.New("file not found")
	changeDirectory = "/home/jj/git/golang-repo/packt-workshop/ch12-files-sys/e1202"
)

func createBackup(working, backup string) error {
	_, err := os.Stat(working)
	if err != nil {
		if os.IsNotExist(err) {
			return ErrFileNotFound
		}
		return err
	}
	workFile, err := os.Open(working)
	if err != nil {
		return err
	}

	content, err := ioutil.ReadAll(workFile)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(backup, content, 0644)
	if err != nil {
		return err
	}
	return nil
}

func addNotes(workingFile, notes string) error {
	notes += "\n"
	f, err := os.OpenFile(workingFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0770)
	if err != nil {
		return err
	}
	defer func() {
		fmt.Println("Closing file:", workingFile)
		f.Close()
	}()

	if _, err := f.Write([]byte(notes)); err != nil {
		return err
	}
	return nil
}

func listDir(dir string) error {
	files, err := ioutil.ReadDir(dir)
	dirFiles := []string{}
	if err != nil {
		return err
	}
	for _, f := range files {
		//fmt.Println(f.Name())
		dirFiles = append(dirFiles, f.Name())
	}
	fmt.Println(dirFiles)
	return nil
}
func main() {
	workingDir, err := os.Getwd()
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println("Working directory:", workingDir)
	if workingDir != changeDirectory {
		os.Chdir(changeDirectory)
	}
	fmt.Println("Files in directory:")
	err = listDir("./")
	if err != nil {
		fmt.Println("Error:", err)
	}
	err = createBackup("notes.txt", "backup.txt")
	if err != nil {
		log.Fatal("Error:", err)
	}

	err = addNotes("notes.txt", "my new note")
	if err != nil {
		log.Fatal("Error:", err)
	}
	for i := 0; i < 10; i++ {
		err = addNotes("notes.txt", fmt.Sprintf("Message number: %v", i))
		if err != nil {
			log.Fatal("Error:", err)
		}
		fmt.Println("written to file, note:", i)
	}
}
