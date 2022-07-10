package database_dir

import (
	"RMS/model_dir"
	"database/sql"
)

func InsertPerson(item model_dir.NewPerson) (sql.Result, error) {
	query := "insert into person(name,email,password,createdby) values($1,$2,$3,$4)"
	res, err := db.Exec(query, item.Name, item.Email, item.Password, item.CreatedBy)
	return res, err
}
func InsertRole(id int, role string) (sql.Result, error) {
	query := "insert into roles (pid, role) values ($1,$2)"
	res, err := db.Exec(query, id, role)
	return res, err
}
func InsertLocation(id int, address model_dir.Point) (sql.Result, error) {
	query := "insert into location (pid, address) values ($1,point($2,$3))"
	res, err := db.Exec(query, id, address.Lat, address.Long)
	return res, err
}

func GetPassword(credentials model_dir.Credentials) (model_dir.Credentials, error) {
	var pass model_dir.Credentials
	er := db.Get(&pass, "Select email,password from person where email=$1", credentials.Email)
	return pass, er
}
func GetId(credentials model_dir.Credentials) (int, error) {
	var id int
	er := db.Get(&id, "Select id from person where email=$1 and password=$2", credentials.Email, credentials.Password)
	return id, er
}

func GetPersonRole(email string) (model_dir.Authentication, error) {
	var auth model_dir.Authentication
	er := db.Get(&auth, "with roleCte as(select  id,r.role as role, case  when r.role = 'admin'::roleType then 1 else case when r.role= 'sub-admin' then 2 else 3 end end as precedence from person p join roles r on p.id = r.pid where email=$1 order by p.id, (role))select id,role from roleCte order by (precedence) limit 1;", email)
	return auth, er
}
func GetPersonDetails(pageNum string, pageSize string) ([]model_dir.PersonList, error) {
	var items []model_dir.PersonList
	er := db.Select(&items, "with pagingCTE as(Select id, name, email, createdby, r.role as role, l.address as address, row_number() over (order by id) as rowNumber from person inner join roles r on person.id = r.pid left join location l on person.id = l.pid) select id, name, email, createdby, role, address from pagingCTE where rowNumber between ($1-1)*$2+1 and $1*$2", pageNum, pageSize)
	return items, er
}
func GetPersonDetailsBySub(id string, pageNum string, pageSize string) ([]model_dir.PersonList, error) {
	var items []model_dir.PersonList
	er := db.Select(&items, "with pagingCTE as(Select id, name, email, createdby, r.role as role, l.address as address, row_number() over (order by id) as rowNumber from person inner join roles r on person.id = r.pid left join location l on person.id = l.pid where createdby=$3) select id, name, email, createdby, role, address from pagingCTE where rowNumber between ($1-1)*$2+1 and $1*$2", pageNum, pageSize, id)
	return items, er
}
func GetSubAdmin(pageNum string, pageSize string) ([]model_dir.SubAdmin, error) {
	var items []model_dir.SubAdmin
	er := db.Select(&items, "with pagingCTE as(Select id, name, email, r.role as role, row_number() over (order by id) as rowNumber from person inner join roles r on person.id = r.pid ) select id, name, email, role from pagingCTE where role='sub-admin' and rowNumber between ($1-1)*$2+1 and $1*$2", pageNum, pageSize)
	return items, er
}

//
//func CheckPerson(email string) (bool, error) {
//	var check bool
//	er := db.Select(&check,"select exists(select 1 from person where email=$1)", email)
//	return check, er
//}
