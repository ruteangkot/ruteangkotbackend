package controller

import (
	"net/http"

	"github.com/gocroot/config"
	"github.com/gocroot/model"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/gocroot/helper"
	"github.com/gocroot/helper/atdb"
	"github.com/gocroot/helper/watoken"
)

func GetDataUser(respw http.ResponseWriter, req *http.Request) {
	var docuser model.Userdomyikado
	httpstatus := http.StatusUnauthorized
	phonenumber, err := watoken.DecodeGetId(config.PublicKeyWhatsAuth, helper.GetSecretFromHeader(req))
	if err == nil {
		httpstatus = http.StatusOK
	}
	docuser, err = atdb.GetOneDoc[model.Userdomyikado](config.Mongoconn, "user", primitive.M{"phonenumber": phonenumber})
	if err != nil {
		httpstatus = http.StatusNotFound
	}
	helper.WriteResponse(respw, httpstatus, docuser)
}
