package addrbk

type Exception struct {
    Message string `json:"message"`
}

// Address Entity
type Address struct {
	ID     	   int 		`bson:"_id"`
	FirstName  string      	`json:"firstname"`
	LastName   string       `json:"lastname"`
	Email      string       `json:"email"`
	Phone  	   string       `json:"phone"`
}

type Result struct{
	Status      int       `bson:"status"`
        Res         Address   `json:"address"` 
	ErrMessage  string    `json:"errMessage"`
}
  
//Array of Address Objects
type Addresses []Address
