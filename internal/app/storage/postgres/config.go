package postgres

import (
	"github.com/go-pg/pg"
	"github.com/vi350/vk-internship/internal/app/e"
	"log"
	"os"
)

func newPgOptions() *pg.Options {
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

	addr := host + ":" + port

	return &pg.Options{
		Addr:     addr,
		User:     user,
		Password: password,
		Database: database,
	}
}

func connect() (p *pg.DB, err error) {
	defer func() { err = e.WrapIfErr("error performing request: ", err) }()
	DB := pg.Connect(newPgOptions())

	if err = statusErr(DB); err != nil {
		return nil, err
	}

	return p, nil
}

func statusErr(p *pg.DB) (err error) {
	defer func() { err = e.WrapIfErr("error performing select 1: ", err) }()

	if _, err = p.Exec("SELECT 1"); err != nil {
		return err
	}

	return nil
}
