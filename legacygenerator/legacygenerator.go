package main

import (
	"encoding/csv"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {
	fileName := "legacy.csv"

	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println(err)
	}
	writer := csv.NewWriter(file)
	firstRow := []string{"Output", "MaxOutput"}
	writer.Write(firstRow)
	output := 100
	outputMax := 200

	for {
		time.Sleep(time.Millisecond * 100)
		output = output + rand.Intn(3) - 1
		if output > outputMax {
			output = outputMax
		} else if output < 0 {
			output = 0
		}
		row := []string{strconv.Itoa(output), strconv.Itoa(outputMax)}
		writer.Write(row)
		// fmt.Println("written line")
		writer.Flush()
	}
}
