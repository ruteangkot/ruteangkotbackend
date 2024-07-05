package gocroot

import (
	"net/http"

	"github.com/gocroot/config"
	"github.com/gocroot/controller"
	"github.com/gocroot/route"

	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
)

func init() {
	config.InitDB()
	functions.HTTP("WebHook", route.URL)
}
func main() {
	http.HandleFunc("/request-reset-password", controller.RequestResetPassword)
	http.HandleFunc("/reset-password", controller.ResetPassword)
}