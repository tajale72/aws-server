package model

//User is a struct
type User struct {
	Name   string   `json:"name"`
	Age    int      `json:"age"`
	Arr    []string `json:"arr"`
	Origin string   `json:origin`
}
