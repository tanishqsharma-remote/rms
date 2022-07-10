package handler_dir

import (
	"RMS/database_dir"
	"RMS/model_dir"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"log"
	"net/http"
	"strings"
)

func CreateRestaurant(w http.ResponseWriter, r *http.Request) {
	var rest model_dir.NewRestaurant
	err := json.NewDecoder(r.Body).Decode(&rest)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println(err)
		return
	}
	row, chErr := database_dir.CheckRestaurant(rest.Name, rest.Address)
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

	_, er := database_dir.InsertRestaurant(rest)
	if er != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(er)
		return
	}
}
func CreateBulkRestaurant(w http.ResponseWriter, r *http.Request) {
	var rest []*model_dir.NewRestaurant
	err := json.NewDecoder(r.Body).Decode(&rest)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println(err)
		return
	}
	_, er := database_dir.BatchInsertRestaurant(rest)
	if er != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(er)
		return
	}
}

func GetDistance(w http.ResponseWriter, r *http.Request) {
	var loc model_dir.Distance
	err := json.NewDecoder(r.Body).Decode(&loc)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println(err)
		return
	}
	var dist float64
	dist, chErr := database_dir.GetRestaurantLocation(loc)
	if chErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(chErr)
		return
	}
	distance, _ := json.Marshal(dist)
	_, er := w.Write(distance)
	if er != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(er)
		return
	}

}

func GetRestaurantBySub(w http.ResponseWriter, r *http.Request) {
	claims, _ := r.Context().Value("email").(jwt.MapClaims)
	id := fmt.Sprint(claims["id"])

	rows, err := database_dir.GetRestaurantBySubAdmin(id, r.URL.Query().Get("pageNum"), r.URL.Query().Get("pageSize"))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}

	var items []model_dir.RestaurantLister
	for rows.Next() {
		var item model_dir.RestaurantList
		var addItem model_dir.RestaurantLister
		err := rows.Scan(&item.ID, &item.Name, &item.Address)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println(err)
			return
		}
		item.Address = strings.ReplaceAll(item.Address, "(", "")
		item.Address = strings.ReplaceAll(item.Address, ")", "")
		split := strings.Split(item.Address, ",")
		addItem.Address.Lat = split[0]
		addItem.Address.Long = split[1]
		addItem.ID = item.ID
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

	rows, err := database_dir.GetRestaurant(r.URL.Query().Get("pageNum"), r.URL.Query().Get("pageSize"))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}

	var items []model_dir.RestaurantLister
	for rows.Next() {
		var item model_dir.RestaurantList
		var addItem model_dir.RestaurantLister
		err := rows.Scan(&item.ID, &item.Name, &item.Address)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println(err)
			return
		}
		item.Address = strings.ReplaceAll(item.Address, "(", "")
		item.Address = strings.ReplaceAll(item.Address, ")", "")
		split := strings.Split(item.Address, ",")
		addItem.Address.Lat = split[0]
		addItem.Address.Long = split[1]
		addItem.ID = item.ID
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
