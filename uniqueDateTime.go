package main

import (
	"bufio"
	"fmt"
	"os"
)

// Data formatting logic to ensure date is ISO 8601 valid before adding to list
func isDateTimeValid(currDateTime string) bool {

	// How the string can be separated into sub-sections
	// [Y,Y,Y,Y,-,M,M,-,D,D,T, h, h,:, m, m,:, s, s,TZD]
	// [0,1,2,3,-,5,6,-,8,9,T,11,12,:,14,15,:,17,18,19]

	// Ensure ISO 8601 date-time formatting symbols are present (e.g. -,T,:)
	if currDateTime[4:5] != "-" || currDateTime[7:8] != "-" || currDateTime[10:11] != "T" || currDateTime[13:14] != ":" || currDateTime[16:17] != ":" {
		return false
	}

	// currentDateTime sub-sections
	currYear := currDateTime[:4]
	currMonth := currDateTime[5:7]
	currDay := currDateTime[8:10]
	currHour := currDateTime[11:13]
	currMin := currDateTime[14:16]
	currSec := currDateTime[17:19]
	currTZD := currDateTime[19:]

	// Ensure date and time within acceptable ranges
	if "0000" <= currYear && currYear <= "9999" && "01" <= currMonth && currMonth <= "12" && "01" <= currDay && currDay <= "31" && "00" <= currHour && currHour <= "23" && "00" <= currMin && currMin <= "59" && "00" <= currSec && currSec <= "59" {

		// Ensure proper Time-Zone Designation exists

		// if TZD == "Z" add date to slice
		if currTZD == "Z" {

			return true

			// TZD has to have a "+" or "-" and have the hh:mm format to be accepted
		} else if len(currTZD) == 6 && (currTZD[0:1] == "+" || currTZD[0:1] == "-") && currTZD[3:4] == ":" {

			tzdHour := currTZD[1:3]
			tzdMin := currTZD[4:6]

			// TZD hours range from 1 to 12, TZD mins can either be 00 or 30
			if "00" <= tzdHour && tzdHour <= "12" && (tzdMin == "00" || tzdMin == "30") {
				return true
			}
		}
	}
	return false
}

// Checks if current Valid Date is unique
func existingValidDates(dateSlice []string, dateValue string) bool {
	for _, val := range dateSlice {
		if val == dateValue {
			return true
		}
	}
	return false
}

func main() {

	// test data slice (for output)
	validUniqueTestDates := []string{}

	// counters for CLI output information
	totalDatesCounter := 0
	invalidDatesCounter := 0
	validDatesCounter := 0
	duplicateDatesCounter := 0
	validUniqueDatesCounter := 0

	// Open the raw data file
	rawDataFile, err := os.Open("DateTimeTestData.txt")
	if err != nil {
		fmt.Println("error opening data file: ", err)
	}
	defer rawDataFile.Close()

	// Create a file scanner
	dataScanner := bufio.NewScanner(rawDataFile)

	// Scan data and append to validTestDates
	for dataScanner.Scan() {
		totalDatesCounter += 1
		currentDateTime := dataScanner.Text()

		// check if current date-time value is formatted correctly
		if isDateTimeValid(currentDateTime) == true {
			validDatesCounter += 1
			// Check if the valid date-time value is NOT in valid dates slice
			if existingValidDates(validUniqueTestDates, currentDateTime) == false {
				validUniqueDatesCounter += 1

				validUniqueTestDates = append(validUniqueTestDates, currentDateTime)

			} else {
				duplicateDatesCounter += 1
			}

		} else {
			invalidDatesCounter += 1
		}

	}
	if err := dataScanner.Err(); err != nil {
		fmt.Println("error scanning data file: ", err)
	}

	// show different dates types; total, invalid, valid, or valid AND unique in command terminal
	fmt.Printf("Out of the %v total dates,\n%v are invalid\n%v are valid\n%v are valid BUT duplicate\n%v are valid AND unique.\n", totalDatesCounter, invalidDatesCounter, validDatesCounter, duplicateDatesCounter, validUniqueDatesCounter)

	// Create new output data file
	validDataFile, err := os.Create("Output.txt")
	if err != nil {
		fmt.Println("error creating new data file: ", err)
	}
	defer validDataFile.Close()

	writer := bufio.NewWriter(validDataFile)

	// Iterate through valid date-time data in Slice and write to new output file
	for _, validDate := range validUniqueTestDates {
		_, err := writer.WriteString(validDate + "\n")
		if err != nil {
			fmt.Println("Error writing to buffer", err)
			return
		}

		err = writer.Flush()
		if err != nil {
			fmt.Println("Error flushing buffer:", err)
			return
		}
	}

}
