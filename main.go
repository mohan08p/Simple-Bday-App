package main

import (
	"github.com/Simple-Bday-App/helper"
	"github.com/Simple-Bday-App/models"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
	"strings"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	_ "github.com/go-sql-driver/mysql"
)

var router *chi.Mux
var db *sql.DB

const (
	dbName = "person"
	dbPass = "password"
	dbHost = "mysqldb"
	dbPort = "3306"
)

// init
func init() {
	router = chi.NewRouter()
	router.Use(middleware.Recoverer)

	var err error

	dbSource := fmt.Sprintf("root:%s@tcp(%s:%s)/%s?charset=utf8", dbPass, dbHost, dbPort, dbName)
	db, err = sql.Open("mysql", dbSource)
	helper.Catch(err)
}

func routers() *chi.Mux {
	router.Get("/", ping)

	router.Get("/hello", AllUsers)
	router.Get("/hello/{username}", DetailUser)
	router.Post("/hello/create", CreateUser)
	router.Put("/hello/update/{username}", UpdateUser)
	router.Delete("/hello/{username}", DeleteUser)

	return router
}

// server starting point
func ping(w http.ResponseWriter, r *http.Request) {
	helper.RespondwithJSON(w, http.StatusOK, map[string]string{"message": "Pong"})
}

func AllUsers(w http.ResponseWriter, r *http.Request) {
	errors := []error{}
	payload := []models.User{}

	rows, err := db.Query("Select id, username, date_of_birth From users")
	helper.Catch(err)

	defer rows.Close()

	for rows.Next() {
		data := models.User{}

		er := rows.Scan(&data.ID, &data.Username, &data.DOB)

		if er != nil {
			errors = append(errors, er)
		}
		payload = append(payload, data)
	}

	helper.RespondwithJSON(w, http.StatusOK, payload)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	json.NewDecoder(r.Body).Decode(&user)

	query, err := db.Prepare("Insert users SET id=?, username=?, date_of_birth=?")
	helper.Catch(err)

	_, er := query.Exec(user.ID, user.Username, user.DOB)
	helper.Catch(er)
	defer query.Close()

	helper.RespondwithJSON(w, http.StatusCreated, map[string]string{"message": "successfully created"})
}

func DetailUser(w http.ResponseWriter, r *http.Request) {
	payload := models.User{}
	username := chi.URLParam(r, "username")

	row := db.QueryRow("Select id, username, date_of_birth From users where username=?", username)

	err := row.Scan(
		&payload.ID,
		&payload.Username,
		&payload.DOB,
	)

	loc, _ := time.LoadLocation("UTC")
	now := time.Now().In(loc)
	
	s1 := strings.Split(payload.DOB, "-")
	y1, _ := strconv.Atoi(s1[0])
	y2 := int(now.Year())
	
	layout := "2006-01-02"
	t, _ := time.Parse(layout, payload.DOB)
	diff := now.Sub(t)
	days := (int(diff.Hours() / 24)) % 365
	leapYears := (y2 - y1)/4
	res := 365-(days - leapYears)
	t1 := ""
	if res == 0 {
		t1 = fmt.Sprintln("Hello", payload.Username," Happy birthday.. !! Enjoy !")
	} else {
		t1 = fmt.Sprintln("Hello", payload.Username," your birthday is in", res, "days")
	}

	if err != nil {
		helper.RespondWithError(w, http.StatusNotFound, "no rows in result set")
		return
	}

	response, _ := json.Marshal(t1)
	fmt.Println(t1)
	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	username := chi.URLParam(r, "username")
	json.NewDecoder(r.Body).Decode(&user)

	query, err := db.Prepare("Update users set id=?, date_of_birth=? where username=?")
	helper.Catch(err)
	_, er := query.Exec(user.ID, user.DOB, username)
	helper.Catch(er)

	defer query.Close()

	helper.RespondwithJSON(w, http.StatusOK, map[string]string{"message": "update successfully"})
}

//DeleteUser ....
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	username := chi.URLParam(r, "username")

	query, err := db.Prepare("delete from users where username=?")
	helper.Catch(err)
	_, er := query.Exec(username)
	helper.Catch(er)
	query.Close()

	helper.RespondwithJSON(w, http.StatusOK, map[string]string{"message": "successfully deleted"})
}

func main() {
	routers()
	http.ListenAndServe(":8080", Logger())
}

// Logger return log messages
func Logger() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(time.Now(), r.Method, r.URL)
		router.ServeHTTP(w, r) // dispatch the request
	})
}
