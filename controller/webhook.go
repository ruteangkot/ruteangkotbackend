package controller

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/go-playground/webhooks/v6/github"
	"github.com/gocroot/config"
	"github.com/gocroot/helper"
	"github.com/gocroot/helper/atdb"
	"github.com/gocroot/model"
)

func PostWebHookGithub(respw http.ResponseWriter, req *http.Request) {
	var resp model.Response
	httpstatus := http.StatusUnauthorized
	hook, err := github.New(github.Options.Secret("MyGitHubSuperSecretSecret...?"))
	if err != nil {
		resp.Response = err.Error()
		httpstatus = http.StatusUnauthorized
	}
	payload, err := hook.Parse(req, github.ReleaseEvent, github.PullRequestEvent)
	if err != nil {
		resp.Response = err.Error()
		httpstatus = http.StatusServiceUnavailable
	}
	switch pyl := payload.(type) {
	case github.PushPayload:
		var komsg, msg string
		for i, komit := range pyl.Commits {
			kommsg := strings.TrimSpace(komit.Message)
			appd := strconv.Itoa(i+1) + ". " + kommsg + "\n_" + komit.Author.Name + "_\n"
			dokcommit := model.PushReport{
				Username: strings.ToLower(komit.Author.Username),
				Email:    komit.Author.Email,
				Repo:     pyl.Repository.URL,
				Ref:      pyl.Ref,
				Message:  kommsg,
			}
			atdb.InsertOneDoc(config.Mongoconn, "pushrepo", dokcommit)
			komsg += appd
		}
		msg = pyl.Pusher.Name + "\n" + pyl.Sender.Login + "\n" + pyl.Repository.Name + "\n" + pyl.Ref + "\n" + pyl.Repository.URL + "\n" + komsg
		dt := &model.TextMessage{
			To:       "msg.Chat_number",
			IsGroup:  false,
			Messages: msg,
		}
		profile, err := helper.GetAppProfile("WAPhoneNumber", config.Mongoconn)
		if err != nil {
			resp.Response = err.Error()
			httpstatus = http.StatusUnauthorized
		}
		resp, err = helper.PostStructWithToken[model.Response]("Token", profile.Token, dt, config.WAAPIMessage)
		if err != nil {
			resp.Response = err.Error()
			httpstatus = http.StatusUnauthorized
		}
	}
	helper.WriteResponse(respw, httpstatus, resp)
}
