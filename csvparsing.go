package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"sort"
	"strconv"
	"time"
)

var people []Person ///slice with our data
type Gender int

const (
	Male Gender = iota
	Female
	gender
)

type Person struct {
	Id               int
	FirstName        string
	LastName         string
	Email            string
	Gender           string
	DateRegistration time.Time
	Loan             float64
}

func (g Gender) String() string {
	return [...]string{"Male", "Female", "gender"}[g]
}
func parsing_csv() {
	csvFille, _ := os.Open("MOCK_DATA.csv")
	reader := csv.NewReader(csvFille)
	for i := 0; i < 101; i++ {
		line, _ := reader.Read()
		id, _ := strconv.Atoi(line[0])
		loan, _ := strconv.ParseFloat(line[6], 64)
		dataPeople, _ := time.Parse("1/2/2006", line[5])
		people = append(people, Person{
			Id:               id,
			FirstName:        line[1],
			LastName:         line[2],
			Email:            line[3],
			Gender:           line[4],
			DateRegistration: dataPeople,
			Loan:             loan,
		})
	}

}
func main() {

}

func loanrange() {
	var (
		first_range float64 = 253372.042169
		last_range  float64 = 571786.599688
	)
	for i := 0; i < len(people); i++ {
		fmt.Println(people[i].Loan)
	}
	fmt.Println("До проверки по диапозону")
	for i := 0; i < len(people); i++ {
		if people[i].Loan >= first_range && people[i].Loan <= last_range {
			fmt.Println(people[i])
		}
	}
}
func prozent_Man_Woman() {
	var m int
	var f int
	for _, value := range people {
		if value.Gender == "" {
		} else if value.Gender == "Male" {
			m++
		} else if value.Gender == "Female" {
			f++
		}

	}
	fmt.Println("Женщины", (f*len(people)-1)/100, "%")
	fmt.Println("Мужчины", (m*len(people)-1)/100, "%")
}
func genderPercent() {
	countMale := 0
	countFemale := 0
	var g Gender
	for i := 1; i <= len(people)-1; i++ {

		if people[i].Gender == "Male" {
			g = Male
		} else if people[i].Gender == "Female" {
			g = Female
		} else if people[i].Gender == "gender" {
			g = 0
		}
		switch g {
		case Male:
			countMale++
		case Female:
			countFemale++
		case gender:
		}
	}
	fmt.Println("Number of men", countMale)
	fmt.Println("Number of women", countFemale)
	percentMale := (countMale * (len(people) - 1))
	percentFemale := (countFemale * (len(people) - 1))
	fmt.Println("Percent of men", percentMale)
	fmt.Println("Percent of women", percentFemale)
}
func dataPeople(dateFirst, dateLast time.Time, people []Person) { //////// date!!!!!!!!

	for _, value := range people {

		if value.DateRegistration.Unix() >= dateFirst.Unix() && value.DateRegistration.Unix() <= dateLast.Unix() {

			fmt.Println(value)

		}

	}

}
func sortLoan() { /// sort the data by Loan
	sort.Slice(people, func(i, j int) bool {
		return people[i].Loan < people[j].Loan
	})
	for _, value := range people {
		fmt.Println(value)
	}
}
