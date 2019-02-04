package users

import (
	"fmt"
	"net/http"
	"encoding/json"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type Users struct {
    id    string
    name  string
}

func connect() (*sql.DB, error) {
    db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/belajar_golang")
    if err != nil {
        return nil, err
    }

    return db, nil
}

func AllUser(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	db, err := connect()
    if err != nil {
        fmt.Println(err.Error())
        return
    }
    defer db.Close()

    rows, err := db.Query("select id, name from user")
    if err != nil {
        fmt.Println(err.Error())
        return
    }
    defer rows.Close()

    var result []Users
	var secret []interface{}
    for rows.Next() {
        var each = Users{}
        var err = rows.Scan(&each.id, &each.name)

        if err != nil {
            fmt.Println(err.Error())
            return
        }

		result = append(result, each)
		secret = append(secret, map[string]interface{} {
			"id" : each.id,
			"name" : each.name,
			// "wo" : []int{234},
		})
	}
	fmt.Fprintf(w, "%s", json.NewEncoder(w).Encode(secret))
	
}

func NewUser(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "New user")
}

func DeleteUser(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Delete User")
}

func UpdateUser(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Update User")
}