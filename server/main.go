package main

import (
	"database/sql"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"net/http"
	"strings"

	"server/models"
)

var db *sql.DB
var err error

const DB_URL string = "./path_to_.db"
const JWT_SECRET string = "password"

func main() {
	// Initialise the global DB pool
	db, err = sql.Open("sqlite3", DB_URL)
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	// Initialise the router
	r := mux.NewRouter()

	log.Fatal(http.ListenAndServe(":8000", r))
}

//---------------------- MIDDLEWARES ------------------------------//
func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET,POST")
		w.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Authorization")

		if r.Method == "OPTIONS" {
			w.WriteHeader(200)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		reqToken := r.Header.Get("Authorization")

		if reqToken == "" {
			http.Error(w, "No auth token", http.StatusForbidden)
			return
		}

		splitToken := strings.Split(reqToken, "Bearer")

		if len(splitToken) != 2 {
			http.Error(w, "Malformed format for auth token", http.StatusForbidden)
			return
		}

		reqToken = strings.TrimSpace(splitToken[1])

		parsedToken, err := jwt.Parse(reqToken, func(token *jwt.Token) (interface{}, error) {
			// Don't forget to validate the alg is what you expect
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, errors.New("Invalid Signing Type")
			}

			return []byte(JWT_SECRET), nil
		})

		// Invalid JWT secret error
		if err != nil {
			http.Error(w, "Authentication failed", http.StatusForbidden)
			return
		}

		// Parsing the claims in the JWT token
		if claims, ok := parsedToken.Claims.(jwt.MapClaims); ok && parsedToken.Valid {
			// If the claims doesn't include the Id or the UserType, throw an error
			if claims["id"] == nil || claims["type"] == nil {
				http.Error(w, "Authentication failed", http.StatusForbidden)
				return
			}

			r.Header.Set("X-User-Claim", claims["id"].(string))
			next.ServeHTTP(w, r)
		} else {
			http.Error(w, "Auth token invalid", http.StatusForbidden)
			return
		}
	})
}

//------------------------------ HANDLERS (Login) ----------------------------------//
// These handlers specifically bypass the authentication middleware, because they do not need any verification
func loginUser(w http.ResponseWriter, r *http.Request) {

}

func registerUser(w http.ResponseWriter, r *http.Request) {

}

//----------------------------- HANDLERS (User) ------------------------------------//
func getUserById(w http.ResponseWriter, r *http.Request) {

}

func getAllAccessibleUsers(w http.ResponseWriter, r *http.Request) {

}

//---------------------------- HANDLERS (Task) ------------------------------------//
func getTasks(w http.ResponseWriter, r *http.Request) {

}

func createTask(w http.ResponseWriter, r *http.Request) {

}

func completeTask(w http.ResponseWriter, r *http.Request) {

}

func verifyTask(w http.ResponseWriter, r *http.Request) {

}
