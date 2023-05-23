package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func getUsers() []*User {
	// Open up our database connection.
	db, err := sql.Open("mysql", "tester:secret@tcp(db:3306)/test")

	if err != nil {
		log.Print(err.Error())
	}
	defer db.Close()

	results, err := db.Query("SELECT * FROM users")
	if err != nil {
		panic(err.Error())
	}

	var users []*User
	for results.Next() {
		var u User
		err = results.Scan(&u.ID, &u.Name)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}

		users = append(users, &u)
	}

	return users
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "welcome")
	//fmt.Printf("go homepage ")
}

func userPage(w http.ResponseWriter, r *http.Request) {
	users := getUsers()

	//fmt.Println("Endpoint Hit: usersPage")
	json.NewEncoder(w).Encode(users)
}

func main() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/users", userPage)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
