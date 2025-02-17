package actions

import (
	"bufio"
	"errors"
	"log"
	"os"
)

func TruncateCSV(file string) {
	filePath := OutputDirPath + file + ".csv"
	fileSeek, err := os.OpenFile(filePath, os.O_RDWR, 0644)
	if err != nil {
		log.Println("Unable to access file.", err)
	}
	defer fileSeek.Close()
	scanner := bufio.NewScanner(fileSeek)

	if scanner.Scan() {
		for {
			header_bytes := scanner.Text()
			truncateErr := fileSeek.Truncate(0)
			_, seekError := fileSeek.Seek(0, 0)
			_, writeErr := fileSeek.WriteString(header_bytes)

			if truncateErr != nil || seekError != nil || writeErr != nil {

				if errors.Is(truncateErr, os.ErrNotExist) || errors.Is(seekError, os.ErrNotExist) || errors.Is(writeErr, os.ErrNotExist) {
					log.Println("File does not exist.")
					break
				}

				continue
			}

			break
		}
	}

	if err := scanner.Err(); err != nil {
		log.Println("Unable to access file.", err)
	}
}
