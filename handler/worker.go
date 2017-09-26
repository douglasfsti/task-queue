package handler

import (
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
	"net/http"
	"time"
)

func Worker(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)

	log.Infof(c, "\ncreateAt: %v\nnow:      %v\n\n", r.FormValue("createAt"), time.Now())
}

