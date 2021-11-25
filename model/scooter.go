package model

type Scooter struct {
	Id 		 		  int 	 `json,scv:"id"`
	Model 		  	  string `json,csv:"model"`
	Brand             string `json,csv:"brand"`
	Capacity  		  int 	 `json,csv:"capacity"`
	MaxWeight  		  int	 `json,csv:"max_weight"`
	MaxDistance   	  int 	 `json,csv:"max_distance"`
	Serial            int    `json,scv:"serial"`
}

type Test struct {
	Id 		 		  int 	 `json:"id"`
	Model 		  	  string `json:"model"`
	Brand             string `json:"brand"`
}

type Scooters struct {
	Scooters []Scooter `json:"scooter"`
}

