package database

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"

	"mid-the-ethereum-fetcher-zxowaz/logger"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func setupDB() (*sql.DB, error) {
	err := godotenv.Load("local.env")
	if err != nil {
		logger.GetInstance().ErrorLogger.Println(err)
		return nil, err
	}

	port, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		logger.GetInstance().ErrorLogger.Println(err)
		return nil, err
	}

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"), port, os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		logger.GetInstance().ErrorLogger.Println(err)
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		logger.GetInstance().ErrorLogger.Println(err)
		return nil, err
	}
	logger.GetInstance().InfoLogger.Println("database connected")
	return db, nil
}

var Database *sql.DB

func init() {
	db, err := setupDB()
	if err != nil {
		logger.GetInstance().ErrorLogger.Println(err)
		return
	}
	Database = db
}
