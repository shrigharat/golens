package main

import (
	"database/sql"
	"fmt"
	"lenslocked/models"

	_ "github.com/jackc/pgx/v4/stdlib"
)

type User struct {
	Name string
	Bio  string
}

type PostgresConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	SSLMode  string
}

func (cfg PostgresConfig) ToString() string {
	return fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host,
		cfg.Port,
		cfg.User,
		cfg.Password,
		cfg.DBName,
		cfg.SSLMode,
	)
}

func main() {
	pgcfg := PostgresConfig{
		Host:     "localhost",
		Port:     "5432",
		User:     "test_user",
		Password: "test123",
		DBName:   "lenslocked",
		SSLMode:  "disable",
	}
	// fmt.Println(pgcfg.ToString())
	db, err := sql.Open(
		"pgx",
		pgcfg.ToString(),
		// "postgresql://test_user@localhost:5432/lenslocked",
	)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	defer db.Close()
	fmt.Println("Connected to Database.")

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS users(
			id SERIAL PRIMARY KEY,
			name TEXT,
			email TEXT UNIQUE NOT NULL
		);
	`)
	if err != nil {
		panic(err)
	}
	fmt.Println("Tables created.")

	// name := "Shrishail Gharat"
	// email := "shri@test.com"
	// _, err = db.Exec(`
	// 	INSERT INTO users(name, email)
	// 	VALUES($1, $2);
	// `, name, email)

	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("Inserted user.")

	userService := models.UserService{
		DB: db,
	}
	newEmail := "bob@cat.com"
	newPassword := "bobthebuilder"
	newUser, err := userService.Create(newEmail, newPassword)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Created new user.")
	fmt.Println(newUser)
}
