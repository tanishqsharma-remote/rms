package handler_dir

import (
	"RMS/database_dir"
	"RMS/model_dir"
	_ "database/sql"
	"encoding/json"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"log"
	"net/http"
)

func CreateByAdmin(w http.ResponseWriter, r *http.Request) {
	var person model_dir.PersonByAdmin
	err := json.NewDecoder(r.Body).Decode(&person)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println(err)
		return
	}
	var newPerson model_dir.NewPerson
	newPerson.Name = person.Name
	newPerson.Email = person.Email
	newPerson.Password = person.Password
	newPerson.CreatedBy = person.CreatedBy

	_, er := database_dir.InsertPerson(newPerson)
	if er != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(er)
		return
	}
	var credentials model_dir.Credentials
	credentials.Email = person.Email
	credentials.Password = person.Password

	var id int
	id, getErr := database_dir.GetId(credentials)
	if getErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(getErr)
		return
	}
	_, roleErr := database_dir.InsertRole(id, person.Role)
	if roleErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(roleErr)
		return
	}
	if person.Role == "user" {
		_, locErr := database_dir.InsertLocation(id, person.Address)
		if locErr != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println(locErr)
			return
		}

	}
}

func CreateBySubAdmin(w http.ResponseWriter, r *http.Request) {
	var person model_dir.UserBySub
	err := json.NewDecoder(r.Body).Decode(&person)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println(err)
		return
	}
	var newPerson model_dir.NewPerson
	newPerson.Name = person.Name
	newPerson.Email = person.Email
	newPerson.Password = person.Password
	newPerson.CreatedBy = person.CreatedBy

	_, er := database_dir.InsertPerson(newPerson)
	if er != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(er)
		return
	}
	var credentials model_dir.Credentials
	credentials.Email = person.Email
	credentials.Password = person.Password
	var id int
	id, getErr := database_dir.GetId(credentials)
	if getErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(getErr)
		return
	}
	_, roleErr := database_dir.InsertRole(id, "user")
	if roleErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(roleErr)
		return
	}
	_, locErr := database_dir.InsertLocation(id, person.Address)
	if locErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(locErr)
		return
	}
}

func CreateByUser(w http.ResponseWriter, r *http.Request) {
	var person model_dir.NewUser
	err := json.NewDecoder(r.Body).Decode(&person)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println(err)
		return
	}
	var newPerson model_dir.NewPerson
	newPerson.Name = person.Name
	newPerson.Email = person.Email
	newPerson.Password = person.Password
	newPerson.CreatedBy = 0

	_, er := database_dir.InsertPerson(newPerson)
	if er != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(er)
		return
	}
	var credentials model_dir.Credentials
	credentials.Email = person.Email
	credentials.Password = person.Password
	var id int
	id, getErr := database_dir.GetId(credentials)
	if getErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(getErr)
		return
	}
	_, roleErr := database_dir.InsertRole(id, "user")
	if roleErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(roleErr)
		return
	}
	_, locErr := database_dir.InsertLocation(id, person.Address)
	if locErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(locErr)
		return
	}
}
func GetSubAdmin(w http.ResponseWriter, r *http.Request) {
	var items []model_dir.SubAdmin
	items, err := database_dir.GetSubAdmin(r.URL.Query().Get("pageNum"), r.URL.Query().Get("pageSize"))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}

	itemsBytes, _ := json.MarshalIndent(items, "", "\t")

	w.Header().Set("Content-Type", "application/json")
	_, er := w.Write(itemsBytes)
	if er != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(er)
		return
	}

}

func GetAllUsers(w http.ResponseWriter, r *http.Request) {

	var items []model_dir.PersonList

	items, err := database_dir.GetPersonDetails(r.URL.Query().Get("pageNum"), r.URL.Query().Get("pageSize"))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}

	itemsBytes, _ := json.MarshalIndent(items, "", "\t")

	w.Header().Set("Content-Type", "application/json")
	_, er := w.Write(itemsBytes)
	if er != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(er)
		return
	}
}
func GetAllUsersBySub(w http.ResponseWriter, r *http.Request) {
	claims, _ := r.Context().Value("email").(jwt.MapClaims)
	id := fmt.Sprint(claims["id"])

	var items []model_dir.PersonList
	items, err := database_dir.GetPersonDetailsBySub(id, r.URL.Query().Get("pageNum"), r.URL.Query().Get("pageSize"))

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}

	itemsBytes, _ := json.MarshalIndent(items, "", "\t")

	w.Header().Set("Content-Type", "application/json")
	_, er := w.Write(itemsBytes)
	if er != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(er)
		return
	}
}

func UpdateLocation(w http.ResponseWriter, r *http.Request) {
	var loc model_dir.NewLocation
	err := json.NewDecoder(r.Body).Decode(&loc)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println(err)
		return
	}
	_, locErr := database_dir.InsertLocation(loc.ID, loc.Address)
	if locErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(locErr)
		return
	}
}
func UpdateRole(w http.ResponseWriter, r *http.Request) {
	var role model_dir.NewRole
	err := json.NewDecoder(r.Body).Decode(&role)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println(err)
		return
	}
	_, roleErr := database_dir.InsertRole(role.ID, role.Role)
	if roleErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(roleErr)
		return
	}
}
