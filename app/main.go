package app

import (
	"net/http"
	"github.com/douglasfsti/task-queue/handler"
)

func init() {
	http.HandleFunc("/", handler.Home)
	http.HandleFunc("/worker", handler.Worker)
}
