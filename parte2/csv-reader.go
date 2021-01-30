package main

import (
	"encoding/csv"
	"io"
	"log"
	"os"
)

//ReadData read a file and return a matrix for use to train o target data
func ReadData(fileName string) ([]string, [][]string) {
	t, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err.Error())
	}
	data := make([][]string, 0)
	r := csv.NewReader((t))
	r.Comment = '#'
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)

		}
		data = append(data, record)
	}
	return data[0], data[1:]
}
