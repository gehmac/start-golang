package config

import (
	"os"

	"github.com/gehmac/anime-list-golang/schemas"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitializeSQLite() (*gorm.DB, error) {
	logger := GetLogger("sqlite")
	_, err := os.Stat("./db/anime.db")
	if os.IsNotExist(err) {
		logger.Info("database file not found, creating...")
		err = os.MkdirAll("./db", os.ModePerm)
		if err != nil {
			return nil, err
		}
		file, err := os.Create("./db/anime.db")
		if err != nil {
			return nil, err
		}
		defer file.Close()
	}
	db, err := gorm.Open(sqlite.Open("./db/anime.db"), &gorm.Config{})
	if err != nil {
		logger.Errf("sqlite opening error: %v", err)
		return nil, err
	}
	err = db.AutoMigrate(&schemas.Anime{})
	logger.Errf("sqlite autoMigrate error: %v", err)
	if err != nil {
		return nil, err
	}
	return db, nil
}
