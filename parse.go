package main

import (
	"encoding/csv"
	"fmt"
	_ "gorm.io/driver/sqlite"
	"gorm.io/gorm"
	_ "gorm.io/gorm"
	"os"
)

type Products struct {
	gorm.Model
	Title string
	Asin string
}

type Reviews struct {
	gorm.Model
	Asin string
	Revie`w string
}

func main() {

	file, err := os.Open("csv/Products.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.FieldsPerRecord = 2
	reader.Comment = '#'

	for {
		record, e := reader.Read()
		if e != nil {
			fmt.Println(e)
			break
		}
		fmt.Println(record)
	}
}
