package actions

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func ValidateIsNewFile(filePath string) {
	if _, err := os.Stat(filePath); err == nil {
		log.Fatalln("File already exists.")
	}
}

func ValidateIsFilePresent(filePath string) {
	if _, err := os.Stat(filePath); err != nil {
		log.Fatalln("File not found!")
	}
}

func ValidateOutputDir() {
	if _, err := os.Stat(OutputDirPath); os.IsNotExist(err) {
		err := os.Mkdir(OutputDirPath, os.ModePerm)
		if err != nil {
			log.Fatalln("Some error occured while configuiring output directory.")
		}
	}
}

func FetchColumnsHeadings(filePath string) []string {
	ValidateIsFilePresent(filePath)
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalln("Cannot read file")
	}
	defer file.Close()

	headings := ""
	scanner := bufio.NewScanner(file)
	if scanner.Scan() {
		headings = scanner.Text()
	}

	return strings.Split(headings, ",")
}
