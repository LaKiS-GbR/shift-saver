package api

import (
	"fmt"
	"net/http"

	"github.com/LaKiS-GbR/shift-saver/pkg/database"
	"github.com/LaKiS-GbR/shift-saver/pkg/database/model"
	"github.com/LaKiS-GbR/shift-saver/pkg/router/util/limit"
	"github.com/LaKiS-GbR/shift-saver/pkg/router/util/session"
	"gorm.io/gorm"
)

func Login(writer http.ResponseWriter, request *http.Request) {
	err := limit.IsOverLimit(request)
	if err != nil {
		writer.WriteHeader(http.StatusTooManyRequests)
		fmt.Fprint(writer, "[login-1] too many requests")
		return
	}
	user := request.URL.Query().Get("username")
	pass := request.URL.Query().Get("password")
	if user == "" || pass == "" {
		handleUnauthorized(writer, "[login-2] user or password incorrect")
		return
	}
	db := database.GetInstance()
	if db == nil {
		handleInternalError(writer, "[login-3] database not initialized")
		return
	}
	candidate, err := getUserByUsername(db, user)
	if err != nil {
		handleUnauthorized(writer, "[login-4] user or password incorrect")
		return
	}
	if !verifyPassword(candidate.Password, pass) {
		handleUnauthorized(writer, "[login-5] user or password incorrect")
		return
	}
	handleSuccess(writer, user)
}

func getUserByUsername(db *gorm.DB, username string) (*model.User, error) {
	candidate := &model.User{}
	result := db.Where("username = ?", username).First(candidate)
	if result.Error != nil {
		return nil, result.Error
	}
	return candidate, nil
}

func verifyPassword(savedPassword, enteredPassword string) bool {
	// TODO: Implement password hashing with bcrypt and salt
	return savedPassword == enteredPassword
}

func handleUnauthorized(writer http.ResponseWriter, message string) {
	writer.WriteHeader(http.StatusUnauthorized)
	fmt.Fprint(writer, message)
}

func handleInternalError(writer http.ResponseWriter, message string) {
	writer.WriteHeader(http.StatusInternalServerError)
	fmt.Fprint(writer, message)
}

func handleSuccess(writer http.ResponseWriter, user string) {
	session.CreateSession(writer, user)
	writer.WriteHeader(http.StatusOK)
}
