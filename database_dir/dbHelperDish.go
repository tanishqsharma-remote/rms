package database_dir

import (
	"RMS/model_dir"
	"database/sql"
	"fmt"
	"strings"
)

func InsertDish(item model_dir.NewDish) (sql.Result, error) {
	query := "insert into dishes(rid, name, price, createdby) values ($1,$2,$3,$4)"
	res, err := db.Exec(query, item.RID, item.Name, item.Price, item.CreatedBy)
	return res, err
}
func BatchInsertDish(items []*model_dir.NewDish) (sql.Result, error) {
	valueStrings := []string{}
	valueArgs := []interface{}{}
	j := 0
	for _, i := range items {
		valueStrings = append(valueStrings, fmt.Sprintf("($%d, $%d, $%d, $%d)", j*4+1, j*4+2, j*4+3, j*4+4))

		valueArgs = append(valueArgs, i.RID)
		valueArgs = append(valueArgs, i.Name)
		valueArgs = append(valueArgs, i.Price)
		valueArgs = append(valueArgs, i.CreatedBy)
		j++
	}
	smt := fmt.Sprintf("insert into dishes(rid,name, price, createdby) values %s", strings.Join(valueStrings, ","))
	res, err := db.Exec(smt, valueArgs...)
	return res, err
}
func GetDish(rId int, pageNum string, pageSize string) (*sql.Rows, error) {
	rows, er := db.Query("with pagingCTE as(Select snum,rid,name,price,row_number() over (order by snum)as rowNumber from dishes)select snum,name,price from pagingCTE where rid=$1 and rowNumber between ($2-1)*$3+1 and $2*$3", rId, pageNum, pageSize)
	return rows, er
}
