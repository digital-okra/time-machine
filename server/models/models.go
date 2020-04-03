package models

type User struct {
	Id       string `json:"id"`
	Username string `json:"username"`
	Utype    string `json:"utype"`
	Amb      int    `json:"amb"`
	Depot    int    `json:"depot"`
	Platoon  int    `json:"platoon"`
	Section  int    `json:"section"`
	Man      int    `json:"man"`
	Name     string `json:"name"`
}

type Task struct {
	Id         string `json:"id"`
	Name       string `json:"name"`
	AssignedTo string `json:"assigned_to"`
	Completed  bool   `json:"completed"`
	Verified   bool   `json:"verified"`
}
