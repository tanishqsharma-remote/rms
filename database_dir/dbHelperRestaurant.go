package database_dir

import (
	"RMS/model_dir"
	"database/sql"
	"fmt"
	"strings"
)

func InsertRestaurant(item model_dir.NewRestaurant) (sql.Result, error) {
	query := "insert into restaurant(name, address, createdby) values($1,point($2,$3),$4)"
	res, err := db.Exec(query, item.Name, item.Address.Lat, item.Address.Long, item.CreatedBy)
	return res, err
}
func BatchInsertRestaurant(items []*model_dir.NewRestaurant) (sql.Result, error) {
	valueStrings := []string{}
	valueArgs := []interface{}{}
	j := 0
	for _, i := range items {
		valueStrings = append(valueStrings, fmt.Sprintf("($%d, %s($%d, $%d), $%d)", j*4+1, "point", j*4+2, j*4+3, j*4+4))

		valueArgs = append(valueArgs, i.Name)
		valueArgs = append(valueArgs, i.Address.Lat)
		valueArgs = append(valueArgs, i.Address.Long)
		valueArgs = append(valueArgs, i.CreatedBy)
		j++
	}
	smt := fmt.Sprintf("insert into restaurant(name, address, createdby) values %s", strings.Join(valueStrings, ","))
	res, err := db.Exec(smt, valueArgs...)
	return res, err
}

func GetRestaurant(pageNum string, pageSize string) (*sql.Rows, error) {
	rows, er := db.Query("with pagingCTE as(Select id, name, address, createdby, row_number() over (order by id)as rowNumber from restaurant)select id,name,address from pagingCTE where rowNumber between ($1-1)*$2+1 and $1*$2", pageNum, pageSize)
	return rows, er
}
func GetRestaurantBySubAdmin(id string, pageNum string, pageSize string) (*sql.Rows, error) {
	rows, er := db.Query("with pagingCTE as(Select id, name, address, createdby, row_number() over (order by id)as rowNumber from restaurant where createdby=$3)select id,name,address from pagingCTE where rowNumber between ($1-1)*$2+1 and $1*$2", pageNum, pageSize, id)
	return rows, er
}
func GetRestaurantLocation(location model_dir.Distance) (float64, error) {
	var loc float64
	er := db.Get(&loc, "select point($2,$3)<->(select address from restaurant where id=$1)", location.RID, location.Lat, location.Long)
	return loc, er
}

func CheckRestaurant(name string, address model_dir.Point) (*sql.Rows, error) {
	rows, er := db.Query("select exists(select 1 from restaurant where name=$1 and address~=point($2,$3))", name, address.Lat, address.Long)
	return rows, er
}
