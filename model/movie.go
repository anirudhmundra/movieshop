package model

type Movie struct{
	Id int				`json:"id"`
	Name string			`json:"name"`
	Price float32		`json:"price"`
	Director string		`json:"director"`
}

type Movies struct{
	Movies []Movie `json:"movies"`
}
