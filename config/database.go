package config

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() error {
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	// 🔹 STEP 1: connect tanpa database
	dsnRoot := fmt.Sprintf("%s:%s@tcp(%s:%s)/",
		user, pass, host, port,
	)

	sqlDB, err := sql.Open("mysql", dsnRoot)
	if err != nil {
		return err
	}

	// 🔹 STEP 2: create database jika belum ada
	_, err = sqlDB.Exec("CREATE DATABASE IF NOT EXISTS " + dbName)
	if err != nil {
		return err
	}

	// 🔹 STEP 3: connect ke database
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		user, pass, host, port, dbName,
	)

	gormDB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	DB = gormDB
	fmt.Println("Database ready")

	return nil
}
