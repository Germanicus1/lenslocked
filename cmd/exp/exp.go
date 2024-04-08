package main

import (
	"fmt"

	_ "github.com/jackc/pgx/v4/stdlib"
	"githubn.com/Germanicus1/lenslocked/models"
)

func main() {
	cfg := models.DefaultPostgresConfig()
	db, err := models.Open(cfg)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		fmt.Println("Could not ping().")
		panic(err)
	}
	fmt.Println("Connected!")

	us := models.UserService{
		DB: db,
	}
	user, err := us.Create("peter@kerschbaumer.es", "peter123")
	if err != nil {
		panic(err)
	}
	fmt.Println(user)

	// name := "Peter"
	// email := "peter1@kerschbaumer.es"

	// // Insert some data
	// row := db.QueryRow(`
	// 	INSERT INTO users (name, email)
	// 	VALUES ($1, $2) RETURNING id`, name, email)
	// var id int
	// err = row.Scan(&id)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("User created. ID:", id)
}
