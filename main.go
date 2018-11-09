package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("csvrepeater must have one argument (filepath to CSV file)")
		os.Exit(1)
	}

	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	r := csv.NewReader(bufio.NewReader(file))
	records, err := r.ReadAll()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	recordsAreEmpty := false
	for recordsAreEmpty != true {
		printed := false
		for i := 0; i < len(records); i++ {

			if len(records[i]) < 2 {
				fmt.Println("Error: CSV row must contain at least two columns")
				os.Exit(1)
			}

			repeats, err := strconv.ParseInt(records[i][1], 10, 64)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			if repeats > 0 {
				fmt.Println(records[i][0])
				repeats--
				records[i][1] = strconv.Itoa(int(repeats))
				printed = true
			}

			recordsAreEmpty = !printed
		}
	}
}
