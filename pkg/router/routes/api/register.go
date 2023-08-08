package api

import (
	"fmt"
	"net/http"
	"regexp"

	"github.com/LaKiS-GbR/shift-saver/pkg/config"
	"github.com/LaKiS-GbR/shift-saver/pkg/database"
	"github.com/LaKiS-GbR/shift-saver/pkg/database/model"
	"github.com/LaKiS-GbR/shift-saver/pkg/router/util/limit"
	"gorm.io/gorm"
)

func Register(writer http.ResponseWriter, request *http.Request) {
	err := limit.IsOverLimit(request)
	if err != nil {
		writer.WriteHeader(http.StatusTooManyRequests)
		fmt.Fprintf(writer, "[register-1] too many requests")
		return
	}
	if !config.Running.Register {
		writer.WriteHeader(http.StatusForbidden)
		fmt.Fprintf(writer, "[register-2] registration is disabled")
		return
	}
	err = request.ParseForm()
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(writer, "[register-3] %s", err)
		return
	}
	user := request.Form.Get("user")
	pass := request.Form.Get("password")
	email := request.Form.Get("email")
	// TODO: Remove spaces from input
	ok := validateInput(user, pass, email)
	if !ok {
		writer.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(writer, "[register-4] invalid input")
		return
	}
	db := database.GetInstance()
	if db == nil {
		writer.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(writer, "[register-7] database not initialized")
		return
	}
	newUser := &model.User{
		Username: user,
		Password: pass,
		Email:    email,
	}
	err = checkIfUserExists(db, newUser)
	if err == nil {
		writer.WriteHeader(http.StatusConflict)
		fmt.Fprintf(writer, "[register-8] user already exists")
		return
	}
	err = createNewUser(db, newUser)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(writer, "[register-9] %s", err)
		return
	}
	http.Redirect(writer, request, "/login", http.StatusFound)
}

func checkIfUserExists(db *gorm.DB, user *model.User) error {
	err := db.Where("username = ? OR email = ?", user.Username, user.Email).First(user).Error
	if err != nil {
		return err
	}
	return nil
}

func createNewUser(db *gorm.DB, user *model.User) error {
	err := db.Create(user).Error
	if err != nil {
		return err
	}
	return nil
}

func validateInput(user, pass, email string) bool {
	if !isValidUsername(user) || !isValidPassword(pass) || !isValidEmail(email) {
		return false
	}
	return true
}

func isValidUsername(username string) bool {
	return len(username) >= 4
}

func isValidPassword(password string) bool {
	return len(password) >= 8
}

func isValidEmail(email string) bool {
	match, _ := regexp.MatchString(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`, email)
	return match
}
