package main

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Headset struct {
	Id     int    `json:"id" gorm:"primary_key"`
	Serial string `json:"serial"`
	Marca  string `json:"marca"`
	Agente string `json:"agente"`
}

type Agent struct {
	Id     int    `json:"id" gorm:"primary_key"`
	Nombre string `json:"nombre"`
	Ciudad string `json:"Ciudad"`
}

type AssignRequest struct {
	HeadsetId int    `json:"headsetId"`
	Agent     string `json:"agent"`
}
