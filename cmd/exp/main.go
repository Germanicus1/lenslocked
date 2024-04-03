package main

import (
	"database/sql"
	"fmt"

	_ "github.com/jackc/pgx/v4/stdlib"
)

type PostgresConfig struct {
	Host string
	Port string
	User string
	Password string
	DbName string
	SSLMode string
}

func (cfg PostgresConfig) String() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", cfg.Host, cfg.Port, cfg.User. cfg.Password, cfg.DbName, cfg.SSLMode)
}

func main() {
	db, err := sql.Open("pgx", "host=localhost port=6543 user=baloo password=junglebook dbname=lenslocked sslmode=disable")
	if err!=nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	d efer db.Close()
}