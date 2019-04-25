package addrbk

import (
    "log"
    "net/http"
    "github.com/gorilla/mux"
)

var controller = &Controller{Repository: Repository{}}

// Route defines a route
type Route struct {
    Name        string
    Method      string
    Pattern     string
    HandlerFunc http.HandlerFunc
}

// Routes defines the list of routes of our API
type Routes []Route

var routes = Routes {
    Route {
        "Index",
        "GET",
        "/",
        controller.Index,
    },
    Route {
        "AddAddress",
        "POST",
        "/address",
        controller.AddAddress,
    },
    Route {
        "UpdateAddress",
        "PUT",
        "/address",
        controller.UpdateAddress,
    },
    // Get Address by {id}
    Route {
        "GetAddress",
        "GET",
        "/address/{id}",
        controller.GetAddress,
    },
    // Delete Address by {id}
    Route {
        "DeleteAddress",
        "DELETE",
        "/address/{id}",
        controller.DeleteAddress,
    },
    // Search Address with string
    Route {
        "SearchAddress",
        "GET",
        "/Searchaddresses/{query}",
        controller.SearchAddress,
    },
    // upload csv file
    Route {
       "Upload",
        "POST",
        "/upload",
        controller.UploadFile,
    },
    // download csv file
    Route {
       "DownloadCSV",
        "GET",
        "/downloadcsv",
        controller.DownloadCsv,
    }}

// NewRouter configures a new router to the API
func NewRouter() *mux.Router {
    router := mux.NewRouter().StrictSlash(true)
    for _, route := range routes { 
        var handler http.Handler
        log.Println(route.Name)
        handler = route.HandlerFunc
        
        router.
         Methods(route.Method).
         Path(route.Pattern).
         Name(route.Name).
         Handler(handler)
    }
    return router
}
