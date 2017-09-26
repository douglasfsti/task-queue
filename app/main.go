package app

import (
	"net/http"
	"gitlab.com/bawi/task-queue/handler"
)

func init() {
	http.HandleFunc("/", handler.Home)
	http.HandleFunc("/worker", handler.Worker)
}
