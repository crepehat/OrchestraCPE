package interact

import (
	"bufio"
	"encoding/csv"
	"io"
	"log"
	"os"
)

func CsvGetValue(filePath string, column int) (string, error) {
	csvFile, _ := os.Open(filePath)
	reader := csv.NewReader(bufio.NewReader(csvFile))
	var lastValue string
	for {
		line, error := reader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
			break
		}
		lastValue = line[column]
	}
	// lastValueInt, ok := lastValue.(int)
	// if !ok {
	// 	return -1, fmt.Errorf("Value was not an integer")
	// }
	return lastValue, nil
}
