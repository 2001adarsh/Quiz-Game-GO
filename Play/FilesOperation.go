package Play

import (
	"encoding/csv"
	"log"
	"os"
)

func readFromCSV(csvFileName *string) [][]string {
	file, err := os.Open(*csvFileName)
	if err != nil {
		log.Fatal("Error in opening of File %s.\n", *csvFileName, err)
		os.Exit(1)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatal("Error: File close:", err)
		}
	}(file)

	records, err := csv.NewReader(file).ReadAll()
	if err != nil {
		log.Fatal("Error in reading CSV.", err)
		return [][]string{}
	}
	return records
}
