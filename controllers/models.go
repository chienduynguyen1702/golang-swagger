package controllers

type Student struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type Teacher struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Course struct {
	ID       int       `json:"id"`
	Name     string    `json:"name"`
	Lectures []Teacher `json:"lectures"`
	Semester string    `json:"semester"`
	Students []Student `json:"students"`
}
