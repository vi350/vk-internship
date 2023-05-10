package postgres

import (
	"database/sql"
	"fmt"
	"github.com/vi350/vk-internship/internal/app/e"
	"log"
	"os"
)

type Config struct {
	DBName   string
	Host     string
	Port     string
	User     string
	Password string
}

func newPgOptions() *Config {
	var host string
	if os.Getenv("INSIDE_A_DOCKER") == "Yes" {
		host = os.Getenv("POSTGRES_CONTAINER_HOST")
	} else {
		host = os.Getenv("POSTGRES_HOST")
	}
	if host == "" {
		log.Panicf("Variable POSTGRES_HOST was not specified in the .env file")
	}
	port := os.Getenv("POSTGRES_PORT")
	if port == "" {
		log.Panicf("Variable POSTGRES_PORT was not specified in the .env file")
	}
	user := os.Getenv("POSTGRES_USER")
	if user == "" {
		log.Panicf("Variable POSTGRES_USER was not specified in the .env file")
	}
	password := os.Getenv("POSTGRES_PASSWORD")
	if password == "" {
		log.Panicf("Variable POSTGRES_PASSWORD was not specified in the .env file")
	}
	database := os.Getenv("POSTGRES_DB")
	if database == "" {
		log.Panicf("Variable POSTGRES_DB was not specified in the .env file")
	}

	return &Config{
		DBName:   database,
		Host:     host,
		Port:     port,
		User:     user,
		Password: password,
	}
}

func connect() (p *sql.DB, err error) {
	defer func() { err = e.WrapIfErr("error connecting to db: ", err) }()
	conf := newPgOptions()
	dbConnectionString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		conf.Host, conf.Port, conf.User, conf.Password, conf.DBName)

	db, err := sql.Open("postgres", dbConnectionString)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
