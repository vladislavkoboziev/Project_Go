package model

import (
	"database/sql"
	"encoding/csv"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"os"
	"strconv"
	"time"
)

type sql_manipulation interface {
	GetAll() []Person
	Insert()
	Update()
	Delete(id int)
}

var people []Person

type Person struct {
	Id               string    `json:"id"`
	FirstName        string    `json:"firstname"`
	LastName         string    `json:"lastname"`
	Email            string    `json:"email"`
	Gender           string    `json:"gender"`
	DateRegistration time.Time `json:"dateregistration"`
	Loan             float64   `json:"loan"`
}

func parsing_csv() {
	csvFille, _ := os.Open("MOCK_DATA.csv")
	reader := csv.NewReader(csvFille)
	for i := 0; i < 101; i++ {
		line, _ := reader.Read()
		loan, _ := strconv.ParseFloat(line[6], 64)
		dataPeople, _ := time.Parse("1/2/2006", line[5])
		people = append(people, Person{
			Id:               line[0],
			FirstName:        line[1],
			LastName:         line[2],
			Email:            line[3],
			Gender:           line[4],
			DateRegistration: dataPeople,
			Loan:             loan,
		})
	}

}
func connections() *sql.DB {
	//connectbd, err := sql.Open("postgres", "postgres://postgres:7154016@localhost/?sslmode=disable")
	connectbd, err := sql.Open("postgres", "postgres://postgres:7154016@localhost/?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	err = connectbd.Ping()
	if err != nil {
		log.Fatal(err)
	}
	return connectbd
}
func (person Person) Insert() {
	connect := connections()
	defer connect.Close()
	rw, err := connect.Exec("insert into table_name (first_name,last_name,email,gender,date_registration,loan) values ($1,$2,$3,$4,$5,$6)",
		person.FirstName, person.LastName, person.Email, person.Gender, time.Now(), person.Loan)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(rw)
} ////2

func (person Person) Update() {
	connect := connections()
	defer connect.Close()
	rs, err := connect.Exec("update table_name set first_name = $1  where id = $2", person.FirstName, person.Id)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(rs.RowsAffected())
} ////3
func (person Person) GetAll() []Person {
	connect := connections()
	defer connect.Close()
	rows, err := connect.Query("select * from table_name")
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()
	for rows.Next() {
		pp := Person{}
		err := rows.Scan(&pp.Id, &pp.FirstName, &pp.LastName, &pp.Email, &pp.Gender, &pp.DateRegistration, &pp.Loan)
		if err != nil {
			fmt.Println(err)
			continue
		}
		people = append(people, pp)
	}
	for _, p := range people {
		fmt.Println(p.Id, p.FirstName, p.LastName, p.Email, p.Gender, p.DateRegistration, p.Loan)
	}
	return people
} ////4
func (person Person) Delete(id int) {
	connect := connections()
	defer connect.Close()
	rs, err := connect.Exec("delete from table_name where id = $1", id)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(rs.RowsAffected())
}

func Filters(name, firstDate, secondDate, gender string) []Person {
	str := "select * from table_name"
	switch {
	case name != "":
		str += " where first_name= '" + name + "'"
		fallthrough

	case gender != "":

		if name != "" && gender != "" {
			str += " and gender= '" + gender + "'"
		} else if name == "" && gender != "" {
			str += " where gender= '" + gender + "'"
		}
	}
	var people []Person
	db := connections()
	defer db.Close()
	rows, err := db.Query(str)
	if err != nil {
		fmt.Println(err)
	}

	defer rows.Close()

	for rows.Next() {
		d := Person{}
		err := rows.Scan(&d.Id, &d.FirstName, &d.LastName, &d.Email, &d.Gender, &d.DateRegistration, &d.Loan)
		if err != nil {
			fmt.Println(err)
			continue
		}
		people = append(people, d)
	}
	return people
}
