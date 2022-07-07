package database_dir

import (
	"RMS/model_dir"
	"database/sql"
	"fmt"
	"strings"
	"time"
)

func InsertPerson(db *sql.DB, item model_dir.NewPerson) (sql.Result, error) {
	query := "insert into person(name,email,password,createdby) values($1,$2,$3,$4)"
	res, err := db.Exec(query, item.Name, item.Email, item.Password, item.CreatedBy)
	return res, err
}
func InsertRole(db *sql.DB, id int, role string) (sql.Result, error) {
	query := "insert into roles (pid, role) values ($1,$2)"
	res, err := db.Exec(query, id, role)
	return res, err
}
func InsertLocation(db *sql.DB, id int, address model_dir.Point) (sql.Result, error) {
	query := "insert into location (pid, address) values ($1,point($2,$3))"
	res, err := db.Exec(query, id, address.X, address.Y)
	return res, err
}

func InsertRestaurant(db *sql.DB, item model_dir.NewRestaurant) (sql.Result, error) {
	query := "insert into restaurant(name, address, createdby) values($1,point($2,$3),$4)"
	res, err := db.Exec(query, item.Name, item.Address.X, item.Address.Y, item.CreatedBy)
	return res, err
}
func BatchInsertRestaurant(db *sql.DB, items []*model_dir.NewRestaurant) (sql.Result, error) {
	valueStrings := []string{}
	valueArgs := []interface{}{}
	j := 0
	for _, i := range items {
		valueStrings = append(valueStrings, fmt.Sprintf("($%d, %s($%d, $%d), $%d)", j*4+1, "point", j*4+2, j*4+3, j*4+4))

		valueArgs = append(valueArgs, i.Name)
		valueArgs = append(valueArgs, i.Address.X)
		valueArgs = append(valueArgs, i.Address.Y)
		valueArgs = append(valueArgs, i.CreatedBy)
		j++
	}
	smt := fmt.Sprintf("insert into restaurant(name, address, createdby) values %s", strings.Join(valueStrings, ","))
	res, err := db.Exec(smt, valueArgs...)
	return res, err
}
func InsertDish(db *sql.DB, item model_dir.NewDish) (sql.Result, error) {
	query := "insert into dishes(rid, name, price, createdby) values ($1,$2,$3,$4)"
	res, err := db.Exec(query, item.RId, item.Name, item.Price, item.CreatedBy)
	return res, err
}
func BatchInsertDish(db *sql.DB, items []*model_dir.NewDish) (sql.Result, error) {
	valueStrings := []string{}
	valueArgs := []interface{}{}
	j := 0
	for _, i := range items {
		valueStrings = append(valueStrings, fmt.Sprintf("($%d, $%d, $%d, $%d)", j*4+1, j*4+2, j*4+3, j*4+4))

		valueArgs = append(valueArgs, i.RId)
		valueArgs = append(valueArgs, i.Name)
		valueArgs = append(valueArgs, i.Price)
		valueArgs = append(valueArgs, i.CreatedBy)
		j++
	}
	smt := fmt.Sprintf("insert into dishes(rid,name, price, createdby) values %s", strings.Join(valueStrings, ","))
	res, err := db.Exec(smt, valueArgs...)
	return res, err
}

func InsertSession(db *sql.DB, sessionToken string, authorized model_dir.Credentials, Expires time.Time) (sql.Result, error) {
	query := "insert into sessions(sessiontoken, email, expiry) VALUES ($1,$2,$3)"
	res, er := db.Exec(query, sessionToken, authorized.Email, Expires)
	return res, er

}

func GetSession(db *sql.DB, sessionToken string) (*sql.Rows, error) {
	rows, err := db.Query("select email,expiry from sessions where sessiontoken=$1", sessionToken)
	return rows, err
}
func GetPassword(db *sql.DB, credentials model_dir.Credentials) (*sql.Rows, error) {
	rows, er := db.Query("Select email,password from person where email=$1", credentials.Email)
	return rows, er
}
func GetId(db *sql.DB, credentials model_dir.Credentials) (*sql.Rows, error) {
	rows, er := db.Query("Select id from person where email=$1 and password=$2", credentials.Email, credentials.Password)
	return rows, er
}

/*
func GetPerson(db *sql.DB, pageNum string, pageSize string) (*sql.Rows, error) {
	rows, er := db.Query("with pagingCTE as(Select id, name, email, createdby,row_number() over (order by id) as rowNumber from person ) select id, name, email, createdby from pagingCTE where rowNumber between ($1-1)*$2+1 and $1*$2", pageNum, pageSize)
	return rows, er
}*/

func GetPersonRole(db *sql.DB, email string) (*sql.Rows, error) {
	rows, er := db.Query("select id,r.role as role from person inner join roles r on person.id = r.pid where email=$1", email)
	return rows, er
}
func GetPersonDetails(db *sql.DB, pageNum string, pageSize string) (*sql.Rows, error) {
	rows, er := db.Query("with pagingCTE as(Select id, name, email, createdby, r.role as role, l.address as address, row_number() over (order by id) as rowNumber from person inner join roles r on person.id = r.pid left join location l on person.id = l.pid) select id, name, email, createdby, role, address from pagingCTE where rowNumber between ($1-1)*$2+1 and $1*$2", pageNum, pageSize)
	return rows, er
}
func GetPersonDetailsBySub(db *sql.DB, id string, pageNum string, pageSize string) (*sql.Rows, error) {
	rows, er := db.Query("with pagingCTE as(Select id, name, email, createdby, r.role as role, l.address as address, row_number() over (order by id) as rowNumber from person inner join roles r on person.id = r.pid left join location l on person.id = l.pid where createdby=$3) select id, name, email, createdby, role, address from pagingCTE where rowNumber between ($1-1)*$2+1 and $1*$2", pageNum, pageSize, id)
	return rows, er
}
func GetSubAdmin(db *sql.DB, pageNum string, pageSize string) (*sql.Rows, error) {
	rows, er := db.Query("with pagingCTE as(Select id, name, email, r.role as role, row_number() over (order by id) as rowNumber from person inner join roles r on person.id = r.pid ) select id, name, email, role from pagingCTE where role='sub-admin' and rowNumber between ($1-1)*$2+1 and $1*$2", pageNum, pageSize)
	return rows, er
}
func GetRestaurant(db *sql.DB, pageNum string, pageSize string) (*sql.Rows, error) {
	rows, er := db.Query("with pagingCTE as(Select id, name, address, createdby, row_number() over (order by id)as rowNumber from restaurant)select id,name,address from pagingCTE where rowNumber between ($1-1)*$2+1 and $1*$2", pageNum, pageSize)
	return rows, er
}
func GetRestaurantBySubAdmin(db *sql.DB, id string, pageNum string, pageSize string) (*sql.Rows, error) {
	rows, er := db.Query("with pagingCTE as(Select id, name, address, createdby, row_number() over (order by id)as rowNumber from restaurant where createdby=$3)select id,name,address from pagingCTE where rowNumber between ($1-1)*$2+1 and $1*$2", pageNum, pageSize, id)
	return rows, er
}
func GetDish(db *sql.DB, rId int, pageNum string, pageSize string) (*sql.Rows, error) {
	rows, er := db.Query("with pagingCTE as(Select snum,rid,name,price,row_number() over (order by snum)as rowNumber from dishes)select snum,name,price from pagingCTE where rid=$1 and rowNumber between ($2-1)*$3+1 and $2*$3", rId, pageNum, pageSize)
	return rows, er
}
func GetRestaurantLocation(db *sql.DB, location model_dir.Distance) (*sql.Rows, error) {
	rows, er := db.Query("select point($2,$3)<->(select address from restaurant where id=$1)", location.RId, location.X, location.Y)
	return rows, er
}

func CheckPerson(db *sql.DB, email string) (*sql.Rows, error) {
	rows, er := db.Query("select exists(select 1 from person where email=$1)", email)
	return rows, er
}

func CheckRestaurant(db *sql.DB, name string, address model_dir.Point) (*sql.Rows, error) {
	rows, er := db.Query("select exists(select 1 from restaurant where name=$1 and address~=point($2,$3))", name, address.X, address.Y)
	return rows, er
}

func DelSession(db *sql.DB, sessionToken string) (sql.Result, error) {
	query := "delete from sessions where sessiontoken=$1"
	res, err := db.Exec(query, sessionToken)
	return res, err
}
