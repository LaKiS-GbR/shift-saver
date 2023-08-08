package web

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/LaKiS-GbR/shift-saver/pkg/router/util/limit"
)

func ProvideLogin(writer http.ResponseWriter, request *http.Request) {
	err := limit.IsOverLimit(request)
	if err != nil {
		writer.WriteHeader(http.StatusTooManyRequests)
		fmt.Fprintf(writer, "[provide-login-1] too many requests")
		return
	}
	tpl, err := template.ParseFS(static, "static/html/pages/login.html")
	if err != nil {
		fmt.Fprintf(writer, "[provide-login-2] could not provide template - error: %s", err)
		return
	}
	writer.Header().Add("Content-Type", "text/html")
	err = tpl.Execute(writer, nil)
	if err != nil {
		fmt.Fprintf(writer, "[provide-login-3] could not execute parsed template - error: %v", err)
	}
}
