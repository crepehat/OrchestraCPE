package inputs

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

func CsvGetValue(filePath string, column int) (int, error) {
	csvFile, _ := os.Open(filePath)
	reader := csv.NewReader(bufio.NewReader(csvFile))
	var lastValue interface{}
	for {
		line, error := reader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}
		lastValue = line[column]
	}
	lastValueInt, ok := lastValue.(int)
	if !ok {
		return -1, fmt.Errorf("Value was not an integer")
	}
	return lastValueInt, nil
}
