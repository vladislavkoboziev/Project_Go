package handle

import (
	"awesomeProject/src/model"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func GetPeople(w http.ResponseWriter, r *http.Request) {
	var people = model.Person{}
	users := people.GetAll()
	json.NewEncoder(w).Encode(users)
}
func InsertPerson(w http.ResponseWriter, r *http.Request) {
	var user = model.Person{}
	user.FirstName = r.FormValue("firstname")
	user.LastName = r.FormValue("lastname")
	user.Email = r.FormValue("email")
	user.Gender = r.FormValue("gender")
	user.Loan, _ = strconv.ParseFloat(r.FormValue("loan"), 64)
	user.Insert()
	json.NewEncoder(w).Encode(user)
}
func UpdatePerson(w http.ResponseWriter, r *http.Request) {
	var user = model.Person{}
	user.Id = r.FormValue("id")
	user.FirstName = r.FormValue("firstname")
	user.Update()
	json.NewEncoder(w).Encode(user)

}
func DeletePerson(w http.ResponseWriter, r *http.Request) {
	var user = model.Person{}
	id, _ := strconv.Atoi(r.FormValue("id"))
	user.Delete(id)
	response := "Ok!"
	json.NewEncoder(w).Encode(response)
}
func UsersFilter(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	firstname := r.FormValue("firstDate")
	secondname := r.FormValue("secondDate")
	gender := r.FormValue("gender")
	fmt.Println("fields for Filter:", name, firstname, secondname, gender)
	users := model.Filters(name, firstname, secondname, gender)
	json.NewEncoder(w).Encode(users)

}
