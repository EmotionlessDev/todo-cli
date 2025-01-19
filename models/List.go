package models

type List struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Tasks []Task `json:"tasks"`
}
