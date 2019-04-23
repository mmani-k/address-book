package addrbk

import (
    "encoding/json"
    "io"
    "io/ioutil"
    "log"
    "net/http"
    "strings"
    "strconv"

    "github.com/gorilla/mux"
 
)

//Controller ...
type Controller struct {
    Repository Repository
}



// Index GET /
func (c *Controller) Index(w http.ResponseWriter, r *http.Request) {
    addresses := c.Repository.GetAddresses() // list of all addresses
    data, _ := json.Marshal(addresses)
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.WriteHeader(http.StatusOK)
    w.Write(data)
    return
}

// AddAddress POST /
func (c *Controller) AddAddress(w http.ResponseWriter, r *http.Request) {
    log.Println("In AddAdress.....")
    var address Address
    body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
    log.Println(body)
    if err != nil {
        log.Fatalln("Error AddAddress", err)
        w.WriteHeader(http.StatusInternalServerError)
        return
    }

    if err := r.Body.Close(); err != nil {
        log.Fatalln("Error AddAddress", err)
    }

    if err := json.Unmarshal(body, &address); err != nil { // unmarshall body contents as a type Candidate
        w.WriteHeader(422) // unprocessable entity
        log.Println(err)
        if err := json.NewEncoder(w).Encode(err); err != nil {
            log.Fatalln("Error AddAddress unmarshalling data", err)
            w.WriteHeader(http.StatusInternalServerError)
            return
        }
    }

    log.Println( "Address : ", address)
    success := c.Repository.AddAddress(address) // adds the Address to the DB
    if !success {
        w.WriteHeader(http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusCreated)
    return
}

// SearchAddress GET /
func (c *Controller) SearchAddress(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    log.Println(vars)

    query := vars["query"] // param query
    log.Println("Search Query - " + query);

    addresses := c.Repository.GetAddressByString(query)
    data, _ := json.Marshal(addresses)

    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.WriteHeader(http.StatusOK)
    w.Write(data)
    return
}

// UpdateAddress PUT /
func (c *Controller) UpdateAddress(w http.ResponseWriter, r *http.Request) {
    var address Address
    body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576)) // read the body of the request
    if err != nil {
        log.Fatalln("Error UpdateAddress", err)
        w.WriteHeader(http.StatusInternalServerError)
        return
    }

    if err := r.Body.Close(); err != nil {
        log.Fatalln("Error UpdateAddress", err)
    }

    if err := json.Unmarshal(body, &address); err != nil { // unmarshall body contents as a type Candidate
        w.Header().Set("Content-Type", "application/json; charset=UTF-8")
        w.WriteHeader(422) // unprocessable entity
        if err := json.NewEncoder(w).Encode(err); err != nil {
            log.Fatalln("Error UpdateAddress unmarshalling data", err)
            w.WriteHeader(http.StatusInternalServerError)
            return
        }
    }

    log.Println(address.ID)
    success := c.Repository.UpdateAddress(address) 
    
    if !success {
        w.WriteHeader(http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.WriteHeader(http.StatusOK)
    return
}

// GetAddress GET - Gets a single address by ID /
func (c *Controller) GetAddress(w http.ResponseWriter, r *http.Request) {
    log.Println("In GetAddress..........")
    vars := mux.Vars(r)
    log.Println(vars)

    id := vars["id"] // param id
    log.Println(id);

    addresid, err := strconv.Atoi(id);

    if err != nil {
        log.Fatalln("Error GetAddress", err)
    }

    result := c.Repository.GetAddressById(addresid)
    data, _ := json.Marshal(result)

    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.WriteHeader(http.StatusOK)
    w.Write(data)
    return
}

// DeleteAddress DELETE /
func (c *Controller) DeleteAddress(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    log.Println(vars)
    id := vars["id"] // param id
    log.Println(id);

    addressid, err := strconv.Atoi(id);

    if err != nil {
        log.Fatalln("Error DeleteAddress", err)
    }

    if err := c.Repository.DeleteAddress(addressid); err != "" {
        log.Println(err);
        if strings.Contains(err, "404") {
            w.WriteHeader(http.StatusNotFound)
        } else if strings.Contains(err, "500") {
            w.WriteHeader(http.StatusInternalServerError)
        }
        return
    }
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.WriteHeader(http.StatusOK)
    return
}
