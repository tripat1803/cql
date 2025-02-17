package actions

import (
	"errors"
	"log"
	"os"
)

func DeleteCSV(file string) {
	filePath := OutputDirPath + file + ".csv"

	DeleteConfigKey(file)

	for {
		err := os.Remove(filePath)

		if err != nil {
			if errors.Is(err, os.ErrNotExist) {
				log.Println("File does not exist.")
				break
			}

			continue
		}

		break
	}

}
