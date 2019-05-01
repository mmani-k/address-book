package addrbk

import (
	"fmt"
	"log"
	"gopkg.in/mgo.v2"
	"strings"
	"gopkg.in/mgo.v2/bson"
)

//Repository ...
type Repository struct{}

// SERVER the DB server
const SERVER = "mongodb://localhost:27017/"

// DBNAME themyadmin:password@ name of the DB instance
const DBNAME = "addressbook"


// COLLECTION is the name of the collection in DB
const COLLECTION = "addresses"

var addressId = 10;

// GetProducts returns the list of Products
func (r Repository) GetAddresses() Addresses {
	session, err := mgo.Dial(SERVER)

	if err != nil {
	 	fmt.Println("Failed to establish connection to Mongo server:", err)
	}

	defer session.Close()

	c := session.DB(DBNAME).C(COLLECTION)
	results := Addresses{}

	if err := c.Find(nil).All(&results); err != nil {
	  	fmt.Println("Failed to write results:", err)
	}

	return results
}

// GetProductById returns a unique Product
func (r Repository) GetAddressById(id int) Result {
	session, err := mgo.Dial(SERVER)

	if err != nil {
	 	fmt.Println("Failed to establish connection to Mongo server:", err)
	}

	defer session.Close()

	c := session.DB(DBNAME).C(COLLECTION)
	var result Address

	var res Result

	if err := c.FindId(id).One(&result); err != nil {
	  	fmt.Println("Failed to get address:", err)
		res.Status = 1	
                res.ErrMessage = "Address not found!"
	}
	if res.Status == 0 {
		res.Res = result
	}

	return res
}

// GetProductsByString takes a search string as input and returns products
func (r Repository) GetAddressByString(query string) Addresses {
	session, err := mgo.Dial(SERVER)

	if err != nil {
	 	fmt.Println("Failed to establish connection to Mongo server:", err)
	}

	defer session.Close()

	c := session.DB(DBNAME).C(COLLECTION)
	result := Addresses{}

	// Logic to create filter
	qs := strings.Split(query, " ")
	and := make([]bson.M, len(qs))
	for i, q := range qs {
    	and[i] = bson.M{"title": bson.M{
        	"$regex": bson.RegEx{Pattern: ".*" + q + ".*", Options: "i"},
    	}}
	}
	filter := bson.M{"$and": and}

	if err := c.Find(&filter).Limit(5).All(&result); err != nil {
	  	fmt.Println("Failed to write result:", err)
	}

	return result
}

// AddProduct adds a Address in the DB
func (r Repository) AddAddress(address Address) bool {
	session, err := mgo.Dial(SERVER)
	defer session.Close()

	addressId += 1
	address.ID = addressId
	session.DB(DBNAME).C(COLLECTION).Insert(address)
	if err != nil {
		log.Fatal(err)
		return false
	}

	fmt.Println("Added New Address ID- ", address.ID)

	return true
}

// UpdateProduct updates a Product in the DB
func (r Repository) UpdateAddress(address Address) bool {
	session, err := mgo.Dial(SERVER)
	defer session.Close()
	log.Printf("In update address ............. : %d", address.ID)
	err = session.DB(DBNAME).C(COLLECTION).UpdateId(address.ID, address)
	
	if err != nil {
		log.Fatal(err)
		return false
	}

	fmt.Println("Updated Address ID - ", address.ID)

	return true
}

// DeleteProduct deletes an Product
func (r Repository) DeleteAddress(id int) string {
	session, err := mgo.Dial(SERVER)
	defer session.Close()

	// Remove product
	if err = session.DB(DBNAME).C(COLLECTION).RemoveId(id); err != nil {
		log.Fatal(err)
		return "INTERNAL ERR"
	}

	fmt.Println("Deleted Address ID - ", id)
	// Write status
	return "OK"
}

func (r Repository) FindAllAddresses() (Addresses, int){
     session, err := mgo.Dial(SERVER)
     defer session.Close()

     var result = Addresses{}
      err = session.DB(DBNAME).C(COLLECTION).Find(nil).All(&result)
      if err != nil {
      	 return result, 1
      }
      
      return result, 0

     
}