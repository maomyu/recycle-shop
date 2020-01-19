package model

import "time"

type User struct {
	ID          string `json:"user"`
	Nickname    string `json:"nickname"`
	Password    string `json:"password"`
	Truename    string `json:"truename"`
	Sex         int    `json:"sex"`
	Email       string `json:"email"`
	HeaderImage string `json:"headerimage"`
	School      string `json:"school"`
	Signature   string `json:"signature"`
	Birthday    time.Time
	StudentID   string `json:"studentid"`
}
