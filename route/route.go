package route

import (
	"net/http"

	"github.com/gocroot/config"
	"github.com/gocroot/controller"
)

func URL(w http.ResponseWriter, r *http.Request) {
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
		switch r.URL.Path {
		case "/data":
			controller.CreateRoute(w, r)
		case "/register":
			controller.Register(w, r)
		case "/login":
			controller.Login(w, r)
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
