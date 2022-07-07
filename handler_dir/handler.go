package handler_dir

import (
	"RMS/database_dir"
	"RMS/model_dir"
	"bytes"
	"database/sql"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

func CreateByAdmin(w http.ResponseWriter, r *http.Request) {
	db := database_dir.DBconnect()
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

	row, chErr := database_dir.CheckPerson(db, person.Email)
	if chErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(chErr)
		return
	}
	var check bool
	for row.Next() {
		ScanErr := row.Scan(&check)
		if ScanErr != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println(ScanErr)
			return
		}
	}
	if check {
		w.WriteHeader(http.StatusBadRequest)
		_, wErr := w.Write([]byte("Username already registered"))
		if wErr != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println(wErr)
			return
		}
		return
	}
	_, er := database_dir.InsertPerson(db, newPerson)
	if er != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(er)
		return
	}
	var credentials model_dir.Credentials
	credentials.Email = person.Email
	credentials.Password = person.Password
	rows, getErr := database_dir.GetId(db, credentials)
	if getErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(getErr)
		return
	}
	var id int
	for rows.Next() {
		ScanErr := rows.Scan(&id)
		if ScanErr != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println(ScanErr)
			return
		}
	}
	_, roleErr := database_dir.InsertRole(db, id, person.Role)
	if roleErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(roleErr)
		return
	}
	if person.Role == "user" {
		_, locErr := database_dir.InsertLocation(db, id, person.Address)
		if locErr != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println(locErr)
			return
		}

	}
}

func CreateBySubAdmin(w http.ResponseWriter, r *http.Request) {
	db := database_dir.DBconnect()
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

	row, chErr := database_dir.CheckPerson(db, person.Email)
	if chErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(chErr)
		return
	}
	var check bool
	for row.Next() {
		ScanErr := row.Scan(&check)
		if ScanErr != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println(ScanErr)
			return
		}
	}
	if check {
		w.WriteHeader(http.StatusBadRequest)
		_, wErr := w.Write([]byte("Username already registered"))
		if wErr != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println(wErr)
			return
		}
		return
	}

	_, er := database_dir.InsertPerson(db, newPerson)
	if er != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(er)
		return
	}
	var credentials model_dir.Credentials
	credentials.Email = person.Email
	credentials.Password = person.Password
	rows, getErr := database_dir.GetId(db, credentials)
	if getErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(getErr)
		return
	}
	var id int
	for rows.Next() {
		ScanErr := rows.Scan(&id)
		if ScanErr != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println(ScanErr)
			return
		}
	}
	_, roleErr := database_dir.InsertRole(db, id, "user")
	if roleErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(roleErr)
		return
	}
	_, locErr := database_dir.InsertLocation(db, id, person.Address)
	if locErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(locErr)
		return
	}
}

func CreateByUser(w http.ResponseWriter, r *http.Request) {
	db := database_dir.DBconnect()
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
	row, chErr := database_dir.CheckPerson(db, person.Email)
	if chErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(chErr)
		return
	}
	var check bool
	for row.Next() {
		ScanErr := row.Scan(&check)
		if ScanErr != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println(ScanErr)
			return
		}
	}
	if check {
		w.WriteHeader(http.StatusBadRequest)
		_, wErr := w.Write([]byte("Username already registered"))
		if wErr != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println(wErr)
			return
		}
		return
	}

	_, er := database_dir.InsertPerson(db, newPerson)
	if er != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(er)
		return
	}
	var credentials model_dir.Credentials
	credentials.Email = person.Email
	credentials.Password = person.Password
	rows, getErr := database_dir.GetId(db, credentials)
	if getErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(getErr)
		return
	}
	var id int
	for rows.Next() {
		ScanErr := rows.Scan(&id)
		if ScanErr != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println(ScanErr)
			return
		}
	}
	_, roleErr := database_dir.InsertRole(db, id, "user")
	if roleErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(roleErr)
		return
	}
	_, locErr := database_dir.InsertLocation(db, id, person.Address)
	if locErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(locErr)
		return
	}
}
func CreateDish(w http.ResponseWriter, r *http.Request) {
	db := database_dir.DBconnect()
	var dish model_dir.NewDish
	err := json.NewDecoder(r.Body).Decode(&dish)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println(err)
		return
	}
	_, er := database_dir.InsertDish(db, dish)
	if er != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(er)
		return
	}
}
func CreateRestaurant(w http.ResponseWriter, r *http.Request) {
	db := database_dir.DBconnect()
	var rest model_dir.NewRestaurant
	err := json.NewDecoder(r.Body).Decode(&rest)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println(err)
		return
	}
	row, chErr := database_dir.CheckRestaurant(db, rest.Name, rest.Address)
	if chErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(chErr)
		return
	}
	var check bool
	for row.Next() {
		ScanErr := row.Scan(&check)
		if ScanErr != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println(ScanErr)
			return
		}
	}
	if check {
		w.WriteHeader(http.StatusBadRequest)
		_, wErr := w.Write([]byte("The Restaurant at given location already exists"))
		if wErr != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println(wErr)
			return
		}
		return
	}

	_, er := database_dir.InsertRestaurant(db, rest)
	if er != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(er)
		return
	}
}

func CreateBulkDish(w http.ResponseWriter, r *http.Request) {
	db := database_dir.DBconnect()
	var dish []*model_dir.NewDish
	err := json.NewDecoder(r.Body).Decode(&dish)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println(err)
		return
	}
	_, er := database_dir.BatchInsertDish(db, dish)
	if er != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(er)
		return
	}
}
func CreateCsvDish(w http.ResponseWriter, r *http.Request) {
	db := database_dir.DBconnect()
	f, _, _ := r.FormFile("csvFile")
	buf := bytes.NewBuffer(nil)
	if _, err := io.Copy(buf, f); err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	create, createErr := os.Create("temp.csv")
	if createErr != nil {
		log.Print(createErr)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	_, copyErr := io.Copy(create, buf)
	if copyErr != nil {
		log.Print(copyErr)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	csvFile, err := os.Open("./temp.csv")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}
	defer func(csvFile *os.File) {
		err := csvFile.Close()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Print(err)
			return
		}
	}(csvFile)
	csvLines, err := csv.NewReader(csvFile).ReadAll()
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	var dishes []*model_dir.NewDish
	for i, line := range csvLines {
		if i == 0 {
			continue
		}
		var dish model_dir.NewDish
		dish.RId, _ = strconv.Atoi(line[0])
		dish.Name = line[1]
		dish.Price, _ = strconv.Atoi(line[2])
		dish.Price, _ = strconv.Atoi(line[3])
		dishes = append(dishes, &dish)
	}
	_, er := database_dir.BatchInsertDish(db, dishes)
	if er != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(er)
		return
	}

	remErr := os.Remove("./temp.csv")
	if remErr != nil {
		log.Print(remErr)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
func CreateBulkRestaurant(w http.ResponseWriter, r *http.Request) {
	db := database_dir.DBconnect()
	var rest []*model_dir.NewRestaurant
	err := json.NewDecoder(r.Body).Decode(&rest)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println(err)
		return
	}
	_, er := database_dir.BatchInsertRestaurant(db, rest)
	if er != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(er)
		return
	}
}

func GetDistance(w http.ResponseWriter, r *http.Request) {
	db := database_dir.DBconnect()
	var loc model_dir.Distance
	err := json.NewDecoder(r.Body).Decode(&loc)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println(err)
		return
	}
	row, chErr := database_dir.GetRestaurantLocation(db, loc)
	if chErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(chErr)
		return
	}
	var dist float64
	for row.Next() {
		ScanErr := row.Scan(&dist)
		if ScanErr != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println(ScanErr)
			return
		}
	}
	distance, _ := json.Marshal(dist)
	_, er := w.Write(distance)
	if er != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(er)
		return
	}

}

func GetSubAdmin(w http.ResponseWriter, r *http.Request) {

	db := database_dir.DBconnect()
	rows, err := database_dir.GetSubAdmin(db, r.URL.Query().Get("pageNum"), r.URL.Query().Get("pageSize"))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}

	var items []model_dir.SubAdmin

	for rows.Next() {
		var item model_dir.SubAdmin
		err := rows.Scan(&item.Id, &item.Name, &item.Email, &item.Role)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println(err)
			return
		}
		items = append(items, item)
	}

	itemsBytes, _ := json.MarshalIndent(items, "", "\t")

	w.Header().Set("Content-Type", "application/json")
	_, er := w.Write(itemsBytes)
	if er != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(er)
		return
	}

	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println(err)
			return
		}
	}(rows)
}

func GetRestaurantBySub(w http.ResponseWriter, r *http.Request) {
	claims, _ := r.Context().Value("email").(jwt.MapClaims)
	id := fmt.Sprint(claims["id"])

	db := database_dir.DBconnect()
	rows, err := database_dir.GetRestaurantBySubAdmin(db, id, r.URL.Query().Get("pageNum"), r.URL.Query().Get("pageSize"))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}

	var items []model_dir.RestaurantLister
	for rows.Next() {
		var item model_dir.RestaurantList
		var addItem model_dir.RestaurantLister
		err := rows.Scan(&item.Id, &item.Name, &item.Address)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println(err)
			return
		}
		item.Address = strings.ReplaceAll(item.Address, "(", "")
		item.Address = strings.ReplaceAll(item.Address, ")", "")
		split := strings.Split(item.Address, ",")
		addItem.Address.X = split[0]
		addItem.Address.Y = split[1]
		addItem.Id = item.Id
		addItem.Name = item.Name
		items = append(items, addItem)
	}

	itemsBytes, _ := json.MarshalIndent(items, "", "\t")

	w.Header().Set("Content-Type", "application/json")
	_, er := w.Write(itemsBytes)
	if er != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(er)
		return
	}

	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println(err)
			return
		}
	}(rows)
}

func GetRestaurant(w http.ResponseWriter, r *http.Request) {

	db := database_dir.DBconnect()
	rows, err := database_dir.GetRestaurant(db, r.URL.Query().Get("pageNum"), r.URL.Query().Get("pageSize"))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}

	var items []model_dir.RestaurantLister
	for rows.Next() {
		var item model_dir.RestaurantList
		var addItem model_dir.RestaurantLister
		err := rows.Scan(&item.Id, &item.Name, &item.Address)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println(err)
			return
		}
		item.Address = strings.ReplaceAll(item.Address, "(", "")
		item.Address = strings.ReplaceAll(item.Address, ")", "")
		split := strings.Split(item.Address, ",")
		addItem.Address.X = split[0]
		addItem.Address.Y = split[1]
		addItem.Id = item.Id
		addItem.Name = item.Name
		items = append(items, addItem)
	}

	itemsBytes, _ := json.MarshalIndent(items, "", "\t")

	w.Header().Set("Content-Type", "application/json")
	_, er := w.Write(itemsBytes)
	if er != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(er)
		return
	}

	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println(err)
			return
		}
	}(rows)
}
func GetDishes(w http.ResponseWriter, r *http.Request) {
	var restId model_dir.RestaurantId
	err := json.NewDecoder(r.Body).Decode(&restId)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println(err)
		return
	}
	db := database_dir.DBconnect()
	rows, err := database_dir.GetDish(db, restId.RId, r.URL.Query().Get("pageNum"), r.URL.Query().Get("pageSize"))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}

	var items []model_dir.DishList

	for rows.Next() {
		var item model_dir.DishList
		err := rows.Scan(&item.SNum, &item.DishName, &item.Price)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println(err)
			return
		}
		items = append(items, item)
	}

	itemsBytes, _ := json.MarshalIndent(items, "", "\t")

	w.Header().Set("Content-Type", "application/json")
	_, er := w.Write(itemsBytes)
	if er != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(er)
		return
	}

	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println(err)
			return
		}
	}(rows)
}
func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	db := database_dir.DBconnect()

	rows, err := database_dir.GetPersonDetails(db, r.URL.Query().Get("pageNum"), r.URL.Query().Get("pageSize"))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}

	var items []model_dir.PersonList

	for rows.Next() {
		var item model_dir.PersonList
		err := rows.Scan(&item.Id, &item.Name, &item.Email, &item.CreatedBy, &item.Role, &item.Address)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println(err)
			return
		}
		items = append(items, item)
	}

	itemsBytes, _ := json.MarshalIndent(items, "", "\t")

	w.Header().Set("Content-Type", "application/json")
	_, er := w.Write(itemsBytes)
	if er != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(er)
		return
	}

	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println(err)
			return
		}
	}(rows)
}
func GetAllUsersBySub(w http.ResponseWriter, r *http.Request) {
	db := database_dir.DBconnect()
	claims, _ := r.Context().Value("email").(jwt.MapClaims)
	id := fmt.Sprint(claims["id"])
	rows, err := database_dir.GetPersonDetailsBySub(db, id, r.URL.Query().Get("pageNum"), r.URL.Query().Get("pageSize"))

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}

	var items []model_dir.PersonList

	for rows.Next() {
		var item model_dir.PersonList
		err := rows.Scan(&item.Id, &item.Name, &item.Email, &item.CreatedBy, &item.Role, &item.Address)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println(err)
			return
		}
		items = append(items, item)
	}

	itemsBytes, _ := json.MarshalIndent(items, "", "\t")

	w.Header().Set("Content-Type", "application/json")
	_, er := w.Write(itemsBytes)
	if er != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(er)
		return
	}

	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println(err)
			return
		}
	}(rows)
}

func Login(w http.ResponseWriter, r *http.Request) {
	db := database_dir.DBconnect()
	var credentials model_dir.Credentials
	err := json.NewDecoder(r.Body).Decode(&credentials)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println(err)
		return
	}
	rows, er := database_dir.GetPassword(db, credentials)
	if er != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(er)
		return
	}
	var authorized model_dir.Credentials
	for rows.Next() {
		ScanErr := rows.Scan(&authorized.Email, &authorized.Password)
		if ScanErr != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println(ScanErr)
			return
		}
	}
	if credentials.Password != authorized.Password {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	ExpiryTime := time.Now().Add(time.Minute * 30).Unix()
	Expires := time.Now().Add(time.Minute * 30)
	sessionToken := uuid.NewString()

	_, exErr := database_dir.InsertSession(db, sessionToken, authorized, Expires)
	if exErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(exErr)
		return
	}
	w.Header().Add("sessionToken", sessionToken)

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	var auth model_dir.Authentication

	rowRole, er := database_dir.GetPersonRole(db, authorized.Email)
	if er != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(er)
		return
	}
	for rowRole.Next() {
		ScanErr := rowRole.Scan(&auth.Id, &auth.Role)
		if ScanErr != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println(ScanErr)
			return
		}
	}

	claims["email"] = authorized.Email
	claims["exp"] = ExpiryTime
	claims["id"] = auth.Id
	claims["role"] = auth.Role
	userTokenString, SignErr := token.SignedString(model_dir.JwtKey)
	if SignErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(SignErr)
		return
	}
	var userToken model_dir.Token
	userToken.Email = authorized.Email
	userToken.TokenString = userTokenString
	EncodeErr := json.NewEncoder(w).Encode(userToken)
	if EncodeErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(EncodeErr)
		return
	}

}
func Logout(w http.ResponseWriter, r *http.Request) {
	db := database_dir.DBconnect()
	sessionToken := r.Header.Get("sessionToken")
	_, execErr := database_dir.DelSession(db, sessionToken)
	if execErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(execErr)
		return
	}
	_, er := io.WriteString(w, "Successfully Logged out")
	if er != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(er)
		return
	}
}
