package database

import (
	"os"
	"path/filepath"
	"sync"

	"github.com/LaKiS-GbR/shift-saver/pkg/config"
	"github.com/LaKiS-GbR/shift-saver/pkg/database/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	instance *gorm.DB
	once     sync.Once
)

func Init() error {
	var err error
	once.Do(func() {
		err = os.MkdirAll(filepath.Dir(config.Running.DBPath), 0755)
		if err != nil {
			return
		}
		db, err := gorm.Open(sqlite.Open(config.Running.DBPath), &gorm.Config{})
		if err != nil {
			return
		}
		instance = db
		err = db.AutoMigrate(&model.User{}, &model.Shift{})
		if err != nil {
			return
		}
	})
	return err
}

func GetInstance() *gorm.DB {
	return instance
}
