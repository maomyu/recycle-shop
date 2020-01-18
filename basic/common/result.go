package common

type Result struct {
	Success   int         `json:"success"`
	Errorcode int         `json:"errorcode"`
	Message   interface{} `message`
}
