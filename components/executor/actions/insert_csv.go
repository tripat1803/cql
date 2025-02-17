package actions

import (
	"log"
	"os"
	"strings"
)

func InsertCSV(
	file string,
	values [][]string,
	columns []string,
) {
	filepath := OutputDirPath + file + ".csv"
	file_headings := FetchColumnsHeadings(filepath)

	var file_columns []string = columns
	if len(columns) == 0 {
		file_columns = file_headings
	}

	for _, row := range values {
		if len(row) != len(file_columns) {
			log.Fatalln("Invalid row length", row)
		}
	}

	csv_values := make([]string, len(values))
	start_idx := 0
	for heading_idx, heading := range file_headings {
		ending_symbol := ","
		if heading_idx == len(file_headings)-1 {
			ending_symbol = ""
		}
		if start_idx < len(file_columns) && file_columns[start_idx] == heading {
			for i := 0; i < len(values); i++ {
				csv_values[i] += values[i][start_idx] + ending_symbol
			}
			start_idx++
		} else {
			for i := 0; i < len(values); i++ {
				csv_values[i] += ending_symbol
			}
		}
	}

	filePtr, err := os.OpenFile(filepath, os.O_APPEND, 0644)
	if err != nil {
		log.Fatalln("Unable to access file")
	}
	defer filePtr.Close()

	_, writeErr := filePtr.WriteString("\n" + strings.Join(csv_values, "\n"))
	if writeErr != nil {
		log.Fatalln("Unable to write to file")
	}
}
