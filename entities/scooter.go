package entities

import "time"

type ScooterUploaded struct {
	PaymentTypeId	int  	`json,scv:"payment_type_id"`
	ModelName 	  	string  `json,csv:"model_name"`
	MaxWeight       int     `json,csv:"max_weight"`
	Speed 			int     `json,scv:"speed"`
	OwnerId  	 	int		`json,scv:"owner_id"`
	SerialNumber 	int		`json,scv:"serial_number"`
}

type Test struct {
	Id		int		`json:"id"`
	Model	string  `json:"model"`
}

type Scooter struct {
	Id		 	 int	`json:"id"`
	ModelId  	 int	`json:"model_id"`
	OwnerId  	 int	`json:"owner_id"`
	SerialNumber int	`json:"serial_number"`
}

type ScooterModel struct {
	Id 		 		  int 	 `json:"id"`
	PaymentTypeId	  int    `json:"payment_type_id"`
	ModelName         string `json:"model_name"`
	MaxWeight         int    `json:"max_weight"`
	Speed 			  int    `json:"speed"`
}

type PaymentType struct {
	Id   int 	 `json:"id"`
	Name string  `json:"name"`
}

type SupplierPrices struct{
	Id 				int 	 `json:"id"`
	Price   		float64  `json:"price"`
	PaymentTypeId   int    	 `json:"payment_type_id"`
	UserId			int 	 `json:"user_id"`
}

type User struct {
	Id  		int 		`json:"id"`
	Email 		string 		`json:"login_email"`
	Blocked		bool 		`json:"is_blocked"`
	Name		string  	`json:"user_name"`
	Surname 	string  	`json:"user_surname"`
	CreatedAt 	time.Time  `json:"created_at"`
	RoleId 		int 		`json:"role_id"`
}
