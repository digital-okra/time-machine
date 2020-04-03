package main

import (
	"database/sql"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"github.com/lithammer/shortuuid"
	_ "github.com/mattn/go-sqlite3"

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

			uid := claims["id"].(string)
			utype := claims["type"].(string)

			r.Header.Set("X-User-Claim", uid)
			r.Header.Set("X-User-Type", utype)

			next.ServeHTTP(w, r)
		} else {
			http.Error(w, "Auth token invalid", http.StatusForbidden)
			return
		}
	})
}

//------------------------------ HANDLERS (Login) ----------------------------------//
// These handlers specifically bypass the authentication middleware, because they do not need any verification
type loginUserRequest struct {
	Username     string `json:"username"`
	PasswordHash string `json:"password_hash"`
}
type registerUserRequest struct {
	Username     string `json:"username"`
	PasswordHash string `json:"password_hash"`
	Type         string `json:"type"`
	Platoon      int    `json:"platoon"`
	Section      int    `json:"section"`
	Man          int    `json:"section"`
	Name         string `json:"name"`
}

func loginUser(w http.ResponseWriter, r *http.Request) {
	var req loginUserRequest

	// Decode the request
	err := json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var uid string
	var passwordhash string

	// Get the user associated to the username if it exists
	sql := `SELECT user, password_hash FROM user WHERE username = ?`
	if err := db.QueryRow(sql, req.Username).Scan(&uid, &passwordhash); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Check if password hashes match then generate JWT
	if passwordhash == req.PasswordHash {
		token, err := createJWT(uid, JWT_SECRET)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Write([]byte(token))
	} else {
		// Password incorrect, throw unauthorized error
		http.Error(w, err.Error(), http.StatusUnauthorized)
	}
}

func registerUser(w http.ResponseWriter, r *http.Request) {
	var req registerUserRequest

	// Decode the request
	err := json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Generate unique uid
	uid := shortuuid.New()

	// Create the SQL prepared statement
	sql := `INSERT INTO user VALUES ?, ?, ?, ?, ?, ?, ?, ?`
	stmt, err := db.Prepare(sql)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	//The existence of the actual content of the parsed request does not need to be checked as it is verified by the NOT NULL constraints

	// Execute the statement
	_, err = stmt.Exec(uid, req.Username, req.PasswordHash, req.Type, req.Platoon, req.Section, req.Man, req.Name)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	// Return the new JWT
	token, err := createJWT(uid, JWT_SECRET)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write([]byte(token))
}

//----------------------------- HANDLERS (User) ------------------------------------//
func getUserById(w http.ResponseWriter, r *http.Request) {
	// Get the current user id
	uid := r.Header.Get("X-User-Claim")

	var user models.User

	// Get the user associated to the id if it exists
	sql := `SELECT user, username, utype, platoon, section, man, name FROM user WHERE user = ?`
	if err := db.QueryRow(sql, uid).Scan(&user.Id, &user.Username, &user.Utype, &user.Platoon, &user.Section, &user.Man, &user.Name); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Marshal to JSON and return
	res, err := json.Marshal(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Write(res)
}

func getAllAccessibleUsers(w http.ResponseWriter, r *http.Request) {
	uid := r.Header.Get("X-User-Claim")
	utype := r.Header.Get("X-User-Type")

	if utype != "admin" {
		http.Error(w, "No admin permissions for this user", http.StatusForbidden)
	}

	// Retrive the platoon/section that the admin is in charge of
	var platoon int
	var section int

	// Get the platoon,section associated to the admin user if it exists
	sql := `SELECT user, username, utype, platoon, section, man, name FROM user WHERE user = ?`
	if err := db.QueryRow(sql, uid).Scan(&platoon, &section); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var users []models.User

	// Get all the users until the admin user
	sql := `SELECT user, username, platoon, section, man, name FROM user WHERE platoon = ? AND section = ?`
	result, err := db.Query(sql, platoon, section)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	defer result.Close()

	for result.Next() {
		var user models.User
		if err := result.Scan(&user.Id, &user.Username, &user.Platoon, &user.Section, &user.Man, &user.Name); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		users = append(users, user)
	}

	// Marshal to JSON and return
	res, err := json.Marshal(users)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Write(res)
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

//------------------------ UTILITIES -----------------------------------------------//
func createJWT(uid string, secret string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id": uid,
	})
	tokenString, err := token.SignedString([]byte(secret))

	return tokenString, err
}
