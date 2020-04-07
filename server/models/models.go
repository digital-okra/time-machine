package models

type User struct {
	Id        string `json:"id"`
	Username  string `json:"username"`
	Utype     string `json:"utype"`
	Amb       int    `json:"amb"`
	Depot     int    `json:"depot"`
	Platoon   int    `json:"platoon"`
	Section   int    `json:"section"`
	Man       int    `json:"man"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Rank      string `json:"rank"`
}

type Task struct {
	Id         string `json:"id"`
	Name       string `json:"name"`
	AssignedTo string `json:"assigned_to"`
	AssignedBy string `json:"assigned_by"`
	Completed  bool   `json:"completed"`
	Verified   bool   `json:"verified"`
	VerifiedBy string `json:"verified_by"`
}
