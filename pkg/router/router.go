package router

import (
	"embed"
	"fmt"
	"io/fs"
	"net/http"
	"time"

	"github.com/LaKiS-GbR/shift-saver/pkg/config"
	"github.com/LaKiS-GbR/shift-saver/pkg/router/routes/api"
	"github.com/LaKiS-GbR/shift-saver/pkg/router/routes/web"
	"golang.org/x/crypto/acme/autocert"
)

var (
	//go:embed routes/web/static
	static embed.FS
	router *http.ServeMux
)

func Init() error {
	router = http.NewServeMux()
	err := provideHandlers()
	if err != nil {
		return err
	}
	server := &http.Server{
		Addr:              fmt.Sprintf(":%v", config.Running.Port),
		ReadTimeout:       3 * time.Second,
		ReadHeaderTimeout: 3 * time.Second,
		WriteTimeout:      3 * time.Second,
		IdleTimeout:       120 * time.Second,
		Handler:           router,
	}
	if config.Running.Host == "localhost" {
		err = server.ListenAndServe()
		if err != nil {
			return err
		}
		return nil
	}
	certer := autocert.Manager{
		Prompt:     autocert.AcceptTOS,
		HostPolicy: autocert.HostWhitelist(config.Running.Host),
		Cache:      autocert.DirCache(config.Running.CertCache),
	}
	server.TLSConfig = certer.TLSConfig()
	err = server.ListenAndServeTLS("", "")
	if err != nil {
		return err
	}
	return nil
}

func provideHandlers() error {
	// Handle static files
	fs, err := fs.Sub(static, "routes/web/static")
	if err != nil {
		return err
	}
	router.Handle("/js/", http.FileServer(http.FS(fs)))
	router.Handle("/css/", http.FileServer(http.FS(fs)))
	router.Handle("/img/", http.FileServer(http.FS(fs)))
	// Handle web routes
	router.HandleFunc("/", web.ProvideIndex)
	router.HandleFunc("/login/", web.ProvideLogin)
	// Handle api routes
	router.HandleFunc("/api/login/", api.Login)
	router.HandleFunc("/api/logout/", api.Logout)
	return nil
}
