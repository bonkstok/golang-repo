package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

type budgetCategory string
type transaction struct {
	ID       int
	payee    string
	spent    float64
	category string
}

const (
	autoFuel   budgetCategory = "fuel"
	food       budgetCategory = "food"
	mortgage   budgetCategory = "mortgage"
	repairs    budgetCategory = "repairs"
	insurance  budgetCategory = "insurance"
	utilities  budgetCategory = "utilities"
	retirement budgetCategory = "retirement"
)

func main() {
	inputFile := flag.String("file", "", "file for input")
	logFile := flag.String("logfile", "", "log file to write to")
	flag.Parse()
	if *inputFile == "" || *logFile == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}
	fmt.Println("Input:", *inputFile)
	fmt.Println("Logfile:", *logFile)

	//setting up logging
	//checkFileIsValid(*logFile)
	logOut, err := os.OpenFile(*logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0770)
	if err != nil {
		log.Fatal("Cannot create logfile")
	}
	defer func() {
		log.Printf("Closing log file %s", logOut.Name())
		logOut.Close()
	}()

	logger := log.New(logOut, "", log.LstdFlags)
	//logger.Println("Adding a line")
	logger.Println("Checking if input file exists:", *inputFile)
	if !checkFileExists(*inputFile) {
		logger.Printf("File: '%s' does not exist. Exiting", *inputFile)
		os.Exit(1)
		//log.Fatal("File does not exist, exiting...")
	}
	logger.Printf("Input file '%s' exists, using it as input", *inputFile)
	//fmt.Printf("%T\n", logger)
	transactions := readCSVFile(*inputFile, true, logger)
	for _, v := range transactions {
		fmt.Println(v)
	}

}

func checkFileExists(file string) bool {
	_, err := os.Stat(file)
	if os.IsNotExist(err) {
		log.Println("File does not exist")
		return false
	}
	//fmt.Println(err)
	return true
}

func readCSVFile(file string, headers bool, logger *log.Logger) []transaction {
	//we already checked if file exist on startup
	toReturn := []transaction{}

	csvFile, err := os.Open(file)
	if err != nil {
		logger.Fatal("Error reading from csv file.")
	}
	defer csvFile.Close()
	reader := csv.NewReader(bufio.NewReader(csvFile))
	lineCounter := 0

	for {
		var row transaction

		line, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			logger.Println("Failed to read line:", line)
			lineCounter++
			continue
		}
		if headers && lineCounter == 0 {
			lineCounter++
			continue
		}

		for i, v := range line {

			switch i {
			case 0:
				v = strings.TrimSpace(v)
				row.ID, err = strconv.Atoi(v)
				if err != nil {
					logger.Println("Skipping line with values:", line)
					break
				}
			case 1:
				v = strings.TrimSpace(v)
				row.payee = v
			case 2:
				v = strings.TrimSpace(v)
				row.spent, err = strconv.ParseFloat(v, 64)
				if err != nil {
					logger.Println("Skipping line with values:", line)
					break
				}
			case 3:
				v = strings.TrimSpace(v)
				row.category = v
			}
		}

		lineCounter++
		toReturn = append(toReturn, row)

	}
	return toReturn
}
