package model

type Scooter struct {
	Id 		 		  int 	 `json:"id"`
	Model 		  	  string `json:"model"`
	Brand             string `json:"brand"`
	BatteryCapacity   int 	 `json:"battery_capacity"`
	MaxWeight  		  int	 `json:"max_weight"`
	MaxDistance   	  int 	 `json:"max_distance"`
	Created   		  string `json:"created_at"`
	Updated   		  string `json:"updated_at"`
}

type Scooters struct {
	Scooters []ScooterParse `json:"scooter"`
}

type ScooterParse struct {
	Id 		 		  int 	 `csv:"id"`
	Model 		  	  string `csv:"model"`
	Brand             string `csv:"brand"`
	BatteryCapacity   int 	 `csv:"battery_capacity"`
	MaxWeight  		  int	 `csv:"max_weight"`
	MaxDistance   	  int	 `csv:"max_distance"`
}
