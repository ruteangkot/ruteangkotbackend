package route

import (
	"fmt"
	"net/http"

	"github.com/gocroot/config"
	"github.com/gocroot/controller"
)

func URL(w http.ResponseWriter, r *http.Request) {
	// Log request method and URL path
	fmt.Printf("Received %s request for %s\n", r.Method, r.URL.Path)
	
	// Set Access Control Headers
	if config.SetAccessControlHeaders(w, r) {
		return
	}

	config.SetEnv()

	switch r.Method {
	case http.MethodGet:
		switch r.URL.Path {
		case "/":
			controller.GetHome(w, r)
		case "/data":
			controller.Getdatarouteangkot(w, r)
		default:
			controller.NotFound(w, r)
		}
	case http.MethodPost:
		if r.URL.Path == "/akun"{
			controller.Register(w, r)
		} else {
			controller.Login(w, r)
		}
		switch r.URL.Path  {
		case "/data":
			controller.CreateRoute(w, r)
		default:
			controller.NotFound(w, r)
		}
	case http.MethodPut:
		if r.URL.Path == "/data" {
			controller.UpdateRoute(w, r)
		} else {
			controller.NotFound(w, r)
		}
	case http.MethodDelete:
		if r.URL.Path == "/data" {
			controller.DeleteRoute(w, r)
		} else {
			controller.NotFound(w, r)
		}
	default:
		controller.NotFound(w, r)
	}
}

func SetAccessControlHeaders(w http.ResponseWriter, r *http.Request) bool {
	// Set CORS headers for the preflight request
	if r.Method == http.MethodOptions {
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type,Login")
		w.Header().Set("Access-Control-Allow-Methods", "POST,GET,PUT,DELETE")
		w.Header().Set("Access-Control-Allow-Origin", "https://ruteangkot.github.io")
		w.Header().Set("Access-Control-Max-Age", "3600")
		w.WriteHeader(http.StatusNoContent)
		return true
	}
	// Set CORS headers for the main request.
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Origin", "https://ruteangkot.github.io")
	w.Header().Set("Access-Control-Allow-Methods", "POST,GET,PUT,DELETE")
	return false
}
