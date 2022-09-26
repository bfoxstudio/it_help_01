package main

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func readCsvFile(filePath string) [][]string {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Unable to read input file "+filePath, err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()

	if err != nil {
		log.Fatal("Unable to parse file as CSV for "+filePath, err)
	}

	return records
}

func main() {
	originalThemplate, err := ioutil.ReadFile("form_confirmation_utf8.csv")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	records := readCsvFile("table.csv")
	for i := range records {
		input := originalThemplate
		input = bytes.Replace(input, []byte("{{0}}"), []byte(records[i][0]), -1)
		input = bytes.Replace(input, []byte("{{1}}"), []byte(records[i][1]), -1)
		input = bytes.Replace(input, []byte("{{2}}"), []byte(records[i][2]), -1)
		input = bytes.Replace(input, []byte("{{3}}"), []byte(records[i][3]), -1)
		input = bytes.Replace(input, []byte("{{4}}"), []byte(records[i][4]), -1)
		input = bytes.Replace(input, []byte("{{5}}"), []byte(records[i][5]), -1)
		input = bytes.Replace(input, []byte("{{6}}"), []byte(records[i][6]), -1)
		input = bytes.Replace(input, []byte("{{7}}"), []byte(records[i][7]), -1)
		input = bytes.Replace(input, []byte("{{8}}"), []byte(records[i][8]), -1)
		input = bytes.Replace(input, []byte("{{9}}"), []byte(records[i][8]), -1)
		if err = ioutil.WriteFile(fmt.Sprintf("modified_%v.csv", i), input, 0666); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
}
