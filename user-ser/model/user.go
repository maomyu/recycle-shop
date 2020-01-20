package model

import "time"

type User struct {
	ID          string    `json:"id"`
	Nickname    string    `json:"nickname"`
	Password    string    `json:"password"`
	Truename    string    `json:"truename"`
	Sex         int       `json:"sex"`
	Email       string    `json:"email"`
	HeaderImage string    `json:"headerimage"`
	School      string    `json:"school"`
	Signature   string    `json:"signature"`
	Birthday    time.Time `json:"birthday"`
	StudentID   string    `json:"studentid"`
	Role        int       `json:"role"`
}
type RegisterResult struct {
	ID          string    `json:"id"`
	Nickname    string    `json:"nickname"`
	Sex         int       `json:"sex"`
	Email       string    `json:"email"`
	HeaderImage string    `json:"headerimage"`
	School      string    `json:"school"`
	Signature   string    `json:"signature"`
	Birthday    time.Time `json:"birthday"`
	StudentID   string    `json:"studentid"`
	Isrealname  int       `json:"isrealname"` //是否实名
	Follow      int       `json:"follow"`
	Fans        int       `json:"fans"`
	Creditscore int       `json:"creditscore"` //信用值
}
