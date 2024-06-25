package route

import (
	"net/http"

	"github.com/gocroot/config"
	"github.com/gocroot/controller"
)

func URL(w http.ResponseWriter, r *http.Request) {
	if config.SetAccessControlHeaders(w, r) {
		return
	}
	config.SetEnv()

	switch r.Method {
	case "GET":
		if r.URL.Path == "/" {
			controller.GetHome(w, r)
		} else if r.URL.Path == "/data" {
			controller.Getdatarouteangkot(w, r)
		} else {
			controller.NotFound(w, r)
		}
	case "POST":
		if r.URL.Path == "/data" {
			controller.CreateRoute(w, r)
		} else if r.URL.Path == "/register" {
			controller.Register(w, r)
		} else if r.URL.Path == "/login" {
			controller.Login(w, r)
		} else {
			controller.NotFound(w, r)
		}
	case "PUT":
		if r.URL.Path == "/data" {
			controller.UpdateRoute(w, r)
		} else {
			controller.NotFound(w, r)
		}
	case "DELETE":
		if r.URL.Path == "/data" {
			controller.DeleteRoute(w, r)
		} else {
			controller.NotFound(w, r)
		}
	default:
		controller.NotFound(w, r)
	}
}
