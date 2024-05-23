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
	payload, err := watoken.Decode(config.PublicKeyWhatsAuth, helper.GetSecretFromHeader(req))
	if err != nil {
		var respn model.Response
		respn.Status = "watoken.Decode"
		respn.Info = helper.GetSecretFromHeader(req)
		respn.Location = config.PublicKeyWhatsAuth
		respn.Response = err.Error()
		helper.WriteJSON(respw, http.StatusOK, respn)
		return
	}
	docuser, err = atdb.GetOneDoc[model.Userdomyikado](config.Mongoconn, "user", primitive.M{"phonenumber": payload.Id})
	if err != nil {
		docuser.PhoneNumber = payload.Id
		docuser.Name = payload.Alias
		docuser.Email = err.Error()
		helper.WriteJSON(respw, http.StatusNotFound, docuser)
		return
	}
	docuser.Email = "ada di database"
	helper.WriteJSON(respw, httpstatus, docuser)
}
