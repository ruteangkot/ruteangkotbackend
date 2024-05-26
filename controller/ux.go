package controller

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/gocroot/config"
	"github.com/gocroot/helper"
	"github.com/gocroot/helper/atdb"
	"github.com/gocroot/helper/watoken"
	"github.com/gocroot/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetLaporan(respw http.ResponseWriter, req *http.Request) {
	id := helper.GetParam(req)
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		var respn model.Response
		respn.Status = "Error : ObjectID Tidak Valid"
		respn.Info = helper.GetSecretFromHeader(req)
		respn.Location = "Encode Object ID Error"
		respn.Response = err.Error()
		helper.WriteJSON(respw, http.StatusBadRequest, respn)
		return
	}
	hasil, err := atdb.GetOneLatestDoc[model.Laporan](config.Mongoconn, "uxlaporan", primitive.M{"_id": objectId})
	if err != nil {
		var respn model.Response
		respn.Status = "Error : Data laporan tidak di temukan"
		respn.Response = err.Error()
		helper.WriteJSON(respw, http.StatusNotImplemented, respn)
		return
	}
	helper.WriteJSON(respw, http.StatusOK, hasil)
}

func PostLaporan(respw http.ResponseWriter, req *http.Request) {
	//otorisasi dan validasi inputan
	payload, err := watoken.Decode(config.PublicKeyWhatsAuth, helper.GetLoginFromHeader(req))
	if err != nil {
		var respn model.Response
		respn.Status = "Error : Token Tidak Valid"
		respn.Info = helper.GetSecretFromHeader(req)
		respn.Location = "Decode Token Error"
		respn.Response = err.Error()
		helper.WriteJSON(respw, http.StatusForbidden, respn)
		return
	}
	var lap model.Laporan
	err = json.NewDecoder(req.Body).Decode(&lap)
	if err != nil {
		var respn model.Response
		respn.Status = "Error : Body tidak valid"
		respn.Response = err.Error()
		helper.WriteJSON(respw, http.StatusBadRequest, respn)
		return
	}
	if lap.Phone == "" || lap.Nama == "" || lap.Solusi == "" {
		var respn model.Response
		respn.Status = "Error : Telepon atau nama atau solusi tidak diisi"
		respn.Response = "Isi lebih lengkap dahulu"
		helper.WriteJSON(respw, http.StatusBadRequest, respn)
		return
	}
	docuser, err := atdb.GetOneDoc[model.Userdomyikado](config.Mongoconn, "user", primitive.M{"phonenumber": payload.Id})
	if err != nil {
		var respn model.Response
		respn.Status = "Error : Data user tidak di temukan"
		respn.Response = err.Error()
		helper.WriteJSON(respw, http.StatusNotImplemented, respn)
		return
	}
	//lojik inputan post
	lap.User = docuser
	lap.Phone = ValidasiNoHP(lap.Phone)
	lap.Petugas = docuser.Name
	lap.NoPetugas = docuser.PhoneNumber

	idlap, err := atdb.InsertOneDoc(config.Mongoconn, "uxlaporan", lap)
	if err != nil {
		var respn model.Response
		respn.Status = "Gagal Insert Database"
		respn.Response = err.Error()
		helper.WriteJSON(respw, http.StatusNotModified, respn)
		return
	}
	message := "*Permintaan Feedback Pekerjaan*\n" + "Petugas : " + docuser.Name + "\nDeskripsi:" + lap.Solusi + "\n Beri Nilai: " + "https://www.do.my.id/rate/#" + idlap.Hex()
	dt := &model.TextMessage{
		To:       lap.Phone,
		IsGroup:  false,
		Messages: message,
	}
	resp, err := helper.PostStructWithToken[model.Response]("Token", config.WAAPIToken, dt, config.WAAPIMessage)
	if err != nil {
		resp.Info = "Tidak berhak"
		resp.Response = err.Error()
		helper.WriteJSON(respw, http.StatusUnauthorized, resp)
		return
	}
	helper.WriteJSON(respw, http.StatusOK, lap)

}

func ValidasiNoHP(nomor string) string {
	nomor = strings.ReplaceAll(nomor, " ", "")
	nomor = strings.ReplaceAll(nomor, "+", "")
	nomor = strings.ReplaceAll(nomor, "-", "")
	return nomor
}
