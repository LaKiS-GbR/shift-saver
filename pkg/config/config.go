package config

import (
	"os"
	"strconv"
)

type Config struct {
	Port      int
	Host      string
	CertCache string
	DBPath    string
	Register  bool
}

var Running *Config

func Init() {
	Running = &Config{
		Port:      80,
		Host:      "localhost",
		CertCache: "./certs",
		DBPath:    "./shifts.db",
		Register:  false,
	}
	host := os.Getenv("HOST")
	if host != "" {
		Running.Host = host
	}
	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err == nil {
		Running.Port = port
	}
	certCache := os.Getenv("CERT_CACHE")
	if certCache != "" {
		Running.CertCache = certCache
	}
	dbPath := os.Getenv("DB_PATH")
	if dbPath != "" {
		Running.DBPath = dbPath
	}
	register, err := strconv.ParseBool(os.Getenv("REGISTER"))
	if err == nil {
		Running.Register = register
	}
}
