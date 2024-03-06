package api

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/Yespolovaz/golangFinal/pkg/bank/helpers"
	"github.com/Yespolovaz/golangFinal/pkg/bank/users"
	"github.com/gorilla/mux"
)

// Interfaces
type Login struct {
	Username string
	Password string
}

type Register struct {
	Username string
	Email    string
	Password string
}

type ErrResponse struct {
	Message string
}

// Handlers
func readBody(r *http.Request) []byte {
	// Ready body
	body, err := io.ReadAll(r.Body)
	helpers.HandleErr(err)

	return body
}

// Healthcheck
func info() string {
	return "This is banking API"
}

func healthCheck(w http.ResponseWriter, r *http.Request) {
	helpers.RespondWithJSON(w, http.StatusOK, map[string]string{"status": "ok", "info": info()})
}

func apiResponse(call map[string]interface{}, w http.ResponseWriter) {
	// Prepare response
	if call["message"] == "Success" {
		response := call
		json.NewEncoder(w).Encode(response)
	} else {
		// Handle error
		response := ErrResponse{Message: "Something went wrong"}
		json.NewEncoder(w).Encode(response)
	}
}

func register(w http.ResponseWriter, r *http.Request) {
	body := readBody(r)

	// Handle login
	var formattedBody Register
	var err = json.Unmarshal(body, &formattedBody)
	helpers.HandleErr(err)
	register := users.Register(
		formattedBody.Username,
		formattedBody.Email,
		formattedBody.Password)

	apiResponse(register, w)
}

func login(w http.ResponseWriter, r *http.Request) {
	body := readBody(r)

	// Handle login
	var formattedBody Login
	var err = json.Unmarshal(body, &formattedBody)
	helpers.HandleErr(err)
	login := users.Login(formattedBody.Username, formattedBody.Password)

	apiResponse(login, w)
}

func getUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["id"]
	auth := r.Header.Get("Authorization")

	user := users.GetUser(userId, auth)

	apiResponse(user, w)
}

// API initializator
func StartApi() {
	router := mux.NewRouter()

	router.Use(helpers.PanicHandler)

	router.HandleFunc("/api/health", healthCheck).Methods("GET")
	router.HandleFunc("/api/auth/login", login).Methods("POST")
	router.HandleFunc("/api/auth/register", register).Methods("POST")
	router.HandleFunc("/api/auth/register", register).Methods("GET")
	router.HandleFunc("/api/users/{id}", getUser).Methods("GET")

	// MARK: check port number
	fmt.Println("Listening on port: 5000")
	log.Fatal(http.ListenAndServe(":5000", router))
}
