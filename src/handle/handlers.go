package handle

import (
	"awesomeProject/src/logers"
	"awesomeProject/src/model"
	"encoding/json"
	"net/http"
	"strconv"
)

func GetPeople(w http.ResponseWriter, r *http.Request) {
	logers.Warning.Println("call GetUser function")
	var people = model.Person{}
	users := people.GetAll()
	json.NewEncoder(w).Encode(users)
}
func InsertPerson(w http.ResponseWriter, r *http.Request) {
	logers.Info.Println("call CreateUser function")
	var user = model.Person{}
	user.FirstName = r.FormValue("firstname")
	user.LastName = r.FormValue("lastname")
	user.Email = r.FormValue("email")
	user.Gender = r.FormValue("gender")
	user.Loan, _ = strconv.ParseFloat(r.FormValue("loan"), 64)
	logers.Debug.Println("fields for Insert:", user)
	user.Insert()
	json.NewEncoder(w).Encode(user)
}
func UpdatePerson(w http.ResponseWriter, r *http.Request) {
	logers.Info.Println("call UpdateUser function")
	var user = model.Person{}
	user.Id = r.FormValue("id")
	user.FirstName = r.FormValue("firstname")
	loan, err := strconv.ParseFloat(r.FormValue("loan"), 64)
	user.Loan = loan
	if err != nil {
		logers.Error.Println(err)
	}
	logers.Debug.Println("fields for Update:", user)
	user.Update()
	json.NewEncoder(w).Encode(user)

}
func DeletePerson(w http.ResponseWriter, r *http.Request) {
	logers.Info.Println("call DeleteUser function")
	var user = model.Person{}
	id, err := strconv.Atoi(r.FormValue("id"))
	if err != nil {
		logers.Error.Println(err)
	}
	logers.Debug.Println("fields for Delete:", user)
	user.Delete(id)
	response := "Ok!"
	json.NewEncoder(w).Encode(response)
}
func UsersFilter(w http.ResponseWriter, r *http.Request) {
	logers.Info.Println("call UserFilter function")
	name := r.FormValue("name")
	firstname := r.FormValue("firstDate")
	secondname := r.FormValue("secondDate")
	gender := r.FormValue("gender")
	logers.Debug.Println("fields for Filter:", name, firstname, secondname, gender)
	users := model.Filters(name, firstname, secondname, gender)
	json.NewEncoder(w).Encode(users)

}
