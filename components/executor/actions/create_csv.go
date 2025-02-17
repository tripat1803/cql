package actions

import (
	"log"
	"os"
	"strings"
)

func CreateCSV(columns []string, file string) {
	ValidateOutputDir()

	fileName := OutputDirPath + file + ".csv"
	ValidateIsNewFile(fileName)
	log.Println("Creating new csv file...")

	new_file, _ := os.Create(fileName)
	headings := strings.Join(columns, ",")
	new_file.Write([]byte(headings))

	AddConfigKey(file)
}
