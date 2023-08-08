package web

import (
	"embed"
	"fmt"
	"net/http"
	"text/template"

	"github.com/LaKiS-GbR/shift-saver/pkg/router/util/limit"
	"github.com/LaKiS-GbR/shift-saver/pkg/router/util/session"
)

var (
	//go:embed static
	static embed.FS
)

func ProvideIndex(writer http.ResponseWriter, request *http.Request) {
	err := limit.IsOverLimit(request)
	if err != nil {
		writer.WriteHeader(http.StatusTooManyRequests)
		fmt.Fprintf(writer, "[provide-index-1] too many requests")
		return
	}
	if !session.IsSessionValid(request) {
		http.Redirect(writer, request, "/login", http.StatusFound)
		return
	}
	template, err := template.ParseFS(static, "static/html/pages/index.html")
	if err != nil {
		fmt.Fprintf(writer, "[provide-index-2] could not provide template - error: %s", err)
		return
	}
	writer.Header().Add("Content-Type", "text/html")
	err = template.Execute(writer, nil)
	if err != nil {
		fmt.Fprintf(writer, "[provide-index-3] could not execute parsed template - error: %v", err)
	}
}
