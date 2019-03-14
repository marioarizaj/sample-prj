package models

import "time"

type Article struct {
	Id int `json:"id"`
	Author string `json:"author"`
	Title string `json:"title"`
	Body string `json:"body"`
	DateCreated time.Time `json:"date_created"`
}