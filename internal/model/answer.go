package model 


type Answer struct {
	Question string `json:"question"`
	Answer string `json:"answer"`
	BeforeAnswer string `json:"beforeAnswer"`
}