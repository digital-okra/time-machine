package models

type User struct {
	Id       string `json:"id"`
	Username string `json:"username"`
	Utype    string `json:"utype"`
	Name     string `json:"name"`
}

type Task struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	Completed bool   `json:"completed"`
	Verified  bool   `json:"verified"`
}
