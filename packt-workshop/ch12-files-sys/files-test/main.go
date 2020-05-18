package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func checkFileExists(file string) bool {
	f, err := os.Stat(file)
	if os.IsNotExist(err) {
		return false
	}
	f.IsDir()
	//fmt.Println(err)
	return true

}

func getFileInfo(file string) error {
	f, err := os.Stat(file)
	if err != nil {
		return err
	}
	fmt.Println("Name:", f.Name())
	fmt.Println("Directory:", f.IsDir())
	fmt.Println("Mode:", f.Mode())
	fmt.Println("Size:", f.Size())
	return nil
}

func main() {

	//check if file exists
	myFakeFile := "knock.txt"
	fmt.Printf("File %s exists: %v\n", myFakeFile, checkFileExists(myFakeFile))
	myRootFile := "root.txt"
	fmt.Printf("File %s exists: %v\n", myRootFile, checkFileExists(myRootFile))

	getFileInfo(myRootFile)

	//create a new file, if exists will truncate else will create the file
	f, err := os.Create("test.txt")
	if err != nil {
		panic(err)
	}
	//can also just defer f.Close() -> just to be a bit more verbose
	defer func() {
		fmt.Println("Closing file:", f.Name())
		f.Close()
	}()
	

	//write to the file
	f.Write([]byte("Writing byte(s) to file\n"))
	f.WriteString("Writing a string, cool\n")

	//can also create a file and write to it within a single command
	err = ioutil.WriteFile("./iotest.txt", []byte("Some lines sniffff\n"), 0770)
	if err != nil {
		fmt.Println("Error:", err)
	}

	//the commands above all truncate the file, you can also append to a file

}
