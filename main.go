package main

import (
	"fmt"

	"github.com/360EntSecGroup-Skylar/excelize/v2"
)

func main() {
	f, err := excelize.OpenFile("test.xlsx")
	if err != nil {
		panic(err)
	}

	rows, err := f.Rows(f.GetSheetName(1))
	if err != nil {
		panic(err)
	}

	var rowCount int
	for rows.Next() {
		rowCount++
		fmt.Println(rowCount)
	}
}
