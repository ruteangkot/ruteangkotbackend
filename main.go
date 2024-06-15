package main

import (
	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
	"github.com/gocroot/route"
)

func init() {
    functions.HTTP("WebHook", route.URL)
}

