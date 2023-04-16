package goserve

import "time"

type Article struct {
	ID        string    `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	Author    string    `json:"author"`
	Published time.Time `json:"published"`
	Updated   time.Time `json:"updated"`
	Tags      []string  `json:"tags"`
	Views     int       `json:"views"`
	Likes     int       `json:"likes"`
	Comments  []Comment `json:"comments"`
}

type Comment struct {
	ID      string    `json:"id"`
	Author  string    `json:"author"`
	Content string    `json:"content"`
	Posted  time.Time `json:"posted"`
}
