package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"

	"github.com/bytedance/gopkg/util/logger"
)

// რა არის საჭირო ეხლა მოდი მოვიფიქროთ
// პირველი წამკითხველა, რომელიც საშუალებას მომცემს წამოვიღო დატა რაიმე .csv file-დან
// Data model და ფუნქცია რომელიც გადამათარგმნინებს, კერძოდ CSV Row-ს Json ობიექტად.
// კონტროლერი დავამტოთ რიც გადასცემს პასუხად? მოდი ეგრე ვქნათ.
// გადავწყვიტე CLI-ზე იყოსთქო.
func main() {
	for _, record := range jsonMapper() {
		record.print(os.Stdout)
	}
}

type DataResponse struct {
	Date     string `json:"date"`
	Name     string `json:"name"`
	Lastname string `json:"lastname"`
	Job      string `json:"job"`
}

func (d *DataResponse) print(w io.Writer) {
	fmt.Println("============================================")
	_, _ = fmt.Fprintf(w, "Name:     %s\n", d.Name)
	_, _ = fmt.Fprintf(w, "Lastname: %s\n", d.Lastname)
	_, _ = fmt.Fprintf(w, "Date:     %s\n", d.Date)
	_, _ = fmt.Fprintf(w, "Job:      %s\n", d.Job)
}

func jsonMapper() []DataResponse {
	fileReader, err := os.Open("data.csv")

	csvReader := csv.NewReader(fileReader)

	all, err := csvReader.ReadAll()
	if err != nil {
		logger.Error("Could not retrieve data from CSV file: ", err.Error())
		return nil
	}

	var response []DataResponse

	for _, record := range all {
		jsonRecord := DataResponse{
			Date:     record[0],
			Name:     record[1],
			Lastname: record[2],
			Job:      record[3],
		}

		response = append(response, jsonRecord)
	}

	return response
}
