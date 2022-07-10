package handler_dir

import (
	"RMS/database_dir"
	"RMS/model_dir"
	"database/sql"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
)

func CreateDish(w http.ResponseWriter, r *http.Request) {
	var dish model_dir.NewDish
	err := json.NewDecoder(r.Body).Decode(&dish)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println(err)
		return
	}
	_, er := database_dir.InsertDish(dish)
	if er != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(er)
		return
	}
}
func CreateBulkDish(w http.ResponseWriter, r *http.Request) {
	var dish []*model_dir.NewDish
	err := json.NewDecoder(r.Body).Decode(&dish)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println(err)
		return
	}
	_, er := database_dir.BatchInsertDish(dish)
	if er != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(er)
		return
	}
}
func CreateCsvDish(w http.ResponseWriter, r *http.Request) {
	f, _, _ := r.FormFile("csvFile")

	create, createErr := os.Create("temp.csv")
	if createErr != nil {
		log.Print(createErr)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	_, copyErr := io.Copy(create, f)
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
		dish.RID, _ = strconv.Atoi(line[0])
		dish.Name = line[1]
		dish.Price, _ = strconv.Atoi(line[2])
		dish.CreatedBy, _ = strconv.Atoi(line[3])
		dishes = append(dishes, &dish)
	}
	_, er := database_dir.BatchInsertDish(dishes)
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
func GetDishes(w http.ResponseWriter, r *http.Request) {
	var restId model_dir.RestaurantId
	err := json.NewDecoder(r.Body).Decode(&restId)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println(err)
		return
	}
	rows, err := database_dir.GetDish(restId.RID, r.URL.Query().Get("pageNum"), r.URL.Query().Get("pageSize"))
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
