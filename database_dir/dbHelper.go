package database_dir

import (
	"RMS/model_dir"
	"database/sql"
	"time"
)

var db = DbConnect()

func InsertSession(sessionToken string, authorized model_dir.Credentials, Expires time.Time) (sql.Result, error) {
	query := "insert into sessions(sessiontoken, email, expiry) VALUES ($1,$2,$3)"
	res, er := db.Exec(query, sessionToken, authorized.Email, Expires)

	return res, er

}
func GetSession(sessionToken string) (model_dir.Session, error) {
	var sess model_dir.Session
	err := db.Get(&sess, "select email,expiry from sessions where sessiontoken=$1", sessionToken)
	return sess, err
}
func DelSession(sessionToken string) (sql.Result, error) {
	query := "delete from sessions where sessiontoken=$1"
	res, err := db.Exec(query, sessionToken)
	return res, err
}
