package api

import (
	"fmt"
	"net/http"

	"github.com/LaKiS-GbR/shift-saver/pkg/router/util/limit"
	"github.com/LaKiS-GbR/shift-saver/pkg/router/util/session"
)

func Logout(writer http.ResponseWriter, request *http.Request) {
	err := limit.IsOverLimit(request)
	if err != nil {
		writer.WriteHeader(http.StatusTooManyRequests)
		fmt.Fprintf(writer, "[logout-1] too many requests")
		return
	}
	session.RemoveSession(writer, request)
	http.Redirect(writer, request, "/login", http.StatusFound)
}
