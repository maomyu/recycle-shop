package common

type Result struct {
	Success   int         `json:"success"`
	Errorcode int         `json:"errorcode"`
	Message   interface{} `json:"message"`
}

type Data struct {
	Tip string `json:"tip"`
}
