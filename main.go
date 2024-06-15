package gocroot

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
	"github.com/gocroot/model"
	"github.com/gocroot/route"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2"
)

func init() {
	functions.HTTP("WebHook", route.URL)
}

var session *mgo.Session

func init() {
    var err error
    session, err = mgo.Dial("localhost")
    if err != nil {
        log.Fatalf("Failed to connect to MongoDB: %v", err)
    }
    session.SetMode(mgo.Monotonic, true)
}

func registerUser(w http.ResponseWriter, r *http.Request) {
    var user model.User
    err := json.NewDecoder(r.Body).Decode(&user)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    collection := session.DB("userDB").C("users")
    err = collection.Insert(user)
    if err != nil {
        http.Error(w, "Error registering new user.", http.StatusInternalServerError)
        return
    }

    log.Printf("New user registered: %+v", user)

    w.WriteHeader(http.StatusOK)
    w.Write([]byte("User registered successfully!"))
}

func main() {
    router := mux.NewRouter()
    router.HandleFunc("/register", registerUser).Methods("POST")

    log.Println("Server started at :8080")
    log.Fatal(http.ListenAndServe(":8080", router))
}