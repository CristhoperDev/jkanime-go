package model

type JsonResult struct {
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}

type LastAnimes struct {
	Id     	string
	Title  	string
	Poster 	string
	//Content ContentAnime
}

type ContentAnime struct {
	Type     string
	Gender 	 []string
	Synopsis string
	Status 	 string
	Episodes string
	Pages	 []string
}